package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/amaghzaz-y/torrex/internal/scraper"
	"github.com/go-chi/chi/v5"
)

type SearchResponse struct {
	Movie  scraper.MovieInfo `json:"movie"`
	Magnet string            `json:"magnet"`
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	queryParam := chi.URLParam(r, "query")
	if queryParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error: bad request" + queryParam))
		return
	}
	magnetChan := make(chan string)
	infoChan := make(chan scraper.MovieInfo)
	go func(query string) {
		res, err := scraper.Torrent().Magnet(query)
		if err != nil {
			log.Println("error searching for magnet :", query)
			magnetChan <- ""
		}
		magnetChan <- res
	}(queryParam)
	go func(query string) {
		res, err := scraper.Info().Movie(query)
		if err != nil {
			log.Println("error searching for magnet :", query)
			infoChan <- scraper.MovieInfo{}
		}
		infoChan <- res
	}(queryParam)
	info := <-infoChan
	magnet := <-magnetChan
	res := &SearchResponse{
		info,
		magnet,
	}
	payload, err := json.Marshal(&res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Write(payload)
}
