package tool

import (
	"math/rand"
	"sync"
	"time"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63
)

//rand不是线程/协程安全，必须加锁
var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var rlock = &sync.Mutex{}

//生成一定长度的随机字节数组
func GenRandomByteArray(size int) []byte {
	rlock.Lock()
	defer rlock.Unlock()

	b := make([]byte, size)

	for i, cache, remain := size-1, r.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = r.Int63(), letterIdxMax
		}

		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits //cache置零
		remain--
	}
	return b
}

//生成一定长度的随机字符串
func GenRandomString(size int) string {
	return string(GenRandomByteArray(size))
}

//随机时间 min+到max之间
func RandomTimeDuration(min time.Duration, max time.Duration) time.Duration {
	if min == max {
		return min
	}

	rlock.Lock()
	defer rlock.Unlock()

	if min < max {
		return min + time.Duration(r.Int63n(int64(max-min)))
	}
	return max + time.Duration(r.Int63n(int64(min-max)))
}

//返回一个min max之间的随机int型
func RandomInt(min, max int) int {
	if min == max {
		return min
	}
	rlock.Lock()
	defer rlock.Unlock()

	if min < max {
		return min + r.Intn(max-min)
	}
	return max + r.Intn(min-max)
}

//返回一个随机min max之间的随机int32位
func RandomInt32(min int32, max int32) int32 {
	if min == max {
		return min
	}

	rlock.Lock()
	defer rlock.Unlock()

	if min < max {
		return min + r.Int31n(max-min)
	}
	return max + r.Int31n(min-max)
}

//返回一个min max之间的int64位
func RandomInt64(min, max int64) int64 {
	if min == max {
		return min
	}

	rlock.Lock()
	defer rlock.Unlock()

	if min < max {
		return min + r.Int63n(max-min)
	}
	return max + r.Int63n(min-max)
}
