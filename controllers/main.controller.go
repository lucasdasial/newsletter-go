package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	Mail "github.com/luccasalves/newsletter-go/modules/mail"
)

func Ping(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, "pong")
}

func SendEmail(ctx *gin.Context) {
	err := Mail.Send()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, "fail")
		return
	}
	ctx.JSON(http.StatusOK, "Sucess")
}
