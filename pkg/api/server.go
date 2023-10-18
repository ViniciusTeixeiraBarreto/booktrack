package api

import (
	"booktrack/internal/http/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	port  string
	fiber *fiber.App
}

func NewServer() Server {
	return Server{
		port:  "5000",
		fiber: fiber.New(),
	}
}

func (s *Server) RunFiber() {
	s.fiber.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	s.fiber.Group("/books").Route("/", routes.BookRouter)

	log.Fatal(s.fiber.Listen(":5000"))
}
