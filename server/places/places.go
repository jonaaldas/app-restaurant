package places

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/jonaaldas/go-restaurant-crud/types"
	"github.com/redis/go-redis/v9"
)

func GetPlacesByText(textQuery string, latitude float64, longitude float64, radius float64, redisClient *redis.Client) ([]types.Restaurant, error) {
	apiKey := os.Getenv("PLACES_API_KEY")
	if apiKey == "" {
		return []types.Restaurant{}, fmt.Errorf("PLACES_API_KEY environment variable is not set")
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
	req.Header.Set("X-Goog-FieldMask", "places.id,places.displayName,places.location,places.rating,places.priceLevel,places.userRatingCount,places.formattedAddress,places.shortFormattedAddress,places.photos")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to make request: %v", err)
		return []types.Restaurant{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return []types.Restaurant{}, fmt.Errorf("text search HTTP %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read body: %v", err)
		return []types.Restaurant{}, err
	}

	// Debug: print raw JSON response
	fmt.Printf("Raw API Response: %s\n", string(body[:min(len(body), 1000)]))

	var textSearchResponse types.TextSearchResponse
	if err := json.Unmarshal(body, &textSearchResponse); err != nil {
		log.Printf("Failed to parse JSON: %v", err)
		return []types.Restaurant{}, err
	}

	if len(textSearchResponse.Places) == 0 {
		return []types.Restaurant{}, nil
	}
	for i, place := range textSearchResponse.Places {
		fmt.Printf("Place %d: Name=%s, Photos=%+v, Address=%s\n", i, place.DisplayName.Text, place.Photos, place.FormattedAddress)
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

	// Use FormattedAddress if available, otherwise fall back to ShortFormattedAddress
	address := place.FormattedAddress
	if address == "" {
		address = place.ShortFormattedAddress
	}

	restaurant := types.Restaurant{
		Name:   place.DisplayName.Text,
		Rating: float32(place.Rating),
		Photos: place.Photos,
		Location: types.Location{
			Lat: place.Location.Latitude,
			Lng: place.Location.Longitude,
		},
		PlaceID:          place.ID,
		WouldTry:         false,
		Reviews:          reviewResponse.Result,
		FormattedAddress: address,
		PriceLevel:       place.PriceLevel,
	}

	return restaurant, nil
}
