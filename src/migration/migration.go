package main

import (
	"github.com/SND1231/post-service/db"
	"github.com/SND1231/post-service/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := db.Connection()
	defer db.Close()

	db.AutoMigrate(&model.Post{}, &model.Like{})
}
