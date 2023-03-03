package main

import (
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
	godotenv.Load()

	s := httpserver.MakeServer()
	s.Start()
}
