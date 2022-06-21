package server

import (
	"github.com/gin-gonic/gin"
	"github.com/luccasalves/newsletter-go/routes"
	"log"
)

type Sever struct {
	port   string
	server *gin.Engine
}

func (s Sever) Init() {
	router := routes.Conf(s.server)
	log.Print("Server is running on http://localhost:" + s.port)
	log.Fatal(router.Run(":" + s.port))
}

func CreateServer() Sever {
	return Sever{
		port:   "3333",
		server: gin.Default(),
	}
}
