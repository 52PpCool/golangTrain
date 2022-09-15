package main

import "fmt"

const (
	BEIJING = 10 * iota //iota常量计数器
	SHANGHAI
	SHENZHEN
)

func main() {
	// var 关键字
	// 方法1 声明一个变量 默认值为0
	var a int
	fmt.Printf("a=%d,a的类型为：%T \n", a, a)

	// 方法2 声明一个变量，并初始化值
	var b int = 100
	fmt.Printf("b=%d,b的类型为：%T \n", b, b)

	// 方法3 初始化去掉值类型，自动匹配
	var c = 100.0
	fmt.Printf("c=%f,c的类型为：%T \n", c, c)

	var cc = "Hello World!"
	fmt.Printf("cc=%s,cc的类型是:%T \n", cc, cc)

	// 方法4 短声明 :=

	e := 100
	fmt.Printf("e=%d,e的类型为：%T \n", e, e)

	ee := "Hello world"
	fmt.Printf("f=%s,f的类型是:%T \n", ee, ee)

	const length = 10
	fmt.Println("length=", length)
	fmt.Println("BEIJING=", BEIJING)
	fmt.Println("SHANGHAI=", SHANGHAI)
	fmt.Println("SHENZHEN=", SHENZHEN)

	// 动态类型 程序运行时系统才能看见的类型
	var number int = 100
	fmt.Println(number)
	number1 := int(100)
	fmt.Printf("number1 Type:%T,data: %v \n", number1, number1)
	number2 := (interface{})(100)
	fmt.Printf("number2 Type:%T,data: %v \n", number2, number2)

}
