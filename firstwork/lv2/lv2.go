package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//lv2
	var a []int
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		b := rand.Intn(math.MaxInt8)
		a = append(a, b)

	}
	sort.Ints(a)
	fmt.Println(a)
	//lv3在lv2的基础上完成，就不再单开go文件了
	fmt.Println("从小到大排序")
	b := Small_To_Big(a)
	fmt.Println(b)
	fmt.Println("从大到小排序")
	c := Big_To_Small(a)
	fmt.Println(c)

}

// 从小到大排序
func Small_To_Big(a []int) (b []int) {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[n-i-1] {
				a[j], a[n-i-1] = a[n-i-1], a[j]
			}
		}
	}
	return a
}

// 从大到小排序
func Big_To_Small(a []int) (b []int) {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] < a[n-i-1] {
				a[j], a[n-i-1] = a[n-i-1], a[j]
			}
		}
	}
	return a
}
