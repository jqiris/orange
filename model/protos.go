package model

import "github.com/jqiris/saki/wall"

type UserLocation struct {
	RoomType  string `protobuf:"bytes,1,opt,name=roomType,proto3" json:"roomType,omitempty"`
	RoomId    string `protobuf:"bytes,2,opt,name=roomId,proto3" json:"roomId,omitempty"`
	SeatIndex int    `protobuf:"varint,3,opt,name=seatIndex,proto3" json:"seatIndex,omitempty"`
}

//base struct
type MjRoomConf struct {
	Type        string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Zimo        int    `protobuf:"varint,2,opt,name=zimo,proto3" json:"zimo,omitempty"`
	Jiangdui    bool   `protobuf:"varint,3,opt,name=jiangdui,proto3" json:"jiangdui,omitempty"`
	Hsz         bool   `protobuf:"varint,4,opt,name=hsz,proto3" json:"hsz,omitempty"`
	Dianganghua int    `protobuf:"varint,5,opt,name=dianganghua,proto3" json:"dianganghua,omitempty"`
	Menqing     bool   `protobuf:"varint,6,opt,name=menqing,proto3" json:"menqing,omitempty"`
	Tiandihu    bool   `protobuf:"varint,7,opt,name=tiandihu,proto3" json:"tiandihu,omitempty"`
	BaseScore   int    `protobuf:"varint,8,opt,name=baseScore,proto3" json:"baseScore,omitempty"`
	MaxFan      int    `protobuf:"varint,9,opt,name=maxFan,proto3" json:"maxFan,omitempty"`
	MaxGames    int    `protobuf:"varint,10,opt,name=maxGames,proto3" json:"maxGames,omitempty"`
	Creator     int64  `protobuf:"varint,11,opt,name=creator,proto3" json:"creator,omitempty"`
}

type MjDissolveData struct {
	EndTime int64  `protobuf:"varint,1,opt,name=endTime,proto3" json:"endTime,omitempty"`
	States  []bool `protobuf:"varint,2,rep,packed,name=states,proto3" json:"states,omitempty"`
}

type MjTingData struct {
	Fan     int    `protobuf:"varint,1,opt,name=fan,proto3" json:"fan,omitempty"`
	Pattern string `protobuf:"bytes,2,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Pai     int    `protobuf:"varint,3,opt,name=pai,proto3" json:"pai,omitempty"`
}

type MjActionData struct {
	Type     string        `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Owner    *MjSeat       `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	State    string        `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	PayTimes int           `protobuf:"varint,4,opt,name=payTimes,proto3" json:"payTimes,omitempty"`
	Ref      *MjActionData `protobuf:"bytes,5,opt,name=ref,proto3" json:"ref,omitempty"`
	Targets  []int         `protobuf:"varint,6,rep,packed,name=targets,proto3" json:"targets,omitempty"`
	Score    int           `protobuf:"varint,7,opt,name=score,proto3" json:"score,omitempty"`
	Iszimo   bool          `protobuf:"varint,8,opt,name=iszimo,proto3" json:"iszimo,omitempty"`
}

type MjQiangGangData struct {
	TurnSeat *MjSeat `protobuf:"bytes,1,opt,name=turnSeat,proto3" json:"turnSeat,omitempty"`
	SeatData *MjSeat `protobuf:"bytes,2,opt,name=seatData,proto3" json:"seatData,omitempty"`
	Pai      int     `protobuf:"varint,3,opt,name=pai,proto3" json:"pai,omitempty"`
	IsValid  bool    `protobuf:"varint,4,opt,name=isValid,proto3" json:"isValid,omitempty"`
}

