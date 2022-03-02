package tools

import (
	"encoding/base64"

	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/utils"
)

func CreateRoomCode(key string, low, high int) int {
	code := utils.NextRoomCode(key, low, high)
	prefix := utils.RangeRand(10, 99)
	return prefix*1000000 + code
}

func IndexOf(arr []int, val int) int {
	for k, v := range arr {
		if v == val {
			return k
		}
	}
	return -1
}

func SliceDel(arr []int, index, num int) []int {
	arrLen := len(arr)
	if index < 0 || index >= arrLen {
		return arr
	}
	lastIndex := index + num
	if lastIndex >= arrLen {
		return arr[:index]
	}
	return append(arr[:index], arr[index+num:]...)
}

func SlicePop(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	return arr[:len(arr)-1]
}
func GetBase64Val(val string) string {
	if bs, err := base64.StdEncoding.DecodeString(val); err == nil {
		return string(bs)
	} else {
		logger.Error(err)
		return val
	}
}
