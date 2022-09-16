package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	/*
			数据类型 1B=8bit
			bool
			string
			int(随系统，64位8字节，32位4字节) 有符号位整数   -2^(n-1) 到 2^(n-1)-1
				int8 int16 int32(rune,代表一个Unicode码) int64
			uint 无符号位整数       0 到 2^n-1
				uint8(byte) uint16 uint32 uint64
				uintptr  ⽆符号的整数类型  没有指定具体的 bit ⼤⼩但是⾜以容纳指针
						型只有在 底层编程 时才需要，特别是Go语⾔和C语⾔函数库或操作系统接⼝相交互的地⽅
			float32  约1.4e-45 到 约3.4e38
			float64  约4.9e-324 到 约1.8e308

			complex64 complex128  复数 x+yi

		当⼀个变量被声明之后，系统⾃动赋予它该类型的零值：
		int 为 0 ， float 为 0.0 ， bool 为 false ， string 为空字符串 ， 指针为 nil
	*/

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
	// 在申明时可以只写整数或者小数部分
	var e = .4555
	var f = 1.
	var planck = 6.62606957e-34
	fmt.Println(e, f, planck)
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

	/*
		字符  有两种类型
		uint8(byte) 代表ASCII码的一个字符  定义 128 个字符，由码位 0 – 127 标识。它涵盖英⽂字⺟，拉丁数字和其他⼀些字符
		int32(rune) 代表一个UTF-8( Unicode)字符
		在书写 Unicode 字符时，需要在 16 进制数之前加上前缀 \u 或者 \U 。如果需要使
		⽤到 4 字节，则使⽤ \u 前缀，如果需要使⽤到 8 个字节，则使⽤ \U 前缀

		判断是否为字⺟：unicode.IsLetter(ch)
		判断是否为数字：unicode.IsDigit(ch)
		判断是否为空⽩符号：unicode.IsSpace(ch)
	*/

	var g = '牛'
	var g1 = "牛"
	fmt.Println("g=", g)
	fmt.Printf("g=%d,g=%c,类型：%T \n", g, g, g)
	fmt.Println("g1=", g1)

	/*
							字符串类型
						go语⾔从底层就⽀持UTF-8编码
					该编码对占⽤字节⻓度的不定性，在Go语⾔中字符串也可能根据需要占⽤
					1 ⾄ 4 个字节，这与其它编程语⾔不同
				当字符为 ASCII 码表上的字符时则占⽤ 1 个字节
			字符串中可以使⽤转义字符来实现换⾏、缩进等效果，常⽤的转义字符包括:
			1. \n： 换⾏符
			2. \r： 回⻋符
			3. \t： tab 键
			4. \u 或 \U： Unicode 字符
			5. \：反斜杠⾃身
		如果使⽤``反引号，会被原样进⾏赋值和输出
	*/
	var str = "\tgo语言\n牛逼！"
	fmt.Println(str)
	fmt.Println(`\tgo语言\n牛逼！`) //原样输出

	var str01 string = "hello"
	var str02 [5]byte = [5]byte{104, 101, 108, 108, 111}
	fmt.Printf("myStr01: %s\n", str01)
	fmt.Printf("myStr02: %s", str02)

	/*
			对于纯 ASCII 码字符串可以通过标准索引法来获取，在⽅括号 [] 内写⼊索引
			//索引从 0 开始计数

			字符串 str 的第 1 个字节：str[0]
			第 i 个字节：str[i - 1]
			最后 1 个字节：str[len(str)-1

		使⽤ len() 函数获取长度
		Unicode字符串⻓度使⽤ utf8.RuneCountInString() 函数
	*/
	str3 := "hello"
	str4 := "你好"
	fmt.Println(len(str3))                    // 1个字⺟占1个字节
	fmt.Println(len(str4))                    // 1个中⽂占3个字节，go从底层⽀持utf8
	fmt.Println(utf8.RuneCountInString(str4)) // 2

	// 字符串拼接符"+" 或者函数writeString()
	str5 := str3 + str4
	fmt.Println(str5)
	//  writeString()
	var stringBuilder bytes.Buffer
	stringBuilder.WriteString(str3)
	stringBuilder.WriteString(str4)
	fmt.Println(stringBuilder.String())

	// 提取某个值
	var str6 = "hello,伟大的go语言"
	fmt.Println(string([]rune(str6)[6]))

	// 遍历  unicode使用for range ascII 使用for range 或者for循环遍历
	/*
		%c 单⼀字符
		%T 动态类型
		%v 本来值的输出
		%+v 字段名+值打印
		%d ⼗进制打印数字
		%p 指针，⼗六进制
		%f 浮点数
		%b ⼆进制
		%s string
	*/
	var str7 = "hello"
	var str8 = "hello,伟大的go语言"
	for i := 0; i < len(str7); i++ {
		fmt.Printf("ascII:%c,%d", str7[i], str7[i])
	}

	for _, s := range str8 {
		fmt.Printf("unicode:%c,%d", s, s)
	}

	/*
			字符串查找

		strings.Index()： 正向搜索⼦字符串
		strings.LastIndex()：反向搜索⼦字符串
	*/
	tracer := "张三来了,张三bye bye"

	//正向
	com := strings.Index(tracer, ",")
	fmt.Println(",位置：", com)
	fmt.Println(tracer[com+1:]) //  张三bye bye

	add := strings.Index(tracer, "+")
	fmt.Println("+所在的位置：", add) // -1

	pos := strings.Index(tracer[com:], "张三")
	fmt.Println(pos) //1

	fmt.Println(com, pos, tracer[5+pos])

	/*
		类型转换
	*/
	
}
