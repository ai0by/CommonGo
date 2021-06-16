package main

import (
	"fmt"
	"time"
)

// 时间结构体
type NowTime struct {
	StrTime        string
	UnixTime       int
	UnixTimeNano   int
	Year           string
	Month          string
	Day            string
	MonthFirstTime int64
}

// 取当前时间日期
func GetNowTime() NowTime {
	current := time.Now()
	oldTime := current.AddDate(0, 0, -7) // 前七天时间
	fmt.Println(oldTime)
	var t NowTime
	t.StrTime = current.Format("2006-01-02 15:04:05")
	t.UnixTime = int(current.Unix())
	t.UnixTimeNano = int(current.UnixNano())
	t.Year = fmt.Sprintf("%d", current.Year())
	t.Month = current.Month().String()
	t.Day = fmt.Sprintf("%d", current.Day())

	// 时间戳转时间
	_ = time.Unix(int64(t.UnixTime), 0)
	_ = time.Unix(int64(t.UnixTime), 0).Format("2006-01-02 15:04:05")
	// 取凌晨0点时间
	loc, _ := time.LoadLocation("Local")                                                                    // 获取时区
	_, _ = time.ParseInLocation("2006-01-02_15:04:05", current.Format("2006-01-02")+"_00:00:00", loc) // 取当日0点时间
	// 月初时间
	year, month, _ := time.Now().Date()
	t.MonthFirstTime = time.Date(year, month, 1, 0, 0, 0, 0, time.Local).Unix()
	return t
}

func main() {
	var timeArr NowTime
	timeArr = GetNowTime()
	println(timeArr.StrTime)

	println("下一天时间戳", timeArr.UnixTime+86400)
	println("上一天时间戳", timeArr.UnixTime-86400)
}
