package places

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/jonaaldas/go-restaurant-crud/database"
	"github.com/jonaaldas/go-restaurant-crud/types"
)

func GetPlacesByText(textQuery string, latitude float64, longitude float64, radius float64) ([]types.Restaurant, error) {
	apiKey := os.Getenv("PLACES_API_KEY")
	if apiKey == "" {
		return []types.Restaurant{}, fmt.Errorf("PLACES_API_KEY environment variable is not set")
	}

	ctx := context.Background()
	redis := database.InitRedis()

	cacheKey := fmt.Sprintf("text_search_%s_%.6f_%.6f_%.0f", textQuery, latitude, longitude, radius)
	cachedRestaurants, found := database.GetTextSearch(ctx, redis, cacheKey)

	if found {
		log.Printf("Cache hit! Returning %d restaurants for text search", len(cachedRestaurants))
		return cachedRestaurants, nil
	}

	requestBody := types.TextSearchRequest{
		TextQuery: textQuery,
		LocationBias: types.LocationBias{
			Circle: types.Circle{
				Center: types.Center{
					Latitude:  latitude,
					Longitude: longitude,
				},
				Radius: radius,
			},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return []types.Restaurant{}, fmt.Errorf("failed to marshal request body: %w", err)
	}

	url := "https://places.googleapis.com/v1/places:searchText"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return []types.Restaurant{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Goog-Api-Key", apiKey)
	req.Header.Set("X-Goog-FieldMask", "places.id,places.displayName,places.formattedAddress,places.location,places.rating,places.priceLevel,places.userRatingCount")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to make request: %v", err)
		return []types.Restaurant{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read body: %v", err)
		return []types.Restaurant{}, err
	}

	var textSearchResponse types.TextSearchResponse
	if err := json.Unmarshal(body, &textSearchResponse); err != nil {
		log.Printf("Failed to parse JSON: %v", err)
		return []types.Restaurant{}, err
	}

	if len(textSearchResponse.Places) == 0 {
		return []types.Restaurant{}, nil
	}

	places := make([]types.Restaurant, len(textSearchResponse.Places))
	var wg sync.WaitGroup
	errChan := make(chan error, len(textSearchResponse.Places))

	for i, place := range textSearchResponse.Places {
		wg.Add(1)
		go func(index int, p types.TextSearchPlace) {
			defer wg.Done()

			restaurant, err := fetchTextSearchPlaceWithReviews(p, apiKey)
			if err != nil {
				errChan <- err
				return
			}
			places[index] = restaurant
		}(i, place)
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		return []types.Restaurant{}, <-errChan
	}

	go func() {
		success := database.SetTextSearch(ctx, redis, cacheKey, places)
		if success {
			log.Printf("Successfully cached %d restaurants for text search", len(places))
		} else {
			log.Printf("Failed to cache restaurants for text search")
		}
	}()

	go func() {
		var wg sync.WaitGroup
		for _, restaurant := range places {
			wg.Add(1)
			go func(r types.Restaurant) {
				defer wg.Done()
				success := database.SetRestaurant(ctx, redis, r.PlaceID, r)
				if success {
					log.Printf("Successfully cached individual restaurant %s", r.PlaceID)
				} else {
					log.Printf("Failed to cache individual restaurant %s", r.PlaceID)
				}
			}(restaurant)
		}
		wg.Wait()
		log.Printf("Finished caching all %d individual restaurants", len(places))
	}()

	return places, nil
}

func fetchTextSearchPlaceWithReviews(place types.TextSearchPlace, apiKey string) (types.Restaurant, error) {
	reviewsUrl := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/details/json?place_id=%s&fields=reviews,rating&key=%s", place.ID, apiKey)

	reviewRes, err := http.Get(reviewsUrl)
	if err != nil {
		return types.Restaurant{}, fmt.Errorf("failed to get reviews for %s: %w", place.DisplayName.Text, err)
	}
	defer reviewRes.Body.Close()

	reviewBody, err := io.ReadAll(reviewRes.Body)
	if err != nil {
		return types.Restaurant{}, fmt.Errorf("failed to read review body for %s: %w", place.DisplayName.Text, err)
	}

	var reviewResponse types.GoogleReviewsReply
	if err := json.Unmarshal(reviewBody, &reviewResponse); err != nil {
		return types.Restaurant{}, fmt.Errorf("failed to parse reviews JSON for %s: %w", place.DisplayName.Text, err)
	}

	restaurant := types.Restaurant{
		Name:   place.DisplayName.Text,
		Rating: float32(place.Rating),
		Photos: []types.Photo{},
		Location: types.Location{
			Lat: place.Location.Latitude,
			Lng: place.Location.Longitude,
		},
		PlaceID:  place.ID,
		WouldTry: false,
		Reviews:  reviewResponse.Result,
	}

	return restaurant, nil
}
