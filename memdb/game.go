package memdb

import (
	"context"
	"github.com/go-redsync/redsync/v4"
	"strings"

	"fmt"

	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/stores"
	"github.com/jqiris/kungfu/v2/utils"
	"github.com/jqiris/orange/constant"
	"github.com/jqiris/orange/protos"
)

func getStoreKey(key string, arg ...any) (string, string) {
	res := key
	if num := len(arg); num > 0 {
		tmp := res + strings.Repeat("_%v", num)
		res = fmt.Sprintf(tmp, arg...)
	}
	return res, res + "_" + constant.StoreLock
}
func unlock(mutex *redsync.Mutex, ctx context.Context) {
	err := stores.Unlock(mutex, ctx)
	if err != nil {
		logger.Error(err)
	}
}

func GetCreatingRoom(roomId int32) bool {
	key, lock := getStoreKey(constant.StoreCreatingRooms)
	mutex, ctx, err := stores.Lock(lock)
	if err != nil {
		logger.Error(err)
		return true
	}
	defer unlock(mutex, ctx)
	field := utils.Int32ToString(roomId)
	var res bool
	err = stores.HGet(key, field, res)
	if err != nil {
		logger.Error(err)
		return true
	}
	return res
}

func SaveCreatingRoom(roomId int32, data bool) {
	key, lock := getStoreKey(constant.StoreCreatingRooms)
	mutex, ctx, err := stores.Lock(lock)
	if err != nil {
		logger.Error(err)
		return
	}
	defer unlock(mutex, ctx)
	field := utils.Int32ToString(roomId)
	err = stores.HSet(key, field, data)
	if err != nil {
		logger.Error(err)
	}
}

func DelCreatingRoom(roomId int32) {
	key, lock := getStoreKey(constant.StoreCreatingRooms)
	mutex, ctx, err := stores.Lock(lock)
	if err != nil {
		logger.Error(err)
		return
	}
	defer unlock(mutex, ctx)
	field := utils.Int32ToString(roomId)
	err = stores.HDel(key, field)
	if err != nil {
		logger.Error(err)
	}
}

func GetRoom(roomId int32) *protos.MjRoom {
	key, lock := getStoreKey(constant.StoreRooms)
	mutex, ctx, err := stores.Lock(lock)
	if err != nil {
		logger.Error(err)
		return nil
	}
	defer unlock(mutex, ctx)
	field := utils.Int32ToString(roomId)
	res := &protos.MjRoom{}
	err = stores.HGet(key, field, res)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return res
}

func SaveRoom(roomId int32, data *protos.MjRoom) {
	key, lock := getStoreKey(constant.StoreRooms)
	mutex, ctx, err := stores.Lock(lock)
	if err != nil {
		logger.Error(err)
		return
	}
	defer unlock(mutex, ctx)
	field := utils.Int32ToString(roomId)
	err = stores.HSet(key, field, data)
	if err != nil {
		logger.Error(err)
	}
}

func DelRoom(roomId int32) {
	key, lock := getStoreKey(constant.StoreRooms)
	mutex, ctx, err := stores.Lock(lock)
	if err != nil {
		logger.Error(err)
		return
	}
	defer unlock(mutex, ctx)
	field := utils.Int32ToString(roomId)
	err = stores.HDel(key, field)
	if err != nil {
		logger.Error(err)
	}
}

func GetGame(roomId int32) *protos.MjGameData {
	key, lock := getStoreKey(constant.StoreGames)
	mutex, ctx, err := stores.Lock(lock)
	if err != nil {
		logger.Error(err)
		return nil
	}
	defer unlock(mutex, ctx)
	field := utils.Int32ToString(roomId)
	res := &protos.MjGameData{}
	err = stores.HGet(key, field, res)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return res
}

func SaveGame(roomId int32, data *protos.MjGameData) {
	key, lock := getStoreKey(constant.StoreGames)
	mutex, ctx, err := stores.Lock(lock)
	if err != nil {
		logger.Error(err)
		return
	}
	defer unlock(mutex, ctx)
	field := utils.Int32ToString(roomId)
	err = stores.HSet(key, field, data)
	if err != nil {
		logger.Error(err)
	}
}

