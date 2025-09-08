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

func GetPlaces(latlong string, radius int, resType string) (types.GoogleAPIPlaceMaster, error) {
	fmt.Print(latlong)
	apiKey := os.Getenv("PLACES_API_KEY")
	if apiKey == "" {
		return types.GoogleAPIPlaceMaster{}, fmt.Errorf("PLACES_API_KEY environment variable is not set")
	}

	url := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=" + latlong + "&radius=" + strconv.Itoa(radius) + "&type=" + resType + "&key=" + apiKey
	fmt.Printf("Request URL: %s\n", url)
	resp, err := http.Get(url)

	var places types.GoogleAPIPlaceMaster

	if err != nil {
		log.Fatal("Failed to make request:", err)
		return types.GoogleAPIPlaceMaster{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read body:", err)
		return types.GoogleAPIPlaceMaster{}, err
	}

	if err := json.Unmarshal(body, &places); err != nil {
		log.Fatal("Failed to parse JSON:", err)
		return types.GoogleAPIPlaceMaster{}, err
	}

	for i, place := range places.Results {
		if i == 1 {
			break
		}

		reviewsUrl := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/details/json?place_id=%s&fields=name,reviews&key=%s", place.PlaceID, apiKey)
		fmt.Printf("Reviews URL: %s\n", reviewsUrl)
		reviewRes, err := http.Get(reviewsUrl)
		if err != nil {
			log.Fatal("Failed to get reviews:", err)
			return types.GoogleAPIPlaceMaster{}, err
		}
		defer reviewRes.Body.Close()

		reviewBody, err := io.ReadAll(reviewRes.Body)
		if err != nil {
			log.Fatal("Failed to read review body:", err)
			return types.GoogleAPIPlaceMaster{}, err
		}

		var reviewResponse types.GoogleAPIPlaceMaster
		if err := json.Unmarshal(reviewBody, &reviewResponse); err != nil {
			log.Fatal("Failed to parse JSON:", err)
			return types.GoogleAPIPlaceMaster{}, err
		}

		fmt.Printf("%+v\n", reviewResponse)
	}
	return places, nil
}
