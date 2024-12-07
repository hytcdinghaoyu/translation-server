package utils

import (
	"math/rand"
	"time"
)

var Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandRange(min, max int) int {
	if min == max {
		return max
	}
	if min > max {
		return Rand.Intn(min-max) + max
	}
	return Rand.Intn(max-min) + min
}

func RandRange32(min, max int32) int32 {
	return int32(RandRange(int(min), int(max)))
}

func RandRange64(min, max int64) int64 {
	if min == max {
		return max
	}
	if min > max {
		return Rand.Int63n(min-max) + max
	}
	return Rand.Int63n(max-min) + min
}
