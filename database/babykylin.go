package database

import (
	"encoding/base64"
	"fmt"

	"github.com/jqiris/orange/model"
	"github.com/jqiris/orange/tools"
)

func GetUserByAccount(account string) (*model.TUser, error) {
	res := &model.TUser{}
	err := db.Where("account = ?", account).First(res).Error
	if err != nil {
		return nil, err
	}
	res.Name = tools.GetBase64Val(res.Name)
	return res, nil
}

func GetUserById(userId int) (*model.TUser, error) {
	res := &model.TUser{}
	err := db.Where("userid = ?", userId).First(res).Error
	if err != nil {
		return nil, err
	}
	res.Name = tools.GetBase64Val(res.Name)
	return res, nil
}

func UpdateUser(userId int, data map[string]interface{}) error {
	return db.Table("t_users").Where("userid=?", userId).Updates(data).Error
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
	res.UserName0 = tools.GetBase64Val(res.UserName0)
	res.UserName1 = tools.GetBase64Val(res.UserName1)
	res.UserName2 = tools.GetBase64Val(res.UserName2)
	res.UserName3 = tools.GetBase64Val(res.UserName3)
	return res, nil
}

func CreateRoom(data *model.TRoom) error {
	return db.Create(data).Error
}

func DeleteRoom(roomId int) error {
	return db.Where("id=?", roomId).Delete(&model.TRoom{}).Error
}

func UpdateRoom(roomId int, data map[string]interface{}) error {
	return db.Table("t_rooms").Where("id=?", roomId).Updates(data).Error
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

func ArchiveGames(uuid string) error {
	sql := "INSERT INTO t_games_archive(SELECT * FROM t_games WHERE room_uuid = ?)"
	return db.Exec(sql, uuid).Error
}

func CreateGame(data *model.TGame) error {
	return db.Create(data).Error
}

func UpdateGame(uuid string, index int, data map[string]interface{}) error {
	return db.Table("t_games").Where("room_uuid=? and game_index=?", uuid, index).Updates(data).Error
}

func CostGems(userId int, cost int) error {
	sql := "UPDATE t_users SET gems = gems -? where userid = ?"
	return db.Exec(sql, cost, userId).Error
}
