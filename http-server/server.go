package httpserver

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	Router *fiber.App
}

func MakeServer() Server {
	return Server{
		Router: fiber.New(),
	}
}

func (s *Server) Start() {
	allowedOrigins := os.Getenv("ALLOW_ORIGINS")

	if allowedOrigins == "" {
		s.Router.Use(cors.New())
	} else {
		s.Router.Use(cors.New(cors.Config{
			AllowOrigins:     allowedOrigins,
			AllowCredentials: true,
		}))
	}

	s.Router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	s.Router.Post("/login", Login)
	s.Router.Get("/home", Home)
	s.Router.Get("/photo", UserPhoto)

	academicGroup := s.Router.Group("/academic")
	academicGroup.Get("/summary", AcademicSummary)
	academicGroup.Get("/history", AcademicHistory)

	s.Router.Listen(":8080")
}
