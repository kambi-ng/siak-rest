package siaklib

import (
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type News struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UserOverview struct {
	Username string `json:"username"`
	Identity string `json:"identity"`
	Role     string `json:"role"`
}

type Homepage struct {
	User UserOverview `json:"user"`
	News []News       `json:"news"`
}

type UserInfo struct {
	Name  string `json:"name,omitempty"`
	Role  string `json:"role,omitempty"`
	Group string `json:"group,omitempty"`
}

func ParseNav(r io.Reader) (*UserInfo, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	nav := doc.Find("#m_b1 > .linfo > strong")
	nameNode := nav.Get(0).FirstChild.NextSibling.NextSibling
	name := strings.TrimSpace(strings.ReplaceAll(nameNode.Data, "–", ""))

	roleNode := nameNode.NextSibling
	role := roleNode.FirstChild.Data

	groupNode := roleNode.NextSibling
	group := strings.TrimSpace(groupNode.Data)
	return &UserInfo{name, role, group}, nil
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
