package initdb

import (
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MYSQLORM *gorm.DB

func init()  {

	connect, err := gorm.Open("mysql", "root:123456@tcp(localhost:3307)/reader?parseTime=true")

	if err != nil {
		fmt.Print(err)
		panic("connect postgres failed ")
	}

	//defer MYSQLORM.Close()

	fmt.Print("Login postgres database success!")

	MYSQLORM = connect;

}