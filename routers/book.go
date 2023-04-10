package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/krishnachowdaryvanam/GOLANG.git/database"
	"github.com/krishnachowdaryvanam/GOLANG.git/handlers"
)

func Router() *gin.Engine {
	router := gin.Default()
	api := handlers.Handler{
		DB: database.GetDB(),
	}
	router.GET("/books", api.GetBooks)

	return router
}
