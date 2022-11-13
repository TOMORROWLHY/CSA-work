package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums numbers
	var n int
	fmt.Println("请输入你要创建数组的长度")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Printf("请添加第%d个数： ", i+1)
		fmt.Scan(&a)
		nums = append(nums, a)
	}
	fmt.Print("你创建的数组是：")
	fmt.Println(nums)
	fmt.Println("升序排列后的数组：")
	sort.Sort(nums)
	fmt.Println(nums)
}

type numbers []int

func (n numbers) Len() int {
	return len(n)
}
func (n numbers) Less(i, j int) bool {
	return n[i] < n[j]
}
func (n numbers) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

//func SortInAscendingOrder(a []int) (b []int) {
//	n := len(a)
//	//for i := 0; i < n; i++ {
//		for j := 0; j < n-i-1; j++ {
//	    	if a[j] > a[n-i-1] {
//				a[j], a[n-i-1] = a[n-i-1], a[j]
//			}
//		}
//	}
//	return a
//}
