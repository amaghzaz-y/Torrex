package api

import (
	"net/http"

	model "github.com/amaghzaz-y/torrex/internal/models"
	"github.com/amaghzaz-y/torrex/internal/scraper"
	"github.com/gofiber/fiber/v2"
	nanoid "github.com/matoous/go-nanoid"
)

// returns model.Room as json
func (a *Api) searchHandler(c *fiber.Ctx) error {
	queryParam := c.Params("query")
	if queryParam == "" {
		return c.Status(http.StatusBadRequest).SendString("invalid request")
	}
	magnetChan := make(chan string)
	movieChan := make(chan model.Movie)
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
			movieChan <- model.Movie{}
		}
		movieChan <- res
	}(queryParam)
	movie := <-movieChan
	magnet := <-magnetChan
	if movie.Title == "" || magnet == "" {
		return c.SendStatus(404)
	}
	id, err := nanoid.Nanoid(32)
	if err != nil {
		return c.SendStatus(502)
	}
	res := model.Room{
		Id:     id,
		Movie:  movie,
		Magnet: magnet,
	}
	err = a.Store.UpsertMovie(magnet, &movie)
	if err != nil {
		return c.SendStatus(502)
	}
	err = a.Store.UpsertRoom(&res)
	if err != nil {
		return c.SendStatus(502)
	}
	return c.JSON(res)
}
