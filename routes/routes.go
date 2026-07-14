package routes

import (
	"elasticsearch-basic/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	r.GET("/", handlers.Home)

	r.GET("/init", handlers.InitBooks)

	r.GET("/search", handlers.SearchBooks)
	r.POST("/insert", handlers.InsertDocument)

}