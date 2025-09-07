package types

type Place struct {
	BusinessStatus      string   `json:"business_status"`
	Geometry            Geometry `json:"geometry"`
	Icon                string   `json:"icon"`
	IconBackgroundColor string   `json:"icon_background_color"`
	IconMaskBaseURI     string   `json:"icon_mask_base_uri"`
	Name                string   `json:"name"`
	OpeningHours        struct {
		OpenNow bool `json:"open_now"`
	} `json:"opening_hours"`
	Photos           []Photo  `json:"photos"`
	PlaceID          string   `json:"place_id"`
	PlusCode         PlusCode `json:"plus_code"`
	PriceLevel       int      `json:"price_level"`
	Rating           float64  `json:"rating"`
	Reference        string   `json:"reference"`
	Scope            string   `json:"scope"`
	Types            []string `json:"types"`
	UserRatingsTotal int      `json:"user_ratings_total"`
	Vicinity         string   `json:"vicinity"`
}

type Geometry struct {
	Location Location `json:"location"`
	Viewport Viewport `json:"viewport"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Viewport struct {
	Northeast Location `json:"northeast"`
	Southwest Location `json:"southwest"`
}

type Photo struct {
	Height           int      `json:"height"`
	HTMLAttributions []string `json:"html_attributions"`
	PhotoReference   string   `json:"photo_reference"`
	Width            int      `json:"width"`
}

type PlusCode struct {
	CompoundCode string `json:"compound_code"`
	GlobalCode   string `json:"global_code"`
}

type GoogleAPIPlaceMaster struct {
	Results       []Place `json:"results"`
	Status        string  `json:"status"`
	NextPageToken string  `json:"next_page_token,omitempty"`
	ErrorMessage  string  `json:"error_message,omitempty"`
}
