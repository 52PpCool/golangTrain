package main

import (
	"fmt"
	"sync"
)

func main() {
	// map相关操作
	// make(map[keyType]valueType)

	// 创建相同的键值对
	steps := make(map[string]string)
	fmt.Println(steps)

	// 创建不同键值对
	score := make(map[string]int)
	fmt.Println(score)

	// 通过字面值创建map
	var step1 map[int8]string = map[int8]string{
		1: "学java",
		2: "学go",
		3: "学分布式",
	}
	fmt.Println(step1)

	step2 := map[int8]string{
		1: "学java",
		2: "学go",
		3: "学分布式",
	}
	fmt.Println(step2)

	// 添加元素
	step2[4] = "学架构"
	fmt.Println(step2)

	// 更新元素
	step2[4] = "学分布式"
	fmt.Println(step2)

	// 获取元素
	fmt.Println(step2[1])

	// 删除元素
	delete(step2, 4)
	fmt.Println(step2)

	// 判断是否存在
	v, ok := step2[2]
	fmt.Println(v)  //学go
	fmt.Println(ok) //true

	v1, ok := step2[5]
	fmt.Printf("%T,%s \n", v1, v1) //string,
	fmt.Println(ok)                // false

	// for range
	for key, value := range step2 {
		fmt.Println("key:", key, "value:", value)

	}

	// len(map) 长度
	fmt.Println(len(step2))

	// 赋值
	step3 := step2
	fmt.Println(step3)

	// 如果需要一个key对应多个value可以使用切片
	//mp1:=make(map[int][]int)
	//mp2:=make(map[int]*[]int)

	/*
			go 语言中没有清空map的操作，因为垃圾回收比清空高效
			map并发读是线程安全的，并发读写是线程不安全的
		sync.Map 提供了并发安全的使用方式
		store 存储
		load 获取
		Delete 删除

		使用Range配合回调参数进行遍历操作
		回调函数在需要迭代遍历时返回true
		终止迭代遍历时返回false
	*/
	// 只能这样创建
	var scene sync.Map
	// 键值都以interface{}类型进行保存
	scene.Store("green", 97)
	scene.Store("empty", 978)
	scene.Store("eg", 977)
	fmt.Println(scene.Load("green"))
	scene.Delete("green")

	//遍历 提供一个函数在每次遍历元素时都会调用函数返回
	scene.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

}
