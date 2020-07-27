package main

import (
	"GoRest/Config"
	"GoRest/Models"
	"GoRest/Router"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mssql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Router.SetupRouter()
	//running
	r.Run()
}
