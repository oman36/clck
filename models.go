package main

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(
		getConfig().DB.Driver,
		getConfig().DB.Connection,
	)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Urls{})
}

type Urls struct {
	gorm.Model
	Url string
}
