package api

import (
	"net/http"

	model "github.com/amaghzaz-y/torrex/internal/models"
	"github.com/amaghzaz-y/torrex/internal/scraper"
	"github.com/labstack/echo/v4"
	nanoid "github.com/matoous/go-nanoid"
)

// GET /rooms
func (a *Api) RoomListHandler(c echo.Context) error {
	return c.JSON(200, a.Rooms())
}

// GET /rooms/:id

func (a *Api) RoomInfoHanlder(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "room id is null")
	}
	room, err := a.Store.GetRoom(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(200, room)
}

// returns model.Room as json
// GET /search/:query
func (a *Api) searchHandler(c echo.Context) error {
	queryParam := c.Param("query")
	if queryParam == "" {
		return c.String(http.StatusBadRequest, "invalid request")
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
		return c.NoContent(404)
	}
	id, err := nanoid.Nanoid(12)
	if err != nil {
		return c.NoContent(502)
	}
	res := model.Room{
		Id:     id,
		Movie:  movie,
		Magnet: magnet,
	}
	err = a.Store.UpsertMovie(magnet, &movie)
	if err != nil {
		return c.NoContent(502)
	}
	err = a.Store.UpsertRoom(&res)
	if err != nil {
		return c.NoContent(502)
	}
	return c.JSON(200, res)
}
