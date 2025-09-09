package places

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jonaaldas/go-restaurant-crud/types"
)

func GetPlaces(latlong string, radius int, resType string) ([]types.Restaurant, error) {
	places := []types.Restaurant{}
	apiKey := os.Getenv("PLACES_API_KEY")
	if apiKey == "" {
		return []types.Restaurant{}, fmt.Errorf("PLACES_API_KEY environment variable is not set")
	}

	url := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=" + latlong + "&radius=" + strconv.Itoa(radius) + "&type=" + resType + "&key=" + apiKey
	resp, err := http.Get(url)

	var googlePlacesRes types.GoogleAPIPlaceMaster

	if err != nil {
		log.Fatal("Failed to make request:", err)
		return []types.Restaurant{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read body:", err)
		return []types.Restaurant{}, err
	}

	if err := json.Unmarshal(body, &googlePlacesRes); err != nil {
		log.Fatal("Failed to parse JSON:", err)
		return []types.Restaurant{}, err
	}

	if len(googlePlacesRes.Results) == 0 {
		return []types.Restaurant{}, err
	}

	for _, place := range googlePlacesRes.Results {

		reviewsUrl := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/details/json?place_id=%s&fields=reviews,rating&key=%s", place.PlaceID, apiKey)

		reviewRes, err := http.Get(reviewsUrl)
		if err != nil {
			log.Fatal("Failed to get reviews:", err)
			return []types.Restaurant{}, err
		}
		defer reviewRes.Body.Close()

		reviewBody, err := io.ReadAll(reviewRes.Body)

		if err != nil {
			log.Fatal("Failed to read review body:", err)
			return []types.Restaurant{}, err
		}

		var reviewResponse types.GoogleReviewsReply

		if err := json.Unmarshal(reviewBody, &reviewResponse); err != nil {
			log.Fatal("Failed to parse JSON:", err)
			return []types.Restaurant{}, err
		}

		restaurant := types.Restaurant{
			Name:     place.Name,
			Rating:   float32(place.Rating),
			Photos:   place.Photos,
			Location: place.Geometry.Location,
			PlaceID:  place.PlaceID,
			WouldTry: false,
			Reviews:  reviewResponse.Result,
		}

		places = append(places, restaurant)

	}

	return places, nil
}
