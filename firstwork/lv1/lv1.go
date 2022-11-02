package main

import (
	"CSA_golong/firstwork/lv1/calculator"
	"fmt"
)

func main() {
	fmt.Println("注意事项：请用空格间隔数字和符号")
	fmt.Println("请输入")
	var (
		a, b, c int
		cmd     string
	)
	fmt.Scan(&a, &cmd, &b)
	switch cmd {
	case "+":
		c = calculator.Add(a, b)
		fmt.Println(c)
	case "-":
		c = calculator.Sub(a, b)
		fmt.Println(c)
	case "*":
		c = calculator.Mul(a, b)
		fmt.Println(c)
	case "/":
		c = calculator.Dis(a, b)
		fmt.Println(c)
	default:
		fmt.Println("error!")
	}
}
