package siaklib

import (
	"fmt"
	"io"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type TermData struct {
	SubjectsTaken     int
	CreditsTaken      int
	CreditsPassed     int
	GradePoint        float64
	GradePointAverage float64

	TotalCreditsTaken            int
	TotalCreditsPassed           int
	TotalCreditsEarned           int
	TotalGradePoint              float64
	TotalGradePointAverage       float64
	TotalPassedGradePointAverage float64
}

type AcademicTerm struct {
	Period string
	Term   string
	Data   *TermData
}

type StudentAcademic struct {
	NPM           string
	Name          string
	Year          int
	Major         string
	Tutor         string
	Status        string
	CreditsPassed int
	GradePoints   float64
	GPA           float64
	CreditsEarned int
}

type StudentSummary struct {
	Student        StudentAcademic
	ScoresOverview map[string]int
	Terms          []AcademicTerm
}

func parseSummaryBox(box *goquery.Selection) (*StudentAcademic, error) {
	npm := box.Get(0).FirstChild.Data
	name := box.Get(1).FirstChild.Data
	year, err := strconv.Atoi(box.Get(2).FirstChild.Data)
	if err != nil {
		return nil, err
	}

	program := box.Get(3).FirstChild.Data
	advisor := box.Get(4).FirstChild.Data
	status := box.Get(5).FirstChild.Data
	passedCredits, err := strconv.Atoi(box.Get(6).FirstChild.Data)
	if err != nil {
		return nil, err
	}

	totalGradePoint, err := strconv.ParseFloat(box.Get(7).FirstChild.Data, 32)
	if err != nil {
		return nil, err
	}

	gpa, err := strconv.ParseFloat(box.Get(8).FirstChild.Data, 32)
	if err != nil {
		return nil, err
	}

	earnedCredits, err := strconv.Atoi(box.Get(9).FirstChild.Data)
	if err != nil {
		return nil, err
	}

	return &StudentAcademic{
		npm, name, year, program, advisor, status, passedCredits, totalGradePoint, gpa, earnedCredits,
	}, nil
}

func parseScoresBox(box *goquery.Selection) map[string]int {
	res := make(map[string]int)
	box.Find(".alt, .x").Each(func(i int, s *goquery.Selection) {
		score := s.Children().Get(0).FirstChild.Data
		total, _ := strconv.Atoi(s.Children().Get(1).FirstChild.Data)
		res[score] = total
	})
	return res
}

func parseStatisticBox(box *goquery.Selection) []AcademicTerm {
	var semesterDatas []AcademicTerm
	box.Find(".alt, .x").Each(func(i int, s *goquery.Selection) {
		datas := s.Find("td")
		period := datas.Get(0).FirstChild.Data
		term := datas.Get(1).FirstChild.Data
		if datas.Length() == 3 {
			// Didn't take anything that semester
			semesterDatas = append(semesterDatas, AcademicTerm{period, term, nil})
			return
		}

		subjectsTaken, _ := strconv.Atoi(datas.Get(2).FirstChild.Data)
		creditsTaken, _ := strconv.Atoi(datas.Get(3).FirstChild.Data)
		credisPassed, _ := strconv.Atoi(datas.Get(4).FirstChild.Data)
		gradePoint, _ := strconv.ParseFloat(datas.Get(5).FirstChild.Data, 32)
		termGPA, _ := strconv.ParseFloat(datas.Get(6).FirstChild.Data, 32)

		totalCreditsTaken, _ := strconv.Atoi(datas.Get(7).FirstChild.Data)
		totalCreditsPassed, _ := strconv.Atoi(datas.Get(8).FirstChild.Data)
		totalGradePoint, _ := strconv.ParseFloat(datas.Get(9).FirstChild.Data, 32)
		totalGPA, _ := strconv.ParseFloat(datas.Get(10).FirstChild.Data, 32)
		totalPassedGPA, _ := strconv.ParseFloat(datas.Get(11).FirstChild.Data, 32)
		totalEarnedCredits, _ := strconv.Atoi(datas.Get(12).FirstChild.Data)

		semesterDatas = append(semesterDatas, AcademicTerm{
			period, term,
			&TermData{
				subjectsTaken, creditsTaken, credisPassed, gradePoint, termGPA,
				totalCreditsTaken, totalCreditsPassed, totalEarnedCredits, totalGradePoint, totalGPA, totalPassedGPA,
			},
		})
	})

	return semesterDatas
}

func ParseAcademicSummaryPage(r io.Reader) (*StudentSummary, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}
	boxes := doc.Find("table.box")

	user, err := parseSummaryBox(doc.FindNodes(boxes.Get(0)).Find("td"))
	if err != nil {
		return nil, err
	}
	scores := parseScoresBox(doc.FindNodes(boxes.Get(1)))
	termData := parseStatisticBox(doc.FindNodes(boxes.Get(2)))

	return &StudentSummary{*user, scores, termData}, nil
}

