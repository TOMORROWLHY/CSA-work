package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg2     sync.WaitGroup
)

func main() {
	wg2.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg2.Wait() //等待goroutine结束
	//最后的运行结果为4，而不会出现上述例子的资源竞争，导致结果出错
	fmt.Println(counter)
}
func incCounter(id int) {
	defer wg2.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1) //安全的对counter加1
		runtime.Gosched()
	}
}
