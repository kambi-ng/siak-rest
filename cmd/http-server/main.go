package main

import (
	"github.com/joho/godotenv"
	httpserver "github.com/kambi-ng/siak-rest/http-server"
)

func main() {
	godotenv.Load()

	s := httpserver.MakeServer()
	s.Start()
}
