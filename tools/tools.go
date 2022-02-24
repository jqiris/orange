package tools

import "github.com/jqiris/kungfu/v2/utils"

func CreateRoomCode(key string, low, high int) int {
	code := utils.NextRoomCode(key, low, high)
	prefix := utils.RangeRand(10, 99)
	return prefix*1000000 + code
}
