package main

import (
	"fmt"
	"time"
)

var reload = make(chan string, 1)
var aim = make(chan string, 1)
var fire = make(chan string, 1)

func main() {
	fmt.Println("二营长！把我的意大利炮拿来！开炮！")
	//循环开炮
	for i := 1; i < 101; i++ {
		if i%3 == 1 {
			go reloadTask1()
			go reloadTask2()
			go reloadTask3()
			go reloadTask4()
			go reloadTask5()
			go reloadTask6()
			go reloadTask7()
			go reloadTask8()
			go reloadTask9()
			go reloadTask10()
			a := <-reload
			fmt.Printf("装弹的是：%s\n", a)
			fmt.Println("装弹完毕")
		} else if i%3 == 2 {
			go aimTask1()
			go aimTask2()
			go aimTask3()
			go aimTask4()
			go aimTask5()
			b := <-aim
			fmt.Printf("瞄准的是：%s\n", b)
			fmt.Println("瞄准就绪")
		} else if i%3 == 0 {
			go fireTask1()
			go fireTask2()
			go fireTask3()
			c := <-fire
			fmt.Printf("发射的是：%s\n", c)
			fmt.Println("点火开炮")
		}
		time.Sleep(time.Second)
	}
}

func reloadTask1() {
	//装弹步骤
	reload <- "reloadSolider1"
}
func reloadTask2() {
	//装弹步骤
	reload <- "reloadSolider2"
}
func reloadTask3() {
	//装弹步骤
	reload <- "reloadSolider3"
}
func reloadTask4() {
	//装弹步骤
	reload <- "reloadSolider4"
}
func reloadTask5() {
	//装弹步骤
	reload <- "reloadSolider5"
}
func reloadTask6() {
	//装弹步骤
	reload <- "reloadSolider6"
}
func reloadTask7() {
	//装弹步骤
	reload <- "reloadSolider7"
}
func reloadTask8() {
	//装弹步骤
	reload <- "reloadSolider8"
}
func reloadTask9() {
	//装弹步骤
	reload <- "reloadSolider9"
}
func reloadTask10() {
	//装弹步骤
	reload <- "reloadSolider10"
}
func aimTask1() {
	//瞄准步骤
	aim <- "aimSolider1"
}
func aimTask2() {
	//瞄准步骤
	aim <- "aimSolider2"
}
func aimTask3() {
	//瞄准步骤
	aim <- "aimSolider3"
}
func aimTask4() {
	//瞄准步骤
	aim <- "aimSolider4"
}
func aimTask5() {
	//瞄准步骤
	aim <- "aimSolider5"
}
func fireTask1() {
	//点火步骤
	fire <- "fireSolider1"
}
func fireTask2() {
	//点火步骤
	fire <- "fireSolider2"
}
func fireTask3() {
	//点火步骤
	fire <- "fireSolider3"
}
