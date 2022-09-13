package main

import "fmt"

func main() {
	/*
		协程调用函数
		go functionname()
	*/

	go showSB(5)
}
func showSB(num int) {
	fmt.Println("你是", num, "号傻逼！")
}
