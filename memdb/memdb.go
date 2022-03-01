package memdb

import (
	"math/rand"
	"time"
)

func GetRandExpire(base time.Duration) time.Duration {
	return base + time.Duration(rand.Intn(3600))*time.Second
}
