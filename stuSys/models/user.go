package models

import "stuSys/models/table"

func Find(username string) (table.Login,bool){
	var people table.Login
	Db.First(&people,"Username=?",username)
	if people.Username==""{
		return people,false
	}
	return people,true
}
func Check(username string,password string) bool{
	people,judge:=Find(username)
	if judge==false{
		return false
	}
	if people.Password==password{
		return true
	}else{
		return false
	}
}
func UpdateSta(username interface{})bool{
	var people table.Login
	Turnstatus(username,people,"Yes")
	return true
}
