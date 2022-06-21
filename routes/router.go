package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luccasalves/newsletter-go/controllers"
)

func Conf(router *gin.Engine) *gin.Engine {
	router.GET("/", controllers.Ping)
	return router
}
