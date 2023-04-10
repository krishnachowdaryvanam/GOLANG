package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/krishnachowdaryvanam/GOLANG/database.go"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) GetBooks(in *gin.Context) {
	books, err := database.GetBooks(h.DB)
	if err != nil {
		log.Panicln(err)
		in.JSON(http.StatusBadRequest, err)
	}
	in.JSON(http.StatusAccepted, books)
}