type MjSeat struct {
	Userid           int64               `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Score            int                 `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	Name             string              `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Ready            bool                `protobuf:"varint,4,opt,name=ready,proto3" json:"ready,omitempty"`
	Seatindex        int                 `protobuf:"varint,5,opt,name=seatindex,proto3" json:"seatindex"`
	NumZiMo          int                 `protobuf:"varint,6,opt,name=numZiMo,proto3" json:"numZiMo,omitempty"`
	NumJiePao        int                 `protobuf:"varint,7,opt,name=numJiePao,proto3" json:"numJiePao,omitempty"`
	NumDianPao       int                 `protobuf:"varint,8,opt,name=numDianPao,proto3" json:"numDianPao,omitempty"`
	NumAnGang        int                 `protobuf:"varint,9,opt,name=numAnGang,proto3" json:"numAnGang,omitempty"`
	NumMingGang      int                 `protobuf:"varint,10,opt,name=numMingGang,proto3" json:"numMingGang,omitempty"`
	NumChaJiao       int                 `protobuf:"varint,11,opt,name=numChaJiao,proto3" json:"numChaJiao,omitempty"`
	Ip               string              `protobuf:"bytes,12,opt,name=ip,proto3" json:"ip,omitempty"`
	Online           bool                `protobuf:"varint,13,opt,name=online,proto3" json:"online,omitempty"`
	Game             *MjGameData         `protobuf:"bytes,14,opt,name=game,proto3" json:"game,omitempty"`
	Holds            []int               `protobuf:"varint,15,rep,packed,name=holds,proto3" json:"holds"`
	Folds            []int               `protobuf:"varint,16,rep,packed,name=folds,proto3" json:"folds"`
	Angangs          []int               `protobuf:"varint,17,rep,packed,name=angangs,proto3" json:"angangs"`
	Diang            []int               `protobuf:"varint,18,rep,packed,name=diang,proto3" json:"diang"`
	Wangangs         []int               `protobuf:"varint,19,rep,packed,name=wangangs,proto3" json:"wangangs"`
	Pengs            []int               `protobuf:"varint,20,rep,packed,name=pengs,proto3" json:"pengs"`
	Que              int                 `protobuf:"varint,21,opt,name=que,proto3" json:"que"`
	Huanpais         []int               `protobuf:"varint,22,rep,packed,name=huanpais,proto3" json:"huanpais"`
	CountMap         map[int]int         `protobuf:"bytes,23,rep,name=countMap,proto3" json:"countMap,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	TingMap          map[int]*MjTingData `protobuf:"bytes,24,rep,name=tingMap,proto3" json:"tingMap,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Pattern          string              `protobuf:"bytes,25,opt,name=pattern,proto3" json:"pattern,omitempty"`
	CanGang          bool                `protobuf:"varint,26,opt,name=canGang,proto3" json:"canGang,omitempty"`
	GangPai          []int               `protobuf:"varint,27,rep,packed,name=gangPai,proto3" json:"gangPai"`
	CanPeng          bool                `protobuf:"varint,28,opt,name=canPeng,proto3" json:"canPeng,omitempty"`
	CanHu            bool                `protobuf:"varint,29,opt,name=canHu,proto3" json:"canHu,omitempty"`
	CanChuPai        bool                `protobuf:"varint,30,opt,name=canChuPai,proto3" json:"canChuPai,omitempty"`
	GuoHuFan         int                 `protobuf:"varint,31,opt,name=guoHuFan,proto3" json:"guoHuFan,omitempty"`
	Hued             bool                `protobuf:"varint,32,opt,name=hued,proto3" json:"hued,omitempty"`
	Iszimo           bool                `protobuf:"varint,33,opt,name=iszimo,proto3" json:"iszimo,omitempty"`
	IsGangHu         bool                `protobuf:"varint,34,opt,name=isGangHu,proto3" json:"isGangHu,omitempty"`
	IsMenQing        bool                `protobuf:"varint,35,opt,name=isMenQing,proto3" json:"isMenQing,omitempty"`
	IsZhongZhang     bool                `protobuf:"varint,36,opt,name=isZhongZhang,proto3" json:"isZhongZhang,omitempty"`
	IsJinGouHu       bool                `protobuf:"varint,37,opt,name=isJinGouHu,proto3" json:"isJinGouHu,omitempty"`
	IsQiangGangHu    bool                `protobuf:"varint,38,opt,name=isQiangGangHu,proto3" json:"isQiangGangHu,omitempty"`
	IsHaiDiHu        bool                `protobuf:"varint,39,opt,name=isHaiDiHu,proto3" json:"isHaiDiHu,omitempty"`
	IsTianHu         bool                `protobuf:"varint,40,opt,name=isTianHu,proto3" json:"isTianHu,omitempty"`
	IsDiHu           bool                `protobuf:"varint,41,opt,name=isDiHu,proto3" json:"isDiHu,omitempty"`
	Qingyise         bool                `protobuf:"varint,42,opt,name=qingyise,proto3" json:"qingyise,omitempty"`
	Actions          []*MjActionData     `protobuf:"bytes,43,rep,name=actions,proto3" json:"actions"`
	Fan              int                 `protobuf:"varint,44,opt,name=fan,proto3" json:"fan,omitempty"`
	LastFangGangSeat int                 `protobuf:"varint,45,opt,name=lastFangGangSeat,proto3" json:"lastFangGangSeat"`
	Numofgen         int                 `protobuf:"varint,46,opt,name=numofgen,proto3" json:"numofgen,omitempty"`
	Diangangs        []int               `protobuf:"varint,47,rep,packed,name=diangangs,proto3" json:"diangangs"`
	Huinfo           []*MjHuData         `protobuf:"bytes,48,rep,name=huinfo,proto3" json:"huinfo"`
}

