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
	// 闭包=函数+引用环境
	X := 100
	func() {
		fmt.Println(X)
	}()

	// 闭包实例2
	generator := playerGen("Pp")
	name, hp := generator()
	fmt.Println(name, hp)

	//类型定义
	s2 := format(formatFun, "%d,%d", 10, 20)
	fmt.Println(s2)
}

// 创建⼀个玩家⽣成器, 输⼊名称, 输出⽣成器
func playerGen(name string) func() (string, int) {
	// ⾎量⼀直为150
	hp := 150
	// 返回创建的闭包
	return func() (string, int) {
		// 将变量引⽤到闭包中
		return name, hp
	}
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

/*
在将函数作为参数类型时，可以类型定义
*/
type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}
func formatFun(s string, x, y int) string {
	return fmt.Sprintln(s, x, y)
}
