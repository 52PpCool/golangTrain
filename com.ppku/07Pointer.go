package main

import (
	"fmt"
)

func main() {
	// 通过定义普通变量获取指针
	x := "面向加薪学习"
	ptr := &x
	fmt.Println("x=", x)       // 面向加薪学习
	fmt.Println("*ptr=", *ptr) //  面向加薪学习
	fmt.Println("&x", &x)      //  0xc00004e250
	fmt.Println("ptr", ptr)    //  0xc00004e250
	fmt.Println("=================")

	//  new 先创建指针并分配好内存再写入值
	ptr2 := new(string)
	*ptr2 = "从0到Go语言微服务架构师"
	fmt.Println(ptr2)  //0xc00004e280
	fmt.Println(*ptr2) //从0到Go语言微服务架构师
	fmt.Println(&ptr2) //0xc00000a030
	fmt.Println("=================")

	// 先声明一个变量，再从其他变量获取内存地址指针变量
	x2 := "学Golang"
	var p *string
	p = &x2
	fmt.Println(p)  // 0xc00004e2a0
	fmt.Println(*p) // 学Golang
	fmt.Println("=================")

	// & 可以从一个变量中取得对应的内存地址
	// * 赋值的左边，指该指针指向的变量， 赋值的右边-指总一个指针变量中取得变量值（解引用）

	mystr := "从0到Go语言微服务架构师"
	myint := 1
	mybool := true
	myfloat := 5.555
	fmt.Printf("数值为：%s,类型为:%T \n", mystr, &mystr)
	fmt.Printf("数值为：%d,类型为:%T \n", myint, &myint)
	fmt.Printf("类型为:%T \n", &mybool)
	fmt.Printf("数值为：%f,类型为:%T \n", myfloat, &myfloat)
	fmt.Println("=================")

	x1 := "从0到Go语言微服务架构师"
	var ptr1 *string
	fmt.Println("ptr1 is ", ptr1)
	ptr1 = &x1
	fmt.Println("ptr1 is ", ptr1)
	fmt.Println("ptr1 = ", *ptr1)
	fmt.Println("=================")

	x3 := 99
	p3 := &x3
	fmt.Println(*p3)
	*p3 = 200
	fmt.Println(p3)
	fmt.Println(*p3)
	fmt.Println("=================")

	x4 := [3]int{10, 20, 30}
	x4[0] = 200
	fmt.Println(x4)
	fmt.Println("=================")

	y := [3]int{10, 20, 30}
	//ChangeArray(&y)
	y1 := &y
	y[2] = 300
	y1[0] = 100
	fmt.Println(y)
	fmt.Println("=================")

}
func ChangeArray(value *[3]int) {
	(*value)[0] = 200
}
func Change(value *int) {
	*value = 200
}

