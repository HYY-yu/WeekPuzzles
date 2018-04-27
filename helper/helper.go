package helper

import (
	"time"
	"log"
	"math"
)

var LOC, _ = time.LoadLocation("Asia/Chongqing")
var TIME_PATTERN = "20060102"

// 计算date对应的时间是否正好是本周。
func IsNowWeek(date time.Time) bool {
	nowSa := FindNowSaturday()
	dateSa := FindRecentSaturday(date)

	return nowSa.Year() == dateSa.Year() && nowSa.Month() == dateSa.Month() && nowSa.Day() == dateSa.Day()
}

func IsEven(a float64) bool {
	return math.Mod(a, 2)/2 == 0
}

func XOR(a, b bool) bool {
	return (a && b) || ((!a) && (!b))
}

func ParseTime(value string) time.Time {
	t, err := time.ParseInLocation(TIME_PATTERN, value, LOC)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

//本周所在的星期六
func FindNowSaturday() time.Time {
	temp := time.Now()
	return FindRecentSaturday(time.Date(temp.Year(), temp.Month(), temp.Day(), 0, 0, 0, 0, LOC))
}

// 根据date找它所在那个周的星期六
func FindRecentSaturday(date time.Time) time.Time {
	deta := 6 - int(date.Weekday())
	if int(date.Weekday()) == 0 {
		deta = -1
	}

	if deta != 0 {
		date = date.AddDate(0, 0, deta)
	}
	return date
}
