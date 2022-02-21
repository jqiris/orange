package model

import "time"

// UserMember [...]
type UserMember struct {
	UserID       int       `gorm:"primaryKey;column:user_id;type:int;not null" json:"user_id"`       // 用户id
	Nickname     string    `gorm:"column:nickname;type:varchar(45)" json:"nickname"`                 // 昵称
	Avatar       string    `gorm:"column:avatar;type:varchar(200)" json:"avatar"`                    // 头像
	Gender       int8      `gorm:"column:gender;type:tinyint;default:0" json:"gender"`               // 性别 0-未知 1-男 2-女
	OpenID       string    `gorm:"column:open_id;type:varchar(45)" json:"open_id"`                   // 登陆open_id
	RegTime      time.Time `gorm:"column:reg_time;type:datetime" json:"reg_time"`                    // 注册时间
	LoginTime    time.Time `gorm:"column:login_time;type:datetime" json:"login_time"`                // 登陆时间
	OfflineState int8      `gorm:"column:offline_state;type:tinyint;default:0" json:"offline_state"` // 离线状态 0-在线 1-离线
	OfflineTime  time.Time `gorm:"column:offline_time;type:datetime" json:"offline_time"`            // 离线时间
	IsNew        bool      `gorm:"column:is_new;type:tinyint(1);default:0" json:"is_new"`            // 是否是新用户 0-否 1-是
}
