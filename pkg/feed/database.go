package feed

import(
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

//MySQl Config
func DbURL()string{
	Host:=    "localhost"
  	Port:=     3306
	User:=     "root"
  	Password:= ""
  	DBName:=   "feed"
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		User,
		Password,
		Host,
		Port,
		DBName,
	)
	
}