package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/krishnachowdaryvanam/GOLANG/database.git"
	"github.com/krishnachowdaryvanam/GOLANG/handlers.git"
)

func Router() *gin.Engine {
	router := gin.Default()
	api := handlers.Handler{
		DB: database.GetDB(),
	}
	router.GET("/books", api.GetBooks)

	return router
}
