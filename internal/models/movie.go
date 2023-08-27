package model

type Movie struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	TagLine     string `json:"tagline"`
	Year        string `json:"year"`
	Score       string `json:"score"`
	Url         string `json:"url"`
	Poster      string `json:"poster"`
	Trailer     string `json:"trailer"`
	BgImg       string `json:"bgimg"`
}
