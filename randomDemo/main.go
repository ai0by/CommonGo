package main

import (
	"math/rand"
	"sync"
	"time"
)
var lock sync.Mutex  // 协程锁

func randomString(len int) string {
	lock.Lock()
	defer lock.Unlock()
	r := rand.New(rand.NewSource(time.Now().Unix()))

	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		rand.Seed(time.Now().UnixNano())
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func main(){
	println("random string : ",randomString(5))
}
