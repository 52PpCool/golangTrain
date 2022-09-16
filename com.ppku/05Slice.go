package main

import "fmt"

func main() {

	// 切片 []Type   数组 [n]Type
	// 方法1 声明切片
	fmt.Println("===========直接空声明 var numberList []int====================")
	var numberList []int
	fmt.Println(numberList)
	fmt.Println(numberList == nil)    //true
	fmt.Println(len(numberList) == 0) //true 判断切片是否为空

	// 方法2 声明一个空切片
	fmt.Println("===========赋值空声明 var numberListEmpty = []int{}=============")
	var numberListEmpty = []int{}
	fmt.Println(numberListEmpty)
	fmt.Println(numberListEmpty == nil)    //false
	fmt.Println(len(numberListEmpty) == 0) //true

	// 方法3 make声明方式  make([]Type,size,cap)
	//指针：是指向第一个切片元素对应的底层数组元素的地址。（切片的第一个元素不一定是数组中第一个元素）
	//长度：切片中的元素个数。
	//容量：从切片的开始位置到底层数据的结尾位置。
	fmt.Println("=============make声明 默认值0 ==================")
	numList := make([]int, 3, 5)
	fmt.Println(numList)
	fmt.Println(numList == nil)    //false
	fmt.Println(len(numList) == 0) //false

	fmt.Println("=============切片 ==================")
	arr := [6]string{"1", "学java", "2", "学go", "3", "学架构"}
	s := arr[:] //[0:len(arr)]
	fmt.Println(s)

	var s1 = arr[0:4] //数组变量[起始位置:结束位置]（切片中不包含结束位置的元素，也就是取值到结束位置-1）
	fmt.Println(s1)
	arr[5] = "学分布式"
	fmt.Println(s1) //[学java 2 学go 3 学分布式]   指针链接实时的
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	//fmt.Println(s1[10]) //  panic: runtime error: index out of range [7] with length 3

	fmt.Println("===============================")
	v := []string{"学java"}
	fmt.Println(v)
	fmt.Println(len(v)) // 4
	fmt.Println(cap(v)) // 6

	//追加一个元素
	v = append(v, "学Go")
	fmt.Println(v)
	fmt.Println(len(v))
	fmt.Println(cap(v))

	//追加2个元素
	fmt.Println(v)
	v = append(v, "学架构", "学分布式")
	fmt.Println(v)

	v = append(v, s1...)
	fmt.Println(v)

	// 切片的复制  copy( destSlice, srcSlice []T) int
	// srcSlice 为数据来源切⽚， destSlice 为复制的⽬标（也就是将 srcSlice 复
	//制到 destSlice）， ⽬标切⽚必须分配过空间且⾜够承载复制的元素个数 ，并且来源和⽬标
	//的 类型必须⼀致 ，copy() 函数的返回值表示实际发⽣复制的元素个数。
	fmt.Println("============切片的复制=============")
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{5, 4, 3}
	i := copy(slice2, slice1)
	fmt.Println("slice2:", slice2, i, len(slice2)) //1,2,3 3 3
	i2 := copy(slice1, slice2)
	fmt.Println("slice1:", slice2, i2, len(slice1)) //1,2,3 3 6

	// 我靠二维切片？
	fmt.Println("===============================")
	f := [][]string{
		{"1", "Go语言极简一本通"},
		{"2", "Go语言微服务架构核心22讲"},
		{"3", "从0到Go语言微服务架构师"},
	}
	fmt.Println(f)
	fmt.Printf("%T \n", f)

	fmt.Println("======================")
}
