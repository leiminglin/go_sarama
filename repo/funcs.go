package repo

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func F1() {
	fmt.Println("f1 func")
}

func SaveMessage(message Message) {
	db.Create(&message)
}

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Message{})
}

type Message struct {
	gorm.Model
	Topic string
	Value string
}
