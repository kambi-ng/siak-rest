package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	httpserver "github.com/kambi-ng/siak-rest/http-server"
)

//	@title			Siak REST API
//	@version		1.0
//	@description	This an REST API for siak
//	@BasePath		/
func main() {
	if os.Getenv("ENV") != "PRODUCTION" {
		err := godotenv.Load()
		if err != nil {
			fmt.Printf("Error loading .env file, %s\n", err.Error())
			os.Exit(1)
		}
	}

	s := httpserver.MakeServer()
	if err := s.Start(); err != nil {
		fmt.Printf("Can't run server %s\n", err.Error())
		os.Exit(1)
	}
}
