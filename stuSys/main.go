package main

import (
	"fmt"
	"net/http"
	"stuSys/pkg/setting"
	"stuSys/routers"
)

func main() {
	router := routers.InitRouter()
	router.StaticFS("/",http.Dir("E:/GoProject/tutorsys/stuSys/tutorsys-frontend"))
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}