type MjGameData struct {
	Conf             *MjRoomConf      `protobuf:"bytes,1,opt,name=conf,proto3" json:"conf,omitempty"`
	Uuid             string           `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	GameIndex        int              `protobuf:"varint,3,opt,name=gameIndex,proto3" json:"gameIndex,omitempty"`
	Button           int              `protobuf:"varint,4,opt,name=button,proto3" json:"button,omitempty"`
	GameSeats        []*MjSeat        `protobuf:"bytes,7,rep,name=gameSeats,proto3" json:"gameSeats,omitempty"`
	NumOfQue         int              `protobuf:"varint,8,opt,name=numOfQue,proto3" json:"numOfQue,omitempty"`
	Turn             int              `protobuf:"varint,9,opt,name=turn,proto3" json:"turn,omitempty"`
	ChuPai           int              `protobuf:"varint,10,opt,name=chuPai,proto3" json:"chuPai,omitempty"`
	State            string           `protobuf:"bytes,11,opt,name=state,proto3" json:"state,omitempty"`
	FirstHupai       int              `protobuf:"varint,12,opt,name=firstHupai,proto3" json:"firstHupai,omitempty"`
	Yipaoduoxiang    int              `protobuf:"varint,13,opt,name=yipaoduoxiang,proto3" json:"yipaoduoxiang,omitempty"`
	Fangpaoshumu     int              `protobuf:"varint,14,opt,name=fangpaoshumu,proto3" json:"fangpaoshumu,omitempty"`
	ActionList       []int            `protobuf:"varint,15,rep,packed,name=actionList,proto3" json:"actionList,omitempty"`
	HupaiList        []int            `protobuf:"varint,16,rep,packed,name=hupaiList,proto3" json:"hupaiList,omitempty"`
	ChupaiCnt        int              `protobuf:"varint,17,opt,name=chupaiCnt,proto3" json:"chupaiCnt,omitempty"`
	HuanpaiMethod    int              `protobuf:"varint,18,opt,name=huanpaiMethod,proto3" json:"huanpaiMethod,omitempty"`
	QiangGangContext *MjQiangGangData `protobuf:"bytes,19,opt,name=qiangGangContext,proto3" json:"qiangGangContext,omitempty"`
	LastHuPaiSeat    int              `protobuf:"varint,20,opt,name=lastHuPaiSeat,proto3" json:"lastHuPaiSeat,omitempty"`
	LastFangGangSeat int              `protobuf:"varint,21,opt,name=lastFangGangSeat,proto3" json:"lastFangGangSeat,omitempty"`
	BaseInfoJson     string           `protobuf:"bytes,22,opt,name=baseInfoJson,proto3" json:"baseInfoJson,omitempty"`
	Wall             *wall.Wall       `json:"wall,omitempty"`
}

type MjRoom struct {
	Uuid       string          `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	RoomId     string          `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	GameType   string          `protobuf:"bytes,3,opt,name=game_type,json=gameType,proto3" json:"game_type,omitempty"`
	NumOfGames int             `protobuf:"varint,4,opt,name=numOfGames,proto3" json:"numOfGames,omitempty"`
	CreateTime int64           `protobuf:"varint,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
	NextButton int             `protobuf:"varint,6,opt,name=nextButton,proto3" json:"nextButton,omitempty"`
	Seats      []*MjSeat       `protobuf:"bytes,7,rep,name=seats,proto3" json:"seats,omitempty"`
	Conf       *MjRoomConf     `protobuf:"bytes,8,opt,name=conf,proto3" json:"conf,omitempty"`
	Dr         *MjDissolveData `protobuf:"bytes,9,opt,name=dr,proto3" json:"dr,omitempty"`
}

type MjHuData struct {
	Ishupai       bool   `protobuf:"varint,1,opt,name=ishupai,proto3" json:"ishupai"`
	Pai           int    `protobuf:"varint,2,opt,name=pai,proto3" json:"pai"`
	Action        string `protobuf:"bytes,3,opt,name=action,proto3" json:"action"`
	IsGangHu      bool   `protobuf:"varint,4,opt,name=isGangHu,proto3" json:"isGangHu,omitempty"`
	IsQiangGangHu bool   `protobuf:"varint,5,opt,name=isQiangGangHu,proto3" json:"isQiangGangHu,omitempty"`
	Iszimo        bool   `protobuf:"varint,6,opt,name=iszimo,proto3" json:"iszimo,omitempty"`
	Target        int    `protobuf:"varint,7,opt,name=target,proto3" json:"target,omitempty"`
	Fan           int    `protobuf:"varint,8,opt,name=fan,proto3" json:"fan"`
	Pattern       string `protobuf:"bytes,9,opt,name=pattern,proto3" json:"pattern"`
	IsHaiDiHu     bool   `protobuf:"varint,10,opt,name=isHaiDiHu,proto3" json:"isHaiDiHu,omitempty"`
	IsTianHu      bool   `protobuf:"varint,11,opt,name=isTianHu,proto3" json:"isTianHu,omitempty"`
	IsDiHu        bool   `protobuf:"varint,12,opt,name=isDiHu,proto3" json:"isDiHu,omitempty"`
	Numofgen      int    `protobuf:"varint,13,opt,name=numofgen,proto3" json:"numofgen,omitempty"`
	Index         int    `protobuf:"varint,14,opt,name=index,proto3" json:"index"`
}
