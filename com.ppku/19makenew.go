package main

import "fmt"

func main() {
	/*
			new 为所有类型分配内存，并且初始化零值，返回指针
			make 只能为slice,map,chan 分配内存，并且初始化，返回的是类型
		1. make 分配空间后，会进⾏初始化，new分配的空间被清零
		2. new 分配返回的是指针，即类型 *Type。make 返回引⽤，即 Type；
		3. new 可以分配任意类型的数据
	*/

	num := new(int)
	fmt.Println(*num)

	//a:=make([]int,2,10)
	//b:=make(map[string]int)
	//c:=make(chan int,10)
}
