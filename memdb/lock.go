package memdb

import (
	"time"

	"github.com/jqiris/kungfu/v2/logger"
	"github.com/jqiris/kungfu/v2/stores"
)

func LockStart(key string, args ...int) bool {
	seconds := 3
	if len(args) > 0 {
		seconds = args[0]
	}
	return lockStart(key, seconds)
}

func LockEnd(key string) {
	lockEnd(key)
}

func lockStart(key string, seconds int) bool {
	key = "Lock:" + key
	if stores.Exists(key) {
		return false
	} else {
		if err := stores.Set(key, 1, time.Duration(seconds)*time.Second); err != nil {
			logger.Errorf("lockStart set error: %v", err)
		}
	}
	return true
}

func lockEnd(key string) {
	key = "Lock:" + key
	if err := stores.Del(key); err != nil {
		logger.Error(err)
	}
}
