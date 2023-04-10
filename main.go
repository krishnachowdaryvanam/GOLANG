package main

import (
	"github.com/krishnachowdaryvanam/GOLANG.git/database"
	"github.com/krishnachowdaryvanam/GOLANG.git/routers"

)

func main() {
	database.SetUp()
	engine := routers.Router()
	err:= engine.Run("127.0.0.1:8000")
	if err := nil{
		log.fatal(err)
	}
}
