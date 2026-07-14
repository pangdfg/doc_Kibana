package main

import (
	"elasticsearch-basic/config"
	"elasticsearch-basic/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.InitElastic()

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8000")
}