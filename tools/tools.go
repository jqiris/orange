package tools

import (
	"encoding/base64"
	"encoding/json"

	"github.com/jqiris/kungfu/v2/logger"
)

func IndexOf(arr []int32, val int32) int {
	for k, v := range arr {
		if v == val {
			return k
		}
	}
	return -1
}

func SliceDel(arr []int32, index, num int) []int32 {
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

func SlicePop(arr []int32) ([]int32, int32) {
	arrLen := len(arr)
	if arrLen == 0 {
		return arr, -1
	}
	return arr[:arrLen-1], arr[arrLen-1]
}
func GetBase64Val(val string) string {
	if bs, err := base64.StdEncoding.DecodeString(val); err == nil {
		return string(bs)
	} else {
		logger.Error(err)
		return val
	}
}

func Stringify(data any) string {
	bs, err := json.Marshal(data)
	if err != nil {
		logger.Error(err)
	}
	return string(bs)
}
