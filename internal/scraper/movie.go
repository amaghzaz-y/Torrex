package scraper

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const letterboxd = "https://letterboxd.com/search/"
const letterboxd_host = "https://letterboxd.com"

type LTBXD struct{}

func Info() *LTBXD {
	return &LTBXD{}
}

type MovieInfo struct {
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

func (*LTBXD) fetchMovieInfoLink(query string) (string, error) {
	query = strings.ReplaceAll(query, " ", "+")
	url := fmt.Sprintf("%s%s", letterboxd, query)
	res, err := http.DefaultClient.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}
	link, exist := doc.Find("li").First().Find("a").First().Attr("href")
	if !exist {
		return "", errors.New("movie not found")
	}
	link = fmt.Sprintf("%s%s", letterboxd_host, link)
	return link, nil
}

func (*LTBXD) fetchMovieInfo(movielink string) (*MovieInfo, error) {
	res, err := http.DefaultClient.Get(movielink)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	bg, _ := doc.Find("#backdrop").Attr("data-backdrop")
	frame, _ := doc.Find("#poster-zoom").Find("img").First().Attr("src")
	title := doc.Find("#featured-film-header").Find("h1").First().Text()
	year := doc.Find("#featured-film-header").Find("small").First().Text()
	tagline := doc.Find("h4").First().Text()
	desc := doc.Find("div .truncate").First().Text()
	score, _ := doc.Find("meta[name='twitter:data2']").Attr("content")
	trailer, _ := doc.Find("div .header").First().Find("a").First().Attr("href")
	info := &MovieInfo{
		Title:       title,
		TagLine:     tagline,
		Year:        year,
		Description: strings.TrimSpace(desc),
		Poster:      frame,
		BgImg:       bg,
		Trailer:     strings.TrimPrefix(trailer, "//"),
		Score:       score,
		Url:         movielink,
	}
	return info, nil
}

func (l *LTBXD) Movie(query string) (*MovieInfo, error) {
	link, err := l.fetchMovieInfoLink(query)
	if err != nil {
		return nil, err
	}
	return l.fetchMovieInfo(link)
}
