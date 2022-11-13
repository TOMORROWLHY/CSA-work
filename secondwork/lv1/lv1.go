package main

import "fmt"

type model struct {
}

func (m model) turnN(n, size int) {
	var a []int
	var b int
	for j := 0; j < n; j++ {
		fmt.Scan(&b)
		a = append(a, b)
	}
	fmt.Println(a)
	for i := 0; i < size/2; i++ {
		a[i], a[size-i-1] = a[size-i-1], a[i]
	}
	fmt.Println(a)
}

func main() {
	var n, size int
	fmt.Scan(&n)
	fmt.Scan(&size)
	var num model
	num.turnN(n, size)
}
