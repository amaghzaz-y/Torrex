package api

import (
	"net/http"

	model "github.com/amaghzaz-y/torrex/internal/models"
	"github.com/amaghzaz-y/torrex/internal/scraper"
	"github.com/gofiber/fiber/v2"
)

type SearchResponse struct {
	Movie  model.Movie `json:"movie"`
	Magnet string          `json:"magnet"`
}

func searchHandler(c *fiber.Ctx) error {
	queryParam := c.Params("query")
	if queryParam == "" {
		return c.Status(http.StatusBadRequest).SendString("invalid request")
	}
	magnetChan := make(chan string)
	infoChan := make(chan model.Movie)
	go func(query string) {
		res, err := scraper.Torrent().Magnet(query)
		if err != nil {
			magnetChan <- ""
		}
		magnetChan <- res
	}(queryParam)
	go func(query string) {
		res, err := scraper.Info().Movie(query)
		if err != nil {
			infoChan <- model.Movie{}
		}
		infoChan <- res
	}(queryParam)
	info := <-infoChan
	magnet := <-magnetChan
	if info.Title == "" || magnet == "" {
		return c.SendStatus(404)
	}
	res := &SearchResponse{
		info,
		magnet,
	}
	return c.JSON(res)
}
