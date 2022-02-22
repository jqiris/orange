package database

import "github.com/jqiris/orange/model"

func GetUserMember(userId int) (*model.UserMember, error) {
	user := &model.UserMember{}
	err := db.Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByOpenId(openId string) (*model.UserMember, error) {
	user := &model.UserMember{}
	err := db.Where("open_id= ?", openId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUserMember(userId int, updates map[string]interface{}) error {
	return db.Model(&model.UserMember{}).Updates(updates).Error
}
