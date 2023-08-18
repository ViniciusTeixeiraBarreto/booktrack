package api

import (
	"booktrack/internal/http/routes"
	"booktrack/pkg/api/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	s.server.Use(middleware.SetDatabaseContext)

	router := routes.ConfigRoutes(s.server)

	log.Println("server is running at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}
