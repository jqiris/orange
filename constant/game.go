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
	CLIENT_PORT     = 10000
	TOKEN_KEY       = "!@#$%^&"
)
