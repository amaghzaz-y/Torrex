package model

type Room struct {
	Id     string `json:"id"`
	Movie  Movie  `json:"movie"`
	Magnet string
	Path   string
}
