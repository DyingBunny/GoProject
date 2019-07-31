package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"stuSys/models/table"
)

var Db *gorm.DB

func Connect() {
	db, err := gorm.Open("mysql", "root:123@/golang?charset=utf8&parseTime=True&loc=Local")
	Db=db
	//db.LogMode(true)
	//defer db.Close()
	if err != nil {
		log.Fatal("Connect error:", err)
		return
	}
}

func CreateTbl(){
	if !Db.HasTable(&table.Login{}) {
		Db.CreateTable(&table.Login{})
	}
}

func Init(){
	Connect()
	CreateTbl()
}

func Add(_name string,_pwd string){
	Db.Create(&table.Login{Username: _name,Password:_pwd})
}

func Turnstatus(username interface{},people table.Login,change interface{}){
	Db.Model(&people).Where("Username=?",username).Update("Login",change)
}



