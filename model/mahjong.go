package model

// MahjongAction [...]
type MahjongAction struct {
	RoomUUID      string `gorm:"primaryKey;column:room_uuid;type:char(20);not null" json:"room_uuid"`
	GameIndex     int    `gorm:"primaryKey;column:game_index;type:smallint;not null" json:"game_index"`
	GameType      string `gorm:"column:game_type;type:char(8);not null" json:"game_type"` // 游戏类型 xzdd-血战到底
	BaseInfo      string `gorm:"column:base_info;type:varchar(1024);not null" json:"base_info"`
	CreateTime    int64  `gorm:"column:create_time;type:int;not null" json:"create_time"`
	ActionRecords string `gorm:"column:action_records;type:varchar(2048)" json:"action_records"`
	Result        string `gorm:"column:result;type:char(255)" json:"result"`
}

// MahjongRoom [...]
type MahjongRoom struct {
	UUID       string `gorm:"primaryKey;unique;column:uuid;type:char(20);not null" json:"uuid"`
	RoomID     string `gorm:"column:room_id;type:char(8);not null" json:"room_id"`                    // 房间号
	GameType   string `gorm:"column:game_type;type:char(8);not null" json:"game_type"`                // 游戏类型 xzdd-血战到底
	ServerID   string `gorm:"column:server_id;type:varchar(45);not null" json:"server_id"`            // 服务器id
	Landlord   int64  `gorm:"column:landlord;type:int;not null;default:0" json:"landlord"`            // 房主
	BaseInfo   string `gorm:"column:base_info;type:varchar(256);not null;default:0" json:"base_info"` // 配置基本信息
	CreateTime int64  `gorm:"column:create_time;type:int;not null" json:"create_time"`                // 创建时间
	NumOfTurns int    `gorm:"column:num_of_turns;type:int;not null;default:0" json:"num_of_turns"`    // 第几轮
	NextButton int    `gorm:"column:next_button;type:int;not null;default:0" json:"next_button"`      // 下一位
	UserID0    int64  `gorm:"column:user_id0;type:int;not null;default:0" json:"user_id0"`            // 用户1
	UserIcon0  string `gorm:"column:user_icon0;type:varchar(128);not null;default:''" json:"user_icon0"`
	UserName0  string `gorm:"column:user_name0;type:varchar(32);not null;default:''" json:"user_name0"`
	UserScore0 int    `gorm:"column:user_score0;type:int;not null;default:0" json:"user_score0"`
	UserID1    int64  `gorm:"column:user_id1;type:int;not null;default:0" json:"user_id1"`
	UserIcon1  string `gorm:"column:user_icon1;type:varchar(128);not null;default:''" json:"user_icon1"`
	UserName1  string `gorm:"column:user_name1;type:varchar(32);not null;default:''" json:"user_name1"`
	UserScore1 int    `gorm:"column:user_score1;type:int;not null;default:0" json:"user_score1"`
	UserID2    int64  `gorm:"column:user_id2;type:int;not null;default:0" json:"user_id2"`
	UserIcon2  string `gorm:"column:user_icon2;type:varchar(128);not null;default:''" json:"user_icon2"`
	UserName2  string `gorm:"column:user_name2;type:varchar(32);not null;default:''" json:"user_name2"`
	UserScore2 int    `gorm:"column:user_score2;type:int;not null;default:0" json:"user_score2"`
	UserID3    int64  `gorm:"column:user_id3;type:int;not null;default:0" json:"user_id3"`
	UserIcon3  string `gorm:"column:user_icon3;type:varchar(128);not null;default:''" json:"user_icon3"`
	UserName3  string `gorm:"column:user_name3;type:varchar(32);not null;default:''" json:"user_name3"`
	UserScore3 int    `gorm:"column:user_score3;type:int;not null;default:0" json:"user_score3"`
}
