package routes

import (
	"lineman_asm_1/controller"

	"github.com/gin-gonic/gin"
)

func CountHandler(router *gin.Engine) {
	router.GET("/covid/summary", controller.CovidSummaryController)
}