package constant

import "time"

//缓存时间
const (
	StoreHalfDay = time.Hour * 6
)

//用户相关
const (
	StoreUserMember = "storeUserMember"
)

//分布式数据存储
const (
	StoreCreatingRooms = "storeCreatingRooms"
	StoreRooms         = "storeRooms"
	StoreGames         = "storeGames"
	StoreUserSeats     = "storeUserSeats"
	StoreTotalRooms    = "storeTotalRooms"
	StoreUserLocations = "storeUserLocations"
	StoreLock          = "lock"
)
