package game

type RoomConf struct {
	Type        string `json:"type"`
	Zimo        int    `json:"zimo"`
	Jiangdui    bool   `json:"jiangdui"`
	Hsz         bool   `json:"hsz"`
	Dianganghua int    `json:"dianganghua"`
	Menqing     bool   `json:"menqing"`
	Tiandihu    bool   `json:"tiandihu"`
	BaseScore   int    `json:"baseScore"`
	MaxFan      int    `json:"maxFan"`
	MaxGames    int    `json:"maxGames"`
	Creator     int    `json:"creator"`
}

type DissolveData struct {
	EndTime int64  `json:"endTime"`
	States  []bool `json:"states"`
}

type Room struct {
	Uuid       string        `json:"uuid"`
	Id         int           `json:"id"`
	NumOfGames int           `json:"numOfGames"`
	CreateTime int64         `json:"createTime"`
	NextButton int           `json:"nextButton"`
	Seats      []*Seat       `json:"seats"`
	Conf       RoomConf      `json:"conf"`
	GameMgr    GameMgr       `json:"gameMgr"`
	Dr         *DissolveData `json:"dr"`
}
