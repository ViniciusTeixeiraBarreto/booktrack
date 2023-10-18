package api

import (
	"booktrack/internal/http/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Teixeira",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	s.fiber.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	s.fiber.Group("/books").Route("/", routes.BookRouter)
	s.fiber.Group("/author").Route("/", routes.AuthorRouter)

	log.Fatal(s.fiber.Listen(":5000"))
}
