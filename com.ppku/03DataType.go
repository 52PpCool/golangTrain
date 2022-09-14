package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {

	// Integer有符号位整型
	var num8 int8 = 127
	var num16 int16 = 32767
	var num32 int32 = math.MaxInt32
	var num64 int64 = math.MaxInt64
	var num int = math.MaxInt

	fmt.Printf("num8的类型是 %T,大小是:%d,数值是：%d \n", num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16的类型是 %T,大小是:%d,数值是：%d \n", num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32的类型是 %T,大小是:%d,数值是：%d \n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64的类型是 %T,大小是:%d,数值是：%d \n", num8, unsafe.Sizeof(num64), num64)
	fmt.Printf("num的类型是 %T,大小是:%d,数值是：%d \n", num, unsafe.Sizeof(num), num)
	fmt.Println("=======================================")

	// 无符号位整型
	var unum8 uint8 = 128
	var unum16 uint16 = 32768
	var unum32 uint32 = math.MaxUint32
	var unum64 uint64 = math.MaxUint64
	var unum uint = math.MaxUint

	fmt.Printf("unum8的类型是 %T,大小是:%d,数值是：%d \n", unum8, unsafe.Sizeof(unum8), unum8)
	fmt.Printf("unum16的类型是 %T,大小是:%d,数值是：%d \n", unum16, unsafe.Sizeof(unum16), unum16)
	fmt.Printf("unum32的类型是 %T,大小是:%d,数值是：%d \n", unum32, unsafe.Sizeof(unum32), unum32)
	fmt.Printf("unum64的类型是 %T,大小是:%d,数值是：%d \n", unum64, unsafe.Sizeof(unum64), unum64)
	fmt.Printf("unum的类型是 %T,大小是:%d,数值是：%d \n", unum, unsafe.Sizeof(unum), unum)
	fmt.Println("=======================================")

	// 浮点型
	var float32 float32 = math.MaxFloat32
	var float64 float64 = math.MaxFloat64
	fmt.Printf("float32的类型是：%T,数值为%g \n", float32, float32)
	fmt.Printf("float64的类型是：%T,数值为%g \n", float64, float64)
	fmt.Println("=======================================")

	// UTF-8编码
	var x byte = 66
	var y uint8 = 65
	fmt.Printf("x=%c\n", x)
	fmt.Printf("y=%c\n", y)
	fmt.Println("=======================================")

	// 占用字节数
	var x1 int8 = 65
	fmt.Printf("x1=%c \n", x1)
	fmt.Printf("x1=%d \n", x1)
	fmt.Printf("x1=%d个字节 \n", unsafe.Sizeof(x1))
	fmt.Println("=======================================")

	// bool类型
	a := true
	b := false
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println(a || b)
	fmt.Println(a && b)
	fmt.Println("=======================================")

	// 复数构建 complex
	// real 和 imag 函数分别返回复数的实部和虚部
	var z1 complex64 = complex(1, 2)  // 1+2i
	var z2 complex128 = complex(3, 4) // 3+4i

	fmt.Println("z1=", z1)
	fmt.Println("z2=", z2)
	fmt.Println("real(z1)=", real(z1))
	fmt.Println("imag(z1)=", imag(z1))
	fmt.Println("=======================================")

	// 字符
	var g = '牛'
	var g1 = "牛"
	fmt.Println("g=", g)
	fmt.Printf("g=%d,g=%c,类型：%T \n", g, g, g)
	fmt.Println("g1=", g1)
}
