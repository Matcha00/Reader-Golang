package main

import (
	"reader/APP/Model"

	"reader/APP/init"
	"reader/APP/ui"
)

func main() {

	StartTable()

	defer initdb.MYSQLORM.Close()

	//api_server.New().Start()

	ui.New().Start()


	
}

func StartTable()  {
	initdb.MYSQLORM.AutoMigrate(&Model.Reader{})
}