package main

import httpserver "github.com/kambi-ng/siak-rest/http-server"

func main() {
	s := httpserver.MakeServer()
	s.Start()
}
