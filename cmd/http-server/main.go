package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	httpserver "github.com/kambi-ng/siak-rest/http-server"
)

// TODO: This is fake and gay

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file, %s\n", err.Error())
		os.Exit(1)
	}

	s := httpserver.MakeServer()
	if err := s.Start(); err != nil {
		fmt.Printf("Can't run server %s\n", err.Error())
		os.Exit(1)
	}
}
