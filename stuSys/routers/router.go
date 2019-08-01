package routers

import (
	"github.com/gin-gonic/gin"
	"stuSys/pkg/setting"
	"stuSys/routers/api"
)

func InitRouter() *gin.Engine{

	gin.SetMode(setting.RunMode)

	r:=gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	apiLogin:=r.Group("/api")
		{
		//登录
		apiLogin.POST("/auth",api.Post)
	}

	return r
}
