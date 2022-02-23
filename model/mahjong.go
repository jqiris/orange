package model

import "time"

// GameMahjongRooms [...]
type GameMahjongRooms struct {
	ID         int       `gorm:"primaryKey;column:id;type:int;not null" json:"id"`
	Code       int       `gorm:"column:code;type:int;not null;default:0" json:"code"`         // 房间号
	Sid        string    `gorm:"column:sid;type:varchar(45)" json:"sid"`                      // 服务器id
	Gid        int       `gorm:"column:gid;type:int;not null;default:0" json:"gid"`           // 游戏类型 1001-血流红中
	Tid        int       `gorm:"column:tid;type:int;not null;default:0" json:"tid"`           // 桌号
	Landlord   int       `gorm:"column:landlord;type:int;not null;default:0" json:"landlord"` // 房主id
	State      int8      `gorm:"column:state;type:tinyint;not null;default:0" json:"state"`   // 状态 0-匹配开始 1-游戏开始 2-游戏结束
	Round      int8      `gorm:"column:round;type:tinyint;not null;default:0" json:"round"`   // 第几轮
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`         // 创建时间
	EndTime    time.Time `gorm:"column:end_time;type:datetime" json:"end_time"`               // 游戏结束时间
}
