package store

import (
	"testing"

	model "github.com/amaghzaz-y/torrex/internal/models"
)

func TestMovieRW(t *testing.T) {
	s := New("torrex.dev.data")

	movie := &model.Movie{
		Url:   "http://localmovie.to",
		Title: "Some Random",
	}
	err := s.UpsertMovie("somemagnet", movie)
	if err != nil {
		t.Fatal(err)
	}
	m, err := s.GetMovieByUrl("somemagnet")
	if err != nil {
		t.Fatal(err)
	}
	if m.Url != movie.Url || m.Title != movie.Title {
		t.Fatal(m.Url, m.Title, "!=", movie.Url, movie.Title)
	}
}
