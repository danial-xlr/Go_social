package main

import (
	"fmt"
	"gosocial/feed"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	var r feed.Repository
	feed.DB,err=gorm.Open("mysql",feed.DbURL())
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer r.Close()
	feed.DB.AutoMigrate(&feed.Feed{})
	s:=feed.NewService(r)	
	profileURL :="localhost:8000" 
	relationURL :="localhost:8002" 
	postURL :="localhost:8001"
	port:=8003
	err=feed.ListenGRPC(s,profileURL,relationURL,postURL,port)
	if err!=nil{
		panic(err)
	}
}