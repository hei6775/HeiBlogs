package tool

import (
	"time"
)

//获取这一分钟时间戳
func GetTisMinute() time.Time {
	return time.Now().Truncate(time.Minute)
}

//获取下一分钟时间戳
func GetNextMinute() time.Time {
	return GetTisMinute().Add(1 * time.Minute)
}

//获取当前和下一分钟的时差
func GetDurationToNextMinute() time.Duration {
	nextMinute := GetNextMinute()
	return nextMinute.Sub(time.Now())
}

//获取距离某时刻n天的开始时刻
func GetTheDayBeginTime(src time.Time, days int) (desc time.Time) {
	desc = time.Date(src.Year(), src.Month(), src.Day(), 0, 0, 0, 0, src.Location())
	if days != 0 {
		desc = desc.AddDate(0, 0, days)
	}
	return
}

//得到前一天的0点
func GetYestoday(src time.Time) (desc time.Time) {
	h := src.Hour() + 24
	desc = src.Truncate(time.Hour).Add(time.Duration(-h) * time.Hour)
	return
}

//得到今天的零点时间
func GetToday(src time.Time)(desc time.Time){
	h := src.Hour()
	desc = src.Truncate(time.Hour).Add(time.Duration(-h) * time.Hour)
	return
}

//获取某周周一0点0分0秒时刻
func GetThisWeekMondayBeginTime(src time.Time) (desc time.Time) {
	t := GetYestoday(src)
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	d := time.Duration(-weekday+1) * 24 * time.Hour
	return t.Truncate(time.Hour).Add(d)
}

// 获取整点和半小时
// 例如: 当前是19:20, 则获取到的是19:30, 如果是19:40, 获取到的是20:00
func GetNextHalfHour() time.Time {
	now := time.Now()
	next := now.Truncate(30 * time.Minute)
	if now.After(next) {
		next = next.Add(30 * time.Minute)
	}
	return next
}

const DefaultTimeFormat = "2006-01-02 15:04:05"

//解析时间字符串，返回UTC时间
func ParseTimeStringByFormat(format string, strTime string) (time.Time, error) {
	t, err := time.Parse(format, strTime)
	return t, err
}

func ParseTimeDefaultString(strTime string) (time.Time, error) {
	t, err := time.Parse(DefaultTimeFormat, strTime)
	return t, err
}

//解析时间字符串，返回本地时间
func ParseTimeStringLocationByFormat(format string, strTime string) (time.Time, error) {
	t, err := time.ParseInLocation(format, strTime, time.Local)
	return t, err
}

func ParseTimeDefaultStringLocation(strTime string) (time.Time, error) {
	t, err := time.ParseInLocation(DefaultTimeFormat, strTime, time.Local)
	return t, err
}
