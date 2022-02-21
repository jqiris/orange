package cache

import (
	"fmt"

	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/stores"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/database"
	"github.com/jqiris/orange/model"
)

//GetUserMember 获取用户信息
func GetUserMember(userId int) (*model.UserMember, error) {
	var err error
	var info *model.UserMember
	key := fmt.Sprintf("%s_%d", constant.StoreUserMember, userId)
	if stores.Exists(key) {
		if err = stores.Get(key, &info); err != nil {
			return nil, err
		}
		return info, nil
	} else {
		info, err = database.GetUserMember(userId)
		if err != nil {
			return nil, err
		}
		if err = stores.Set(key, info, GetRandExpire(constant.StoreHalfDay)); err != nil {
			logger.Errorf("GetUserMember set error: %v", err)
		}
		return info, nil
	}
}
