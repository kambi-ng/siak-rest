package httpserver

import "github.com/gofiber/fiber/v2"

type Server struct {
	Router *fiber.App
}

func MakeServer() Server {
	return Server{
		Router: fiber.New(),
	}
}

func (s *Server) Start() {
	s.Router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	s.Router.Post("/login", Login)
	s.Router.Get("/home", Home)

	s.Router.Listen(":8080")
}
