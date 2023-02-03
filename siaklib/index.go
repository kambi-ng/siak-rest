package siaklib

import (
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type News struct {
	Title   string
	Content string
}

type UserOverview struct {
	Username string
	Identity string
	Role     string
}

type Homepage struct {
	User UserOverview
	News []News
}

func ParseWelcomePage(r io.Reader) (*Homepage, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	infoRow := doc.Find(".infocol > dl > dd")
	username := infoRow.Get(0).FirstChild.Data
	identity := infoRow.Get(1).FirstChild.Data
	role := infoRow.Get(2).FirstChild.Data

	var news []News
	doc.Find(".newsitem").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".title").Text()
		content := strings.TrimSpace(s.Find(".content").Text())
		news = append(news, News{title, content})
	})

	return &Homepage{UserOverview{username, identity, role}, news}, nil
}
