package main

import "fmt"

func main() {
	// 函数可以作为参数传递
	// 高阶函数： 接受一个或者多个函数作为参数 或者 返回值是一个函数 的函数

	// 匿名函数命名
	bookFunc := func() {
		fmt.Println("喜欢go简洁的语法")
	}
	bookFunc()
	fmt.Printf("Type: %T \n", bookFunc) // 类型是func()

	// 先定义函数再调用高阶函数传参处理
	b := func(x, y string) string {
		return x + y
	}
	printRes(b)

	s := showInfo()
	fmt.Println(s("欢喜", "go"))

	// 闭包：函数体访问的变量在函数体外部  匿名函数的特例
	X := 100
	func() {
		fmt.Println(X)
	}()
}

// 高阶函数
func printRes(show func(x, y string) string) {
	fmt.Println(show("喜欢", "go简洁语法"))
}

func showInfo() func(x, y string) string {
	return func(x, y string) string {
		return x + y
	}
}
