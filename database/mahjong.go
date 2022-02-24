package database

import "github.com/jqiris/orange/model"

func GetMahjongRoom(code int) (*model.GameMahjongRooms, error) {
	res := &model.GameMahjongRooms{}
	err := db.Where("code = ?", code).First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreateMahjongRoom(data *model.GameMahjongRooms) error {
	return db.Create(data).Error
}
