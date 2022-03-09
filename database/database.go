package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)
var (
	MjActionArchivePrefix = "mahjong_action_"
	MjRoomArchivePrefix   = "mahjong_room_"
)

func TableMjActionArchive(gameType string) string {
	return MjActionArchivePrefix + gameType + "s"
}

func TableMjRoomArchive(gameType string) string {
	return MjRoomArchivePrefix + gameType + "s"
}
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

func ErrRecordNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}
