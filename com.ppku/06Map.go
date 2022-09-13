package main

import "fmt"

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
}
