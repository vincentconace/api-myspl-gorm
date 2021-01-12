package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/vincentconace/api-msqly-gorm/Config"
	"github.com/vincentconace/api-msqly-gorm/Models"
	"github.com/vincentconace/api-msqly-gorm/Routes"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
