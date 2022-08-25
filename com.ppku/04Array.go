package main

import "fmt"

func main() { // 数组操作
	// 声明时没有指定元素值，默认零值
	var arr [5]int
	fmt.Println(arr)

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println(arr)

	// 声明时进行初始化
	var arr1 = [5]int{15, 25, 35, 45, 55}
	fmt.Println(arr1)

	// 段声明方式
	arr2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(arr2)

	// b部分初始化
	arr3 := [5]int{5, 6, 7, 8}
	fmt.Println(arr3)

	//指定索引
	var arr4 = [5]int{1: 100, 4: 200}
	fmt.Println(arr4)

	// 直接赋值
	arr5 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr5)
	fmt.Println("============================")

	// 特别注意：数组的长度是类型的一部分。[3]int与[5]int不是同一个类型。
	arr6 := [3]int{15, 20, 25}
	arr7 := [5]int{15, 20, 25, 30, 35}
	fmt.Printf("type of arr1 %T\n", arr6)
	fmt.Printf("type of arr2 %T\n", arr7)
	fmt.Println("============================")

	// 多维数组 3个2维数组   len(数组名)
	arr8 := [3][2]string{
		{"1", "学java"},
		{"2", "学go"},
		{"3", "学架构"},
	}
	fmt.Println(arr8)
	fmt.Println("数组长度：", len(arr8))
	fmt.Println("============================")

	// 遍历 for range
	for index, value := range arr8 {
		fmt.Printf("第%d个数组：", index)
		for i, v := range value {
			fmt.Printf("第%d元素：", i)
			fmt.Print(v + ",")
		}
		fmt.Println("")
	}

	fmt.Println("============================")
	for _, v := range arr8 {
		fmt.Printf("type:%T,value:%s \n", v, v)
	}

	fmt.Println("==============数组是值类型==============")
	copyArray := arr8
	fmt.Println(copyArray)
}
