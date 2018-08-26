package main

import (
	"reader/Reader-Golang/APP/ui"
	"reader/Reader-Golang/APP/init"
	"reader/Reader-Golang/APP/Model"
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