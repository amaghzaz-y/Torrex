package scraper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var rarbg = []string{"https://www.rarbgproxy.to/search/?search=", "&order=seeders&by=DESC"}

const rarbg_host = "https://www.rarbgproxy.to"

func FetchMovieLink(query string) (string, error) {
	query = strings.ReplaceAll(query, " ", "+")
	url := fmt.Sprintf("%s%s%s", rarbg[0], query, rarbg[1])
	res, err := http.DefaultClient.Get(url)
	if err != nil {
		return "", fmt.Errorf("error: movie search on yts failed, status : %s", res.Status)
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}
	x := doc.Find("tr .table2ta_rarbgproxy").Children().Filter("td .tlista_rarbgproxy").Children().Filter("a")
	url = x.Nodes[1].Attr[0].Val
	url = fmt.Sprintf("%s%s", rarbg_host, url)
	return url, nil
}
