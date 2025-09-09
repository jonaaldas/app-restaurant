package places

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/jonaaldas/go-restaurant-crud/types"
)

func GetPlaces(latlong string, radius int, resType string) ([]types.Restaurant, error) {
	apiKey := os.Getenv("PLACES_API_KEY")
	if apiKey == "" {
		return []types.Restaurant{}, fmt.Errorf("PLACES_API_KEY environment variable is not set")
	}

	url := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=" + latlong + "&radius=" + strconv.Itoa(radius) + "&type=" + resType + "&key=" + apiKey
	resp, err := http.Get(url)
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

	var googlePlacesRes types.GoogleAPIPlaceMaster

	if err := json.Unmarshal(body, &googlePlacesRes); err != nil {
		log.Printf("Failed to parse JSON: %v", err)
		return []types.Restaurant{}, err
	}

	if len(googlePlacesRes.Results) == 0 {
		return []types.Restaurant{}, nil
	}

	// Use concurrent processing for fetching reviews
	places := make([]types.Restaurant, len(googlePlacesRes.Results))
	var wg sync.WaitGroup
	errChan := make(chan error, len(googlePlacesRes.Results))

	for i, place := range googlePlacesRes.Results {
		wg.Add(1)
		go func(index int, p types.Place) {
			defer wg.Done()

			restaurant, err := fetchPlaceWithReviews(p, apiKey)

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

func fetchPlaceWithReviews(place types.Place, apiKey string) (types.Restaurant, error) {
	reviewsUrl := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/details/json?place_id=%s&fields=reviews,rating&key=%s", place.PlaceID, apiKey)

	reviewRes, err := http.Get(reviewsUrl)
	if err != nil {
		return types.Restaurant{}, fmt.Errorf("failed to get reviews for %s: %w", place.Name, err)
	}
	defer reviewRes.Body.Close()

	reviewBody, err := io.ReadAll(reviewRes.Body)
	if err != nil {
		return types.Restaurant{}, fmt.Errorf("failed to read review body for %s: %w", place.Name, err)
	}

	var reviewResponse types.GoogleReviewsReply
	if err := json.Unmarshal(reviewBody, &reviewResponse); err != nil {
		return types.Restaurant{}, fmt.Errorf("failed to parse reviews JSON for %s: %w", place.Name, err)
	}

	return types.Restaurant{
		Name:     place.Name,
		Rating:   float32(place.Rating),
		Photos:   place.Photos,
		Location: place.Geometry.Location,
		PlaceID:  place.PlaceID,
		WouldTry: false,
		Reviews:  reviewResponse.Result,
	}, nil
}
