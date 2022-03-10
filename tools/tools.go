package tools

import (
	"encoding/base64"
	"encoding/json"
	"github.com/jqiris/kungfu/v2/logger"
)

func IndexOf[T comparable](arr []T, val T) int {
	for k, v := range arr {
		if v == val {
			return k
		}
	}
	return -1
}

func SliceDel[T comparable](arr []T, index, num int) []T {
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

func SlicePop[T comparable](arr []T) (bool, []T, T) {
	arrLen := len(arr)
	if arrLen == 0 {
		return false, arr, *new(T)
	}
	return true, arr[:arrLen-1], arr[arrLen-1]
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
