package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type crawler struct {
	Name  string
	Money string
}

func mysql(name string, money string) {
	db, err := gorm.Open("mysql", "root:123@/golang?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		log.Fatal("Connect error:", err)
		return
	}
	
	if !db.HasTable(&crawler{}) {
		db.CreateTable(&crawler{})
	}
	
	people := crawler{Name: name, Money: money}
	db.Create(&people)
}
