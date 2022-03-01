package database

import (
	"encoding/base64"
	"fmt"

	"github.com/jqiris/orange/model"
)

func GetUserByAccount(account string) (*model.TUser, error) {
	res := &model.TUser{}
	err := db.Where("account = ?", account).First(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetUserById(userId int) (*model.TUser, error) {
	res := &model.TUser{}
	err := db.Where("userid = ?", userId).First(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreateAccount(data *model.TAccount) error {
	return db.Create(data).Error
}

func GetAccount(account string) (*model.TAccount, error) {
	res := &model.TAccount{}
	err := db.Where("account = ?", account).First(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetRoomById(roomId int) (*model.TRoom, error) {
	res := &model.TRoom{}
	err := db.Where("id = ?", roomId).First(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreateRoom(data *model.TRoom) error {
	return db.Create(data).Error
}

func UpdateSeatInfo(roomId, seatIndex, userId int, icon, name string) error {
	colUserId := fmt.Sprintf("user_id%d", seatIndex)
	colIcon := fmt.Sprintf("user_icon%d", seatIndex)
	colUserName := fmt.Sprintf("user_name%d", seatIndex)
	name = base64.StdEncoding.EncodeToString([]byte(name))
	data := map[string]interface{}{
		colUserId:   userId,
		colIcon:     icon,
		colUserName: name,
	}
	return db.Table("t_rooms").Where("id = ?", roomId).Updates(data).Error
}
