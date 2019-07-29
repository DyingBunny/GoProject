package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type user struct{
	name string
	pwd string
}

var db *gorm.DB

func connect() {
	Db, err := gorm.Open("mysql", "root:123@/golang?charset=utf8&parseTime=True&loc=Local")
	db=Db
	defer db.Close()
	if err != nil {
		log.Fatal("Connect error:", err)
		return
	}
}

func createTbl(){
	if !db.HasTable(&user{}) {
		db.CreateTable(&user{})
	}
}

func add(_name string,_pwd string){
	db.Create(&user{name:_name,pwd:_pwd})
}

