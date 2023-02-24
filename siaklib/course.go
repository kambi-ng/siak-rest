package siaklib

import (
	"io"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type Course struct {
	Code      string   `json:"code,omitempty"`
	Name      string   `json:"name,omitempty"`
	ClassName string   `json:"class_name,omitempty"`
	Credits   int      `json:"credits,omitempty"`
	Period    string   `json:"period,omitempty"`
	Schedule  []string `json:"schedule,omitempty"`
	Rooms     []string `json:"rooms,omitempty"`
	Lecturers []string `json:"lecturers,omitempty"`
}

func ParseCourseClasses(r io.Reader) ([]Course, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	splitByBr := func(node *html.Node) []string {
		scheduleHtml, _ := doc.FindNodes(node).Html()
		schedule := strings.Split(scheduleHtml, "<br/>")
		return schedule
	}

	courses := make([]Course, 0)
	table := doc.Find("table.box")
	table.Find(".alt, .x").Each(func(i int, s *goquery.Selection) {
		datas := s.Find("td")
		courseColumn := doc.FindNodes(datas.Get(2))

		code := getTextFromNode(datas.Get(1))
		name := courseColumn.Find("a").Text()
		className := courseColumn.Find("span").Text()
		credits, _ := strconv.Atoi(getTextFromNode(datas.Get(3)))

		period := splitByBr(datas.Get(4))[0]
		schedule := splitByBr(datas.Get(5))
		rooms := splitByBr(datas.Get(6))
		lecturers := splitByBr(datas.Get(7))

		courses = append(courses, Course{
			Code:      code,
			Name:      name,
			ClassName: className,
			Credits:   credits,
			Period:    period,
			Schedule:  schedule,
			Rooms:     rooms,
			Lecturers: lecturers,
		})
	})

	return courses, nil
}
