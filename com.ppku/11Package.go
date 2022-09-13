package main

import "fmt"

func init() {
	fmt.Println("初始化函数")
}
func main() {
	/*
				包的初始化顺序
				1.包级别的变量
				2. init() 根据解析顺序调用

		import _“fmt” :给fmt包起一个别名，匿名：无法使用当前包的方法，但是会执行当前包的内部的init（）方法
		import aa “fmt”:给fmt包起一个别名，aa，aa.println()来直接调用
		import .“fmt”:将当前fmt包中的全部方法，导入到当前本包的作用中，fmt包中的全部方法可以直接使用API调用，不需要用fmt.API来调用
	*/
}