func DelGame(roomId int32) {
	key, lock := getStoreKey(constant.StoreGames)
	mutex, ctx, err := stores.Lock(lock)
	if err != nil {
		logger.Error(err)
		return
	}
	defer unlock(mutex, ctx)
	field := utils.Int32ToString(roomId)
	err = stores.HDel(key, field)
	if err != nil {
		logger.Error(err)
	}
}

func GetSeat(userId int64) *protos.MjSeat {
	key, lock := getStoreKey(constant.StoreUserSeats)
	mutex, ctx, err := stores.Lock(lock)
	if err != nil {
		logger.Error(err)
		return nil
	}
	defer unlock(mutex, ctx)
	field := utils.Int64ToString(userId)
	res := &protos.MjSeat{}
	err = stores.HGet(key, field, res)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return res
}

func SaveSeat(userId int64, data *protos.MjSeat) {
	key, lock := getStoreKey(constant.StoreUserSeats)
	mutex, ctx, err := stores.Lock(lock)
	if err != nil {
		logger.Error(err)
		return
	}
	defer unlock(mutex, ctx)
	field := utils.Int64ToString(userId)
	err = stores.HSet(key, field, data)
	if err != nil {
		logger.Error(err)
	}
}

func DelSeat(userId int64) {
	key, lock := getStoreKey(constant.StoreUserSeats)
	mutex, ctx, err := stores.Lock(lock)
	if err != nil {
		logger.Error(err)
		return
	}
	defer unlock(mutex, ctx)
	field := utils.Int64ToString(userId)
	err = stores.HDel(key, field)
	if err != nil {
		logger.Error(err)
	}
}
func GetTotalRooms() int {
	key, lock := getStoreKey(constant.StoreTotalRooms)
	mutex, ctx, err := stores.Lock(lock)
	if err != nil {
		logger.Error(err)
		return 0
	}
	defer unlock(mutex, ctx)
	return stores.GetInt(key)
}

func IncrTotalRooms() {
	key, lock := getStoreKey(constant.StoreTotalRooms)
	mutex, ctx, err := stores.Lock(lock)
	if err != nil {
		logger.Error(err)
		return
	}
	defer unlock(mutex, ctx)
	_, err = stores.Incr(key)
	if err != nil {
		logger.Error(err)
	}
}
func DecrTotalRooms() {
	key, lock := getStoreKey(constant.StoreTotalRooms)
	mutex, ctx, err := stores.Lock(lock)
	if err != nil {
		logger.Error(err)
		return
	}
	defer unlock(mutex, ctx)
	_, err = stores.Decr(key)
	if err != nil {
		logger.Error(err)
	}
}

// GetLocation 用户位置信息-全局
func GetLocation(userId int64) *protos.UserLocation {
	key, lock := getStoreKey(constant.StoreUserLocations)
	mutex, ctx, err := stores.Lock(lock)
	if err != nil {
		logger.Error(err)
		return nil
	}
	defer unlock(mutex, ctx)
	field := utils.Int64ToString(userId)
	res := &protos.UserLocation{}
	err = stores.HGet(key, field, res)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return res
}

func SaveLocation(userId int64, data *protos.UserLocation) {
	key, lock := getStoreKey(constant.StoreUserLocations)
	mutex, ctx, err := stores.Lock(lock)
	if err != nil {
		logger.Error(err)
		return
	}
	defer unlock(mutex, ctx)
	field := utils.Int64ToString(userId)
	err = stores.HSet(key, field, data)
	if err != nil {
		logger.Error(err)
	}
}

func DelLocation(userId int64) {
	key, lock := getStoreKey(constant.StoreUserLocations)
	mutex, ctx, err := stores.Lock(lock)
	if err != nil {
		logger.Error(err)
		return
	}
	defer unlock(mutex, ctx)
	field := utils.Int64ToString(userId)
	err = stores.HDel(key, field)
	if err != nil {
		logger.Error(err)
	}
}

func GetUserRoom(userId int64) string {
	if v := GetLocation(userId); v != nil {
		return v.RoomId
	}
	return ""
}

func GetUserSeat(userId int64) int32 {
	if v := GetLocation(userId); v != nil {
		return v.SeatIndex
	}
	return -1
}