type SubjectScore struct {
	Code       string
	Curriculum string
	Name       string
	Class      string
	Credits    int
	Status     string
	FinalScore string
	FinalIndex string
}

type SemesterScore struct {
	Period   string
	Semester int
	Scores   []SubjectScore
}

func getTextFromNode(elem *html.Node) string {
	for {
		if elem.Type == html.TextNode {
			break
		}

		elem = elem.FirstChild
	}

	return elem.Data
}

func parseSubjectRow(row *goquery.Selection) SubjectScore {
	children := row.Children()

	code := children.Get(1).FirstChild.FirstChild.Data
	curriculum := children.Get(2).FirstChild.Data
	name := children.Get(3).FirstChild.Data
	class := children.Get(4).FirstChild.Data
	credits, _ := strconv.Atoi(children.Get(5).FirstChild.Data)
	status := children.Get(6).FirstChild.Data

	var finalScore, finalIndex string
	if children.Length() == 10 {
		finalScore = getTextFromNode(children.Get(7).FirstChild)
		finalIndex = getTextFromNode(children.Get(8).FirstChild)
	} else {
		finalScore = getTextFromNode(children.Get(7).FirstChild)
		finalIndex = finalScore
	}

	return SubjectScore{
		code, curriculum, name, class, credits, status, finalScore, finalIndex,
	}
}

func ParseAcademicHistoryPage(r io.Reader) (*[]SemesterScore, error) {
	yearTermRegExp := regexp.MustCompile(`Tahun Ajaran (?P<YearS>\d+)\/(?P<YearE>\d+) Term (?P<Semester>\d)`)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	var semesterScores []SemesterScore
	var currentSemester int
	var period string
	var subjectScores []SubjectScore = make([]SubjectScore, 0)

	rows := doc.Find(".box tr")
	for i := 1; i < rows.Length(); i++ {
		elem := doc.FindNodes(rows.Get(i))
		match := yearTermRegExp.FindStringSubmatch(elem.Text())
		if match != nil {
			if i != 1 {
				semesterScores = append(semesterScores, SemesterScore{
					Semester: currentSemester,
					Period:   period,
					Scores:   subjectScores,
				})
			}

			paramsMap := make(map[string]string)
			for i, name := range yearTermRegExp.SubexpNames() {
				if i > 0 && i <= len(match) {
					paramsMap[name] = match[i]
				}
			}

			currentSemester, _ = strconv.Atoi(paramsMap["Semester"])
			period = fmt.Sprintf("%s/%s", paramsMap["YearS"], paramsMap["YearE"])
			subjectScores = make([]SubjectScore, 0)
			continue
		}

		subjectScores = append(subjectScores, parseSubjectRow(elem))
	}

	semesterScores = append(semesterScores, SemesterScore{
		Semester: currentSemester,
		Period:   period,
		Scores:   subjectScores,
	})

	return &semesterScores, nil
}
