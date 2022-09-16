package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
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
	var uNumber8 uint8 = 128
	var uNumber16 uint16 = 32768
	var uNumber32 uint32 = math.MaxUint32
	var uNumber64 uint64 = math.MaxUint64
	var uNumber uint = math.MaxUint

	fmt.Printf("unum8的类型是 %T,大小是:%d,数值是：%d \n", uNumber8, unsafe.Sizeof(uNumber8), uNumber8)
	fmt.Printf("unum16的类型是 %T,大小是:%d,数值是：%d \n", uNumber16, unsafe.Sizeof(uNumber16), uNumber16)
	fmt.Printf("unum32的类型是 %T,大小是:%d,数值是：%d \n", uNumber32, unsafe.Sizeof(uNumber32), uNumber32)
	fmt.Printf("unum64的类型是 %T,大小是:%d,数值是：%d \n", uNumber64, unsafe.Sizeof(uNumber64), uNumber64)
	fmt.Printf("unum的类型是 %T,大小是:%d,数值是：%d \n", uNumber, unsafe.Sizeof(uNumber), uNumber)
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
	fmt.Println(a || b) // 短路或
	fmt.Println(a && b) // 短路与
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
			Go语言不存在隐式类型转换，因此所有的类型转换都必须显式的声明
		类型 B 的值 = 类型 B(类型 A 的值)
	*/

	// 初始化一个int32整数
	var a1 int32 = math.MaxInt32 // 1047483647

	// 输出变量的16进制 和10进制的值
	fmt.Printf("int32: 0x%x,%d \n", a1, a1)

	// 将数值转换为16进制,发生数值截断
	a2 := int16(a1)
	fmt.Printf("int16:0x%x,%d \n", a2, a2)

	// 将常量保存为float32
	var s = 44.155
	fmt.Println(int(s))

	/*
			修改字符串

		[]byte string 可以强制类型转换
	*/
	s1 := "localhost:8080"
	strByte := []byte(s1)

	// 下标修改
	strByte[len(s1)-1] = '1'
	fmt.Println(strByte)

	// 强制类型转换
	s2 := string(strByte)
	fmt.Println(s2)

	// 字符串和其他类型转换
	newStr := "1"
	intValue, _ := strconv.Atoi(newStr)
	fmt.Printf("%T,%d \n", intValue, intValue) // int,1

	//int 转 str
	intValues := 1
	strValue := strconv.Itoa(intValues)
	fmt.Printf("%T,%s", strValue, strValue)

	/*
		浮点数与字符串
		4个参数，，1：要转换的浮点数 2. 格式标记（b、e、E、f、g、G) 3. 精度 4. 指定浮点类型（32:float32、64:float64）
		 格式标记：
		 ‘b’ (-ddddp±ddd，⼆进制指数)
		 ‘e’ (-d.dddde±dd，⼗进制指数)
		 ‘E’ (-d.ddddE±dd，⼗进制指数)
		 ‘f’ (-ddd.dddd，没有指数)
		 ‘g’ (‘e’:⼤指数，‘f’:其它情况)
		 ‘G’ (‘E’:⼤指数，‘f’:其它情况)

		 如果格式标记为 ‘e’，‘E’和’f’，则 prec 表示⼩数点后的数字位数
		 如果格式标记为 ‘g’，‘G’，则 prec 表示总的数字位数（整数部分+⼩数部分）
	*/

	string1 := "4.554185"
	f1, _ := strconv.ParseFloat(string1, 32)
	fmt.Printf("%T,%f", f1, f1)
}
