package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	db_user := "root"
	db_pass := "root"
	db_host := "localhost"
	db_port := 3303
	db_name := "gorm_two"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db_user, db_pass, db_host, db_port, db_name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB not found")
	}
	DB = db
	return nil
}
