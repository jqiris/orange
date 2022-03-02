package game

type Seat struct {
	UserId      int    `json:"userid"`
	Score       int    `json:"score"`
	Name        string `json:"name"`
	Ready       bool   `json:"ready"`
	SeatIndex   int    `json:"seatindex"`
	NumZiMo     int    `json:"numZiMo"`
	NumJiePao   int    `json:"numJiePao"`
	NumDianPao  int    `json:"numDianPao"`
	NumAnGang   int    `json:"numAnGang"`
	NumMingGang int    `json:"numMingGang"`
	NumChaJiao  int    `json:"numChaJiao"`
	Ip          string `json:"ip"`
	Online      bool   `json:"online"`
	//游戏相关
	Game             *GameData `json:"game"`
	Holds            []int     `json:"holds"`
	Folds            []int     `json:"folds"`
	Angangs          []int     `json:"angangs"`
	Diangangs        []int     `json:"diang"`
	Wangangs         []int     `json:"wangangs"`
	Pengs            []int     `json:"pengs"`
	Que              int
	Huanpais         []int             `json:"huanpais"`
	CountMap         map[int]int       `json:"countMap"`
	TingMap          map[int]*TingData `json:"tingMap"`
	Pattern          string            `json:"pattern"`
	CanGang          bool              `json:"canGang"`
	GangPai          []int             `json:"gangPai"`
	CanPeng          bool              `json:"canPeng"`
	CanHu            bool              `json:"canHu"`
	CanChuPai        bool              `json:"canChuPai"`
	GuoHuFan         int               `json:"guoHuFan"`
	Hued             bool              `json:"hued"`
	Iszimo           bool              `json:"iszimo"`
	IsGangHu         bool              `json:"isGangHu"`
	Actions          []int             `json:"actions"`
	Fan              int               `json:"fan"`
	LastFangGangSeat int               `json:"lastFangGangSeat"`
}
