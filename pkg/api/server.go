package api

import (
	"log"
	"web-api/http/routes"
	"web-api/pkg/database"

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
	router := routes.ConfigRoutes(s.server)

	router.Use(func(gin *gin.Context) {
		ctx := gin.Request.Context()
		ctx = database.SetConnection(ctx, nil)
	})

	log.Println("server is running at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}
