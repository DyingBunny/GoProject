package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func connect() {
	db, err := gorm.Open("mysql", "root:123@/?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		log.Fatal("Connect error:", err)
		return
	} else {
		fmt.Println("Connect success")
	}
}
