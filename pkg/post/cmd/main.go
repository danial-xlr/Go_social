package main

import (
	"fmt"
	"gosocial/post"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	var r post.Repository
	post.DB,err=gorm.Open("mysql",post.DbURL())
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer r.Close()
	post.DB.AutoMigrate(&post.Post{})
	s:=post.NewService(r)
	profileURl:="localhost:8000"
	port:=8001
	err=post.ListenGRPC(s,profileURl,port)
	if err!=nil{
		panic(err)
	}
}