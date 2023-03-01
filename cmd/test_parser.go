package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kambi-ng/siak-rest/siaklib"
)

func main() {
	welcome()
	summary()
	history()
	classes()
	components()
}

func components() {
	f, err := os.Open("html/components/ok.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	res, err := siaklib.ParseCourseDetail(f)
	if err != nil {
		fmt.Printf("Error happened: %v\n", err)
		return
	}
	fmt.Printf("%+v\n", res)

	f, err = os.Open("html/components/missing-table.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	res, err = siaklib.ParseCourseDetail(f)
	if err != nil {
		fmt.Printf("Error happened: %v\n", err)
		return
	}
	fmt.Printf("%+v\n", res)

	f, err = os.Open("html/components/missing-everything.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	res, err = siaklib.ParseCourseDetail(f)
	if err != nil {
		fmt.Printf("Error happened: %v\n", err)
		return
	}
	fmt.Printf("%+v\n", res)
}

func history() {
	f, err := os.Open("html/history.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	res, err := siaklib.ParseAcademicHistoryPage(f)
	if err != nil {
		fmt.Printf("Error happened: %v\n", err)
		return
	}
	fmt.Printf("%+v\n", res)
}

func summary() {
	f, err := os.Open("html/summary.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	res, err := siaklib.ParseAcademicSummaryPage(f)
	if err != nil {
		fmt.Printf("Error happened: %v\n", err)
		return
	}
	fmt.Printf("%+v\n", res)
	fmt.Printf("%+v\n", res.Terms[1].Data)
}

func classes() {
	f, err := os.Open("html/classes.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	res, err := siaklib.ParseCourseClasses(f)
	if err != nil {
		fmt.Printf("Error happened: %v\n", err)
		return
	}
	fmt.Printf("%+v\n", res)
}

func welcome() {
	f, err := os.Open("html/index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	res, err := siaklib.ParseWelcomePage(f)
	if err != nil {
		fmt.Printf("Error happened: %v\n", err)
		return
	}
	fmt.Printf("%+v\n", res)
}
