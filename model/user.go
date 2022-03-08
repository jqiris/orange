package model

import "time"

// UserMember [...]
type UserMember struct {
	UserID    int64     `gorm:"primaryKey;column:user_id;type:int unsigned;not null" json:"user_id"` // 用户id
	Account   string    `gorm:"column:account;type:varchar(64);not null;default:''" json:"account"`
	Name      string    `gorm:"column:name;type:varchar(45)" json:"name"`                             // 昵称
	Headimg   string    `gorm:"column:headimg;type:varchar(256)" json:"headimg"`                      // 头像
	Sex       int8      `gorm:"column:sex;type:tinyint;default:0" json:"sex"`                         // 性别 0-未知 1-男 2-女
	OpenID    string    `gorm:"column:open_id;type:varchar(45)" json:"open_id"`                       // 登陆open_id
	RegTime   time.Time `gorm:"column:reg_time;type:datetime" json:"reg_time"`                        // 注册时间
	LoginTime time.Time `gorm:"column:login_time;type:datetime" json:"login_time"`                    // 登陆时间
	Lv        int16     `gorm:"column:lv;type:smallint;default:1" json:"lv"`                          // 用户等级
	Exp       int       `gorm:"column:exp;type:int;default:0" json:"exp"`                             // 用户经验
	Coins     int       `gorm:"column:coins;type:int;default:0" json:"coins"`                         // 用户金币
	Gems      int       `gorm:"column:gems;type:int;default:0" json:"gems"`                           // 用户宝石
	RoomType  int8      `gorm:"column:room_type;type:tinyint;not null;default:0" json:"room_type"`    // 最后进入房间类型 0-麻将 1-扑克
	RoomID    string    `gorm:"column:room_id;type:char(8)" json:"room_id"`                           // 最后进入房间号
	History   string    `gorm:"column:history;type:varchar(4096);not null;default:''" json:"history"` // 操作记录
}
