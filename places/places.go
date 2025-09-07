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

func GetPlaces(location string, radius int, resType string) (types.GoogleAPIPlaceMaster, error) {
	apiKey := os.Getenv("PLACES_API_KEY")
	fmt.Print(apiKey)
	if apiKey == "" {
		return types.GoogleAPIPlaceMaster{}, fmt.Errorf("PLACES_API_KEY environment variable is not set")
	}

	url := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=" + location + "&radius=" + strconv.Itoa(radius) + "&type=" + resType + "&key=" + apiKey
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

	fmt.Printf("API Response Status: %s\n", places.Status)
	if places.ErrorMessage != "" {
		fmt.Printf("Error Message: %s\n", places.ErrorMessage)
	}

	if places.Status != "OK" {
		return types.GoogleAPIPlaceMaster{}, fmt.Errorf("Google Places API error: %s - %s", places.Status, places.ErrorMessage)
	}

	fmt.Printf("Found %d places\n", len(places.Results))
	return places, nil
}
