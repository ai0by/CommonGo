package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

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


//读取json文件 返回结果
func ReadJson(filePath string) (result string){
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	buf := bufio.NewReader(file)
	for {
		s, err := buf.ReadString('\n')  // 按行读取
		result += s
		if err != nil {
			if err == io.EOF{   // 判断是否为文件尾
				break
			}else{
				panic(err)
			}
		}
	}
	return result
}

func main() {

	// 测试文件写入
	err := ErrLogWright("./fileDemo/err.log", "test"+"\n")
	if err != nil {
		println("write error")
	}

	// 测试json读写
	result := ReadJson("./fileDemo/1.json")
	println("json内容 ： ",result)
}
