package database

import (
	"encoding/base64"
	"fmt"

	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/model"
	"github.com/jqiris/orange/tools"
)

func GetUserByAccount(account string) (*model.UserMember, error) {
	res := &model.UserMember{}
	err := db.Where("account = ?", account).First(res).Error
	if err != nil {
		return nil, err
	}
	res.Name = tools.GetBase64Val(res.Name)
	return res, nil
}

func GetUserById(userId int64) (*model.UserMember, error) {
	res := &model.UserMember{}
	err := db.Where("user_id = ?", userId).First(res).Error
	if err != nil {
		return nil, err
	}
	res.Name = tools.GetBase64Val(res.Name)
	return res, nil
}

func UpdateUser(userId int64, data map[string]any) error {
	return db.Model(&model.UserMember{}).Where("user_id=?", userId).Updates(data).Error
}

func CreateUser(data *model.UserMember) error {
	data.Name = base64.StdEncoding.EncodeToString([]byte(data.Name))
	return db.Create(data).Error
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

func GetMjRoomById(roomId string) (*model.MahjongRoom, error) {
	res := &model.MahjongRoom{}
	err := db.Where("room_id = ?", roomId).First(res).Error
	if err != nil {
		return nil, err
	}
	res.UserName0 = tools.GetBase64Val(res.UserName0)
	res.UserName1 = tools.GetBase64Val(res.UserName1)
	res.UserName2 = tools.GetBase64Val(res.UserName2)
	res.UserName3 = tools.GetBase64Val(res.UserName3)
	return res, nil
}

func CreateMjRoom(data *model.MahjongRoom) error {
	return db.Create(data).Error
}

func ArchiveMjRoom(gameType, roomId string) error {
	t := TableMjRoomArchive(gameType)
	sql := "INSERT INTO " + t + "(SELECT * FROM mahjong_rooms WHERE room_id = ?)"
	return db.Exec(sql, roomId).Error
}

func DeleteMjRoom(roomId string) error {
	return db.Where("room_id=?", roomId).Delete(&model.MahjongRoom{}).Error
}

func UpdateMjRoom(roomId string, data map[string]any) error {
	return db.Model(&model.MahjongRoom{}).Where("room_id=?", roomId).Updates(data).Error
}

func UpdateMjSeatInfo(roomId string, seatIndex int32, userId int64, icon, name string) error {
	colUserId := fmt.Sprintf("user_id%d", seatIndex)
	colIcon := fmt.Sprintf("user_icon%d", seatIndex)
	colUserName := fmt.Sprintf("user_name%d", seatIndex)
	name = base64.StdEncoding.EncodeToString([]byte(name))
	data := map[string]any{
		colUserId:   userId,
		colIcon:     icon,
		colUserName: name,
	}
	return db.Model(&model.MahjongRoom{}).Where("room_id = ?", roomId).Updates(data).Error
}

func ArchiveMjActions(gameType, uuid string) error {
	t := TableMjActionArchive(gameType)
	sql := "INSERT INTO " + t + "(SELECT * FROM mahjong_actions WHERE room_uuid = ?)"
	return db.Exec(sql, uuid).Error
}

func CreateMjAction(data *model.MahjongAction) error {
	return db.Create(data).Error
}

func UpdateMjAction(uuid string, index int32, data map[string]any) error {
	return db.Model(&model.MahjongAction{}).Where("room_uuid=? and game_index=?", uuid, index).Updates(data).Error
}

func CostGems(userId int64, cost int32) error {
	sql := "UPDATE user_members SET gems = gems -? where user_id = ?"
	return db.Exec(sql, cost, userId).Error
}
func GetMjActionArchive(gameType, uuid string) ([]*model.MahjongAction, error) {
	var res []*model.MahjongAction
	t := TableMjActionArchive(gameType)
	err := db.Table(t).Where("room_uuid = ?", uuid).Find(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetMjActionArchiveDetail(gameType, uuid string, index int32) (*model.MahjongAction, error) {
	res := &model.MahjongAction{}
	t := TableMjActionArchive(gameType)
	err := db.Table(t).Where("room_uuid = ? and game_index = ?", uuid, index).First(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetGameNotice(typ, version string) (*model.GameNotice, error) {
	res := &model.GameNotice{}
	do := db.Where("type=?", typ)
	if len(version) > 0 && version != constant.Null {
		do = do.Where("version = ?", version)
	}
	err := do.First(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
