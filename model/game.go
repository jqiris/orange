package model

// GameNotice [...]
type GameNotice struct {
	Type    string `gorm:"primaryKey;column:type;type:varchar(32);not null" json:"type"`
	Msg     string `gorm:"column:msg;type:varchar(1024);not null" json:"msg"`
	Version string `gorm:"column:version;type:varchar(32);not null" json:"version"`
}
