package main

import (
	"fmt"
	"gosocial/relation"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	var r relation.Repository
	relation.DB,err=gorm.Open("mysql",relation.DbURL())
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer r.Close()
	relation.DB.AutoMigrate(&relation.Relation{})
	s:=relation.NewService(r)
	profileURl:="localhost:8000"
	port:=8002
	err=relation.ListenGRPC(s,profileURl,port)
	if err!=nil{
		panic(err)
	}
}