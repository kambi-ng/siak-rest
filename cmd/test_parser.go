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
