package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

// RandString 生成随机字符串
func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func StrSplit (s string, sep string) (ret []int32) {
	strArr := strings.Split(s, sep)
	for _, v := range strArr {
		i, _ := strconv.Atoi(v)
		ret = append(ret, int32(i))
	}
	return ret
}

