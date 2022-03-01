package database

import "github.com/jqiris/orange/model"

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
