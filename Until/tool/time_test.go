package tool

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()

	fmt.Println(GetYestoday(now))
	fmt.Println(GetThisWeekMondayBeginTime(now))
	fmt.Println(GetNextHalfHour())
	tstring := "2006-08-01 15:15:15"
	fmt.Println(ParseTimeDefaultString(tstring))
}
