package types

type Restaurant struct {
	Name     string   `json:"name"`
	Rating   float32  `json:"rating"`
	Photos   []Photo  `json:"photos"`
	Location Location `json:"location"`
	WouldTry bool     `json:"would_try"`
}
