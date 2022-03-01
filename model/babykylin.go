package model

// TUser [...]
type TUser struct {
	Userid  uint   `gorm:"primaryKey;column:userid;type:int unsigned;not null" json:"userid"`         // 用户ID
	Account string `gorm:"unique;column:account;type:varchar(64);not null;default:''" json:"account"` // 账号
	Name    string `gorm:"column:name;type:varchar(32)" json:"name"`                                  // 用户昵称
	Sex     int    `gorm:"column:sex;type:int" json:"sex"`
	Headimg string `gorm:"column:headimg;type:varchar(256)" json:"headimg"`
	Lv      int16  `gorm:"column:lv;type:smallint;default:1" json:"lv"`  // 用户等级
	Exp     int    `gorm:"column:exp;type:int;default:0" json:"exp"`     // 用户经验
	Coins   int    `gorm:"column:coins;type:int;default:0" json:"coins"` // 用户金币
	Gems    int    `gorm:"column:gems;type:int;default:0" json:"gems"`   // 用户宝石
	Roomid  string `gorm:"column:roomid;type:varchar(8)" json:"roomid"`
	History string `gorm:"column:history;type:varchar(4096);not null;default:''" json:"history"`
}

// TAccount [...]
type TAccount struct {
	Account  string `gorm:"primaryKey;column:account;type:varchar(255);not null" json:"account"`
	Password string `gorm:"column:password;type:varchar(255);not null" json:"password"`
}
