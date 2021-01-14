package helper

import "time"

// 测试 函数执行顺序
func Add(x, y int) (z int) {
	defer func() {
		println(z) // 输出: 203
	}()

	z = x + y
	return z + 200 // 执行顺序: (z = z + 200) -> (call defer) -> (return)
}

// 测试信道阻塞 解决方案
func TestFunc() {
	fc := make(chan func() string, 2)
	//go func() {
	//	println((<-fc)())
	//	println((<-fc)())
	//	println((<-fc)())
	//}()
	go func() {
		fc <- func() string { return "Hello, World!  1" }
	}()
	go func() {
		fc <- func() string { return "Hello, World!  2" }
	}()
	go func() {
		fc <- func() string { return "Hello, World!  3" }
	}()
	time.Sleep(1 * time.Second)
	for {
		select {
		case x := <-fc:
			println(x())
		default:
			return
		}
	}

}
