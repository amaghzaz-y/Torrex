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

func (*LTBXD) FetchMovieInfoLink(query string) (string, error) {
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
