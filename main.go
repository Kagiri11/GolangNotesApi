package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:Positive11@tcp(127.0.0.1:3306)/pleya?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Hey we are unable to connect to database due to %v:", err)
		return
	}
	db.AutoMigrate()
}
