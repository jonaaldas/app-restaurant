package types

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type AuthorAttribution struct {
	DisplayName string `json:"displayName"`
	URI         string `json:"uri"`
	PhotoURI    string `json:"photoUri"`
}

type Photo struct {
	Name               string              `json:"name"`
	WidthPx            int                 `json:"widthPx"`
	HeightPx           int                 `json:"heightPx"`
	AuthorAttributions []AuthorAttribution `json:"authorAttributions"`
	FlagContentURI     string              `json:"flagContentUri"`
	GoogleMapsURI      string              `json:"googleMapsUri"`
}

type GoogleReviewsReply struct {
	HTMLAttributions []interface{}       `json:"html_attributions"`
	Result           GoogleReviewsResult `json:"result"`
	Status           string              `json:"status"`
}

type GoogleReviewsResult struct {
	Photos  []GoogleReviewsPhoto  `json:"photos"`
	Rating  float64               `json:"rating"`
	Reviews []GoogleReviewsReview `json:"reviews"`
}

type GoogleReviewsPhoto struct {
	Height           int      `json:"height"`
	HTMLAttributions []string `json:"html_attributions"`
	PhotoReference   string   `json:"photo_reference"`
	Width            int      `json:"width"`
}

type GoogleReviewsReview struct {
	AuthorName              string `json:"author_name"`
	AuthorURL               string `json:"author_url"`
	Language                string `json:"language"`
	OriginalLanguage        string `json:"original_language"`
	ProfilePhotoURL         string `json:"profile_photo_url"`
	Rating                  int    `json:"rating"`
	RelativeTimeDescription string `json:"relative_time_description"`
	Text                    string `json:"text"`
	Time                    int64  `json:"time"`
	Translated              bool   `json:"translated"`
}

// New Google Places API v1 Text Search Types
type TextSearchRequest struct {
	TextQuery    string       `json:"textQuery"`
	LocationBias LocationBias `json:"locationBias"`
}

type LocationBias struct {
	Circle Circle `json:"circle"`
}

type Circle struct {
	Center Center  `json:"center"`
	Radius float64 `json:"radius"`
}

type Center struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type TextSearchResponse struct {
	Places []TextSearchPlace `json:"places"`
}

type TextSearchPlace struct {
	ID                    string         `json:"id"`
	FormattedAddress      string         `json:"formattedAddress"`
	ShortFormattedAddress string         `json:"shortFormattedAddress"`
	Location              TextSearchLoc  `json:"location"`
	Rating                float64        `json:"rating"`
	PriceLevel            string         `json:"priceLevel,omitempty"`
	UserRatingCount       int            `json:"userRatingCount"`
	DisplayName           TextSearchName `json:"displayName"`
	Photos                []Photo        `json:"photos,omitempty"`
}

type TextSearchLoc struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type TextSearchName struct {
	Text         string `json:"text"`
	LanguageCode string `json:"languageCode"`
}
