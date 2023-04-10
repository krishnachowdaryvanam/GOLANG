package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/krishnachowdaryvanam/GOLANG/database.go"
	"github.com/krishnachowdaryvanam/GOLANG/handlers.go"
)

func Router() *gin.Engine {
	router := gin.Default()
	api := handlers.Handler{
		DB: database.GetDB(),
	}
	router.GET("/books", api.GetBooks)

	return router
}
