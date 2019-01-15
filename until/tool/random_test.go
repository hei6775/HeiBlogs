package tool

import (
	"fmt"
	"testing"
	"time"
)

func TestRandom(t *testing.T) {
	fmt.Printf("letterBytes:[%v] \n", letterBytes)
	fmt.Printf("letterIdxBits [%v] [%b] \n", letterIdxBits, letterIdxBits)
	fmt.Printf("letterIdxMask [%v] [%b] \n", letterIdxMask, letterIdxMask)
	fmt.Printf("letterIdxMax [%v] [%b] \n", letterIdxMax, letterIdxMax)

	fmt.Printf("-------------------------------------- \n")
	size := 5
	randomArr := GenRandomByteArray(5)
	randomStr := GenRandomString(5)
	randomTime := RandomTimeDuration(time.Second, 80*time.Second)
	randomInt32 := RandomInt32(0, 80)
	randomInt64 := RandomInt64(0, 80)
	randomInt := RandomInt(0, 80)

	fmt.Printf("size[%v],randomArr[%v],randomStr[%v],randomTime[%v] \n"+
		"randomInt[%v],randomInt32[%v],randomInt64[%v]", size, randomArr, randomStr,
		randomTime, randomInt, randomInt32, randomInt64)
}
