package main

import "time"

// 时间结构体
type NowTime struct {
	StrTime      string
	UnixTime     int
	UnixTimeNano int
}

// 取当前时间日期
func GetNowTime() NowTime {
	current := time.Now()
	var t NowTime
	t.StrTime = current.Format("2006-01-02 15:04:05")
	t.UnixTime = int(current.Unix())
	t.UnixTimeNano = int(current.UnixNano())
	return t
}

func main() {
	var timeArr NowTime
	timeArr = GetNowTime()
	println(timeArr.StrTime)

	println("下一天时间戳", timeArr.UnixTime + 86400)
	println("上一天时间戳", timeArr.UnixTime - 86400)
}
