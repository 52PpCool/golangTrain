package main

import "fmt"

const ( // 常量：不会改变的数据，只能是 bool,整型，字符串型，常量在编译器就确定了
	BEIJING = iota //iota常量计数器,用于⽣成⼀组以相似规则初始化的常量，默认值为0，每一行加1
	SHANGHAI
	SHENZHEN
)
const (
	a = 1
	b
	c = 2
	d
)

func main() {
	// 批量声明的常量除了第一个外其他的初始表达式可以省略
	fmt.Println(a, b, c, d) // "1 1 2 2"
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

	// 方法4 短声明 :=  只能⽤在函数内部 局部变量

	e := 100
	fmt.Printf("e=%d,e的类型为：%T \n", e, e)

	ee := "Hello world"
	fmt.Printf("f=%s,f的类型是:%T \n", ee, ee)

	const length = 10
	fmt.Println("length=", length)
	fmt.Println("BEIJING=", BEIJING)
	fmt.Println("SHANGHAI=", SHANGHAI)
	fmt.Println("SHENZHEN=", SHENZHEN)

	// 批量声明
	//var(
	//	names string
	//	ages int64
	//	length float64
	//)

	// 动态类型 程序运行时系统才能看见的类型
	var number int = 100
	fmt.Println(number)
	number1 := int(100)
	fmt.Printf("number1 Type:%T,data: %v \n", number1, number1)
	number2 := (interface{})(100)
	fmt.Printf("number2 Type:%T,data: %v \n", number2, number2)

	// 类型别名  别名类型只会在代码中存在，编译完成时，不会有别名类型了
	fmt.Println("============类型别名===============")
	type newInt int
	type AliasInt int
	var a1 newInt
	fmt.Printf("type：%T \n", a1)
	var a2 AliasInt
	fmt.Printf("Type:%T \n", a2)
}
