package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitDatabase(data map[string]any) {
	var err error
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?%v",
		data["username"],
		data["password"],
		data["host"],
		data["port"],
		data["database"],
		"charset=utf8mb4&parseTime=True&loc=Local",
	)
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
}
