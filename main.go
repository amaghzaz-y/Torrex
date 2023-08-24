package main

import (
	"fmt"
	_ "net/http/pprof"

	"github.com/amaghzaz-y/torrex/internal/scraper"
)

func main() {
	link, err := scraper.FetchMovieLink("spider man homecoming")
	if err != nil {
		panic(err)
	}
	mag, err := scraper.FetchMovieMagnet(link)
	if err != nil {
		panic(err)
	}
	fmt.Println(mag)
}
