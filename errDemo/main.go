package main

import (
	"fmt"
	"os"
	"time"
)

// 错误结构体
type Error struct {
	ErrCode int
	ErrMsg  string
}

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

// 文件写入
func ErrLogWright(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}

//报错增加
func IfError(err error) {
	t := GetNowTime()
	var content string
	content = "[" + t.StrTime + "] "
	if err != nil {
		err = ErrLogWright("./err.log", content+err.Error()+"\n")
	}
	if err != nil {
		fmt.Println("写入到文件出错")
	}
}


func NewError(code int, msg string) *Error {
	return &Error{ErrCode: code, ErrMsg: msg}
}

func (err *Error) Error() string {
	return err.ErrMsg
}

func main() {
	err := NewError(500,"test err")
	IfError(err)
}
