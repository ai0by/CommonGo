package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var xc sync.WaitGroup

func say(s string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- s // 输出数据
	}
	c <- "end"
	xc.Done() // 手动通知等待结束
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
		time.Sleep(time.Duration(time.Second))
	}
	close(ch)
}

func pump2(ch chan string) {
	for i := 0; ; i++ {
		ch <- strconv.Itoa(i + 5)
		time.Sleep(time.Duration(time.Second))
	}
	close(ch)
}

func suck(ch1 chan int, ch2 chan string) {
	chRate := time.Tick(time.Duration(time.Second * 5)) // 睡眠
	// 循环信道输出
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %s\n", v)
		case <-chRate:
			fmt.Printf("Log log...\n")
		}
	}

}

// 主函数
func main() {
	// 通信
	ch11 := make(chan int)                      // 建立int数据类型信道
	ch22 := make(chan string)                   // 建立String类型信道
	go pump1(ch11)                              // 执行协程 函数 pump1 传递参数 信道 ch1  输出数据
	go pump2(ch22)                              // 执行协程 函数 pump2 传递参数 信道 ch2   输出数据
	go suck(ch11, ch22)                         // 执行协程 函数 suck 传递参数 信道 ch1 ch2  取数据并打印
	time.Sleep(time.Duration(time.Second * 30)) // 睡眠

	//不通信
	xc.Add(2)

	ch1 := make(chan string) // 实例化一个管道
	ch2 := make(chan string) // 实例化一个管道

	go say("Hello", ch1) // 开协程
	go say("World", ch2) // 开协程

	var ch1End string
	var ch2End string
	for {
		if ch1End == "end" && ch2End == "end" {
			break
		}
		ch1End = <-ch1
		ch2End = <-ch2

		if ch1End != "end" {
			println(ch1End) //循环从管道取数据并打印
		} else {
			ch1End = "end"
		}
		if ch2End != "end" {
			println(ch2End) //循环从管道取数据并打印
		} else {
			ch2End = "end"
		}

	}
	xc.Wait() // 等待协程执行完毕结束
}
