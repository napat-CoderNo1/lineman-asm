package main

import (
	"lineman_asm_1/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.CountHandler(r)
	r.Run()
}