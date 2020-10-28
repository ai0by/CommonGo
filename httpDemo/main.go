package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

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

func HttpPost(url, jccpath string, data []byte) ([]byte, error) {
	body := bytes.NewReader(data)

	request, err := http.NewRequest("POST", url, body)

	if err != nil {
		fmt.Println("请求错误500REQ01:", err.Error())
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("JCC-Path", jccpath)

	resp, err := (&http.Client{}).Do(request)

	code := resp.StatusCode
	if code != 200 {
		return nil, errors.New("请求错误500REQ04 请求状态码错误" + string(code))
	}
	if err != nil {
		fmt.Println("请求错误500REQ02", err)
		return nil, err
	}
	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("请求错误500REQ03", err)
		return nil, err
	}
	return respByte, nil
}

func main(){

}
