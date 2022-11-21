package main

import (
	"fmt"
	"sync"
)

var (
	// 逻辑中使用的某个变量
	count1 int
	// 与变量对应的使用互斥锁
	countGuard sync.Mutex
)

func GetCount() int {
	// 锁定
	countGuard.Lock()
	// 在函数退出时解除锁定
	defer countGuard.Unlock()
	return count1
}
func SetCount(c int) {
	countGuard.Lock()
	count1 = c
	countGuard.Unlock()
}
func main() {
	// 可以进行并发安全的设置
	SetCount(1)
	// 可以进行并发安全的获取
	fmt.Println(GetCount())
}
