package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stuSys/models"
	"stuSys/models/table"
)

func Post(router *gin.Engine){
	router.StaticFS("/",http.Dir("E:/GoProject/stuSys/view/dist"))
	router.POST("/auth", func(context *gin.Context) {
		var people table.Login
		if context.BindJSON(&people)==nil{
			username:=people.Username
			password:=people.Password
			context.JSON(http.StatusOK,gin.H{
				"status":"OK",
			})
			if models.Check(username,password){
				models.UpdateSta(username)
			}
		}
	})
}

