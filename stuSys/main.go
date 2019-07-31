package main

import (
	"github.com/gin-gonic/gin"
	"stuSys/models"
	"stuSys/routers/api"
)

func main() {
	router := gin.Default()
	models.Init()
	api.Post(router)
	router.Run() // listen and serve on 0.0.0.0:8080
}