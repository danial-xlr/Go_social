package main

import (
	"fmt"
	"gosocial/profile"	

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	var r profile.Repository
	profile.DB,err = gorm.Open("mysql",profile.DbURL())
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer r.Close()
	profile.DB.AutoMigrate(&profile.Profile{})
	s:=profile.NewService(r)
	err=profile.ListenGRPC(s)
	if err!=nil{
		panic(err)
	}
}