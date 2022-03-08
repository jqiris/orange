package model

// TUser [...]
type TUser struct {
	Userid  int64  `gorm:"primaryKey;column:userid;type:int32 unsigned;not null" json:"userid"`       // 用户ID
	Account string `gorm:"unique;column:account;type:varchar(64);not null;default:''" json:"account"` // 账号
	Name    string `gorm:"column:name;type:varchar(32)" json:"name"`                                  // 用户昵称
	Sex     int32  `gorm:"column:sex;type:int32" json:"sex"`
	Headimg string `gorm:"column:headimg;type:varchar(256)" json:"headimg"`
	Lv      int16  `gorm:"column:lv;type:smallint;default:1" json:"lv"`    // 用户等级
	Exp     int32  `gorm:"column:exp;type:int32;default:0" json:"exp"`     // 用户经验
	Coins   int32  `gorm:"column:coins;type:int32;default:0" json:"coins"` // 用户金币
	Gems    int32  `gorm:"column:gems;type:int32;default:0" json:"gems"`   // 用户宝石
	Roomid  int32  `gorm:"column:roomid;type:varchar(8)" json:"roomid"`
	History string `gorm:"column:history;type:varchar(4096);not null;default:''" json:"history"`
}

// TAccount [...]
type TAccount struct {
	Account  string `gorm:"primaryKey;column:account;type:varchar(255);not null" json:"account"`
	Password string `gorm:"column:password;type:varchar(255);not null" json:"password"`
}

// TGame [...]
type TGame struct {
	RoomUUID      string `gorm:"primaryKey;column:room_uuid;type:char(20);not null" json:"room_uuid"`
	GameIndex     int32  `gorm:"primaryKey;column:game_index;type:smallint;not null" json:"game_index"`
	BaseInfo      string `gorm:"column:base_info;type:varchar(1024);not null" json:"base_info"`
	CreateTime    int64  `gorm:"column:create_time;type:int32;not null" json:"create_time"`
	Snapshots     string `gorm:"column:snapshots;type:char(255)" json:"snapshots"`
	ActionRecords string `gorm:"column:action_records;type:varchar(2048)" json:"action_records"`
	Result        string `gorm:"column:result;type:char(255)" json:"result"`
}

// TGuest [...]
type TGuest struct {
	GuestAccount string `gorm:"primaryKey;column:guest_account;type:varchar(255);not null" json:"guest_account"`
}

// TMessage [...]
type TMessage struct {
	Type    string `gorm:"primaryKey;column:type;type:varchar(32);not null" json:"type"`
	Msg     string `gorm:"column:msg;type:varchar(1024);not null" json:"msg"`
	Version string `gorm:"column:version;type:varchar(32);not null" json:"version"`
}

// TRoom [...]
type TRoom struct {
	UUID       string `gorm:"primaryKey;unique;column:uuid;type:char(20);not null" json:"uuid"`
	ID         int32  `gorm:"unique;column:id;type:char(8);not null" json:"id"`
	BaseInfo   string `gorm:"column:base_info;type:varchar(256);not null;default:0" json:"base_info"`
	CreateTime int64  `gorm:"column:create_time;type:int32;not null" json:"create_time"`
	NumOfTurns int32  `gorm:"column:num_of_turns;type:int32;not null;default:0" json:"num_of_turns"`
	NextButton int32  `gorm:"column:next_button;type:int32;not null;default:0" json:"next_button"`
	UserID0    int64  `gorm:"column:user_id0;type:int32;not null;default:0" json:"user_id0"`
	UserIcon0  string `gorm:"column:user_icon0;type:varchar(128);not null;default:''" json:"user_icon0"`
	UserName0  string `gorm:"column:user_name0;type:varchar(32);not null;default:''" json:"user_name0"`
	UserScore0 int32  `gorm:"column:user_score0;type:int32;not null;default:0" json:"user_score0"`
	UserID1    int64  `gorm:"column:user_id1;type:int32;not null;default:0" json:"user_id1"`
	UserIcon1  string `gorm:"column:user_icon1;type:varchar(128);not null;default:''" json:"user_icon1"`
	UserName1  string `gorm:"column:user_name1;type:varchar(32);not null;default:''" json:"user_name1"`
	UserScore1 int32  `gorm:"column:user_score1;type:int32;not null;default:0" json:"user_score1"`
	UserID2    int64  `gorm:"column:user_id2;type:int32;not null;default:0" json:"user_id2"`
	UserIcon2  string `gorm:"column:user_icon2;type:varchar(128);not null;default:''" json:"user_icon2"`
	UserName2  string `gorm:"column:user_name2;type:varchar(32);not null;default:''" json:"user_name2"`
	UserScore2 int32  `gorm:"column:user_score2;type:int32;not null;default:0" json:"user_score2"`
	UserID3    int64  `gorm:"column:user_id3;type:int32;not null;default:0" json:"user_id3"`
	UserIcon3  string `gorm:"column:user_icon3;type:varchar(128);not null;default:''" json:"user_icon3"`
	UserName3  string `gorm:"column:user_name3;type:varchar(32);not null;default:''" json:"user_name3"`
	UserScore3 int32  `gorm:"column:user_score3;type:int32;not null;default:0" json:"user_score3"`
	IP         string `gorm:"column:ip;type:varchar(16)" json:"ip"`
	Port       int32  `gorm:"column:port;type:int32;default:0" json:"port"`
	ServerId   string `gorm:"column:server_id;type:varchar(45)" json:"server_id"`
}

// TGamesArchive [...]
type TGamesArchive struct {
	RoomUUID      string `gorm:"primaryKey;column:room_uuid;type:char(20);not null" json:"room_uuid"`
	GameIndex     int16  `gorm:"primaryKey;column:game_index;type:smallint;not null" json:"game_index"`
	BaseInfo      string `gorm:"column:base_info;type:varchar(1024);not null" json:"base_info"`
	CreateTime    int32  `gorm:"column:create_time;type:int32;not null" json:"create_time"`
	Snapshots     string `gorm:"column:snapshots;type:char(255)" json:"snapshots"`
	ActionRecords string `gorm:"column:action_records;type:varchar(2048)" json:"action_records"`
	Result        string `gorm:"column:result;type:char(255)" json:"result"`
}

type GameConf struct {
	Difen        int32  `json:"difen"`
	Zimo         int32  `json:"zimo"`
	Jiangdui     bool   `json:"jiangdui"`
	Huansanzhang bool   `json:"huansanzhang"`
	Zuidafanshu  int32  `json:"zuidafanshu"`
	Jushuxuanze  int32  `json:"jushuxuanze"`
	Dianganghua  int32  `json:"dianganghua"`
	Menqing      bool   `json:"menqing"`
	Tiandihu     bool   `json:"tiandihu"`
	Type         string `json:"type"`
}
