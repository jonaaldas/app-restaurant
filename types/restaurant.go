package types

type Restaurant struct {
	Name     string              `json:"name"`
	Rating   float32             `json:"rating"`
	Photos   []Photo             `json:"photos"`
	Location Location            `json:"location"`
	PlaceID  string              `json:"place_id"`
	WouldTry bool                `json:"would_try"`
	Reviews  GoogleReviewsResult `json:"reviews"`
}

type RestaurantId struct {
	PlaceID string `json:"place_id"`
}
