package constant

const (
	MahjongRoomCode     = "mahjongRoomCode" //麻将房号低位数
	MahjongRoomCodeLow  = 100001            //最低6位数
	MahjongRoomCodeHigh = 999999            //最高6位数
)

const (
	ACTION_CHUPAI = 1
	ACTION_MOPAI  = 2
	ACTION_PENG   = 3
	ACTION_GANG   = 4
	ACTION_HU     = 5
	ACTION_ZIMO   = 6
)

const (
	ACCOUNT_PRI_KEY = "^&*#$%()@"
	ROOM_PRI_KEY    = "~!@#$(*&^%$&"
	HALL_ADDR       = "127.0.0.1:9001"
	SERVER_ID       = "001"
	CLIENT_IP       = "127.0.0.1"
	CLIENT_PORT     = 10000
	HALL_IP         = "127.0.0.1"
	HALL_PORT       = 9002
	HTTP_TICK_TIME  = 5000
	TOKEN_KEY       = "!@#$%^&"
	HTTP_PORT       = 9003
	Null            = "null"
)
