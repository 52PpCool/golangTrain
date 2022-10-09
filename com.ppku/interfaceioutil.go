package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("\n------------测试接口---------------")
	testInterFace()

	fmt.Println("\n------------空接口操作---------------")
	emptyInterface()

	fmt.Println("\n------------io操作---------------")
	stringReader()

	fmt.Println("\n------------读取文件---------------")
	readFile()

	fmt.Println("\n------------写文件---------------")
	writeFile()

	fmt.Println("\n------------bufio读写文件---------------")
	write4Bufio() //bufio写文件
	read4Bufio()  //bufio读文件

	fmt.Println("\n------------IoUtil读写文件---------------")
	write4IoUtil() //IoUtil写文件
	read4IoUtil()  //IoUtil读文件

	//实现一个cat命令
	catCommand()
}

func catCommand() {
	flag.Parse() // 解析命令行参数
	if flag.NArg()==0{
		// 如果没有参数默认从标准输入读取内容
		cat(bufio.NewReader(os.Stdin))
	}
	// 依次读取指定文件的内容并答应到终端
	for i := 0; i <flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout,"reading from %s failed, err:%v \n", flag.Arg(i), err)
			continue 
		}
		cat(bufio.NewReader(f))
	}
}

func cat(reader *bufio.Reader) {
	for{
		buf, err := reader.ReadBytes('\n')//这里是字符
		if err == io.EOF {
			break
		} 
		fmt.Fprintf(os.Stdout,"%s",buf)
	}
}

func read4IoUtil() {
	file, err := ioutil.ReadFile("./yyy.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(file))
}

func write4IoUtil() {
	err := ioutil.WriteFile("./yyy.txt", []byte("go go go"), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func read4Bufio() {
	file, err := os.Open("./xxx01.txt")
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != io.EOF {
			break
		}
		if err != nil {
			return
		}
		fmt.Println(string(line))
	}
}

func write4Bufio() {
	// 参数2：打开模式，所有模式d都在上面
	// 参数3是权限控制
	// w写 r读 x执行   w  2   r  4   x  1
	//特殊权限位，拥有者位，同组用户位，其余用户位
	file, err := os.OpenFile("./xxx01.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	// 获取writer对象
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("Hello \n")
	}
	// 刷新缓冲区，强制写出
	writer.Flush()

}

func writeFile() {
	file, err := os.Create("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for i := 0; i < 5; i++ {
		file.WriteString("abcdefghijklmnopqrstuvwxyz\n")
		file.Write([]byte("cd 2222\n"))
	}
}

func readFile() {
	file, err := os.Open("./xxx.txt")
	if err != nil {
		fmt.Println("open file err")
		return
	}
	defer file.Close()
	// 定义接收文件读取的字节数组
	var buf [128]byte
	var content []byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			fmt.Println("read file err", err)
			return
		}
		fmt.Println("read length:", n)
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))
}

/*
*
Reader接口的定义，Read()方法用于读取数据

	type Reader interface{
		Read(p []byte) (n int,err error)
	}

io.Reader 表示一个读取器它将资源读取到缓冲区，数据可以被流式传输和使用
对于用作读取器的类型必须实现Read(p []byte)方法
通过string.NewReader(string) 创建字符串读取器，流式按字节读取
*/
func stringReader() {
	reader := strings.NewReader("zhangsan test")
	// 每次读取四个字节
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				log.Printf("读完了，eof错误：%d", n)
				break
			}
			log.Printf("其他错误，%v", err)
			os.Exit(2)
		}
		log.Printf("读取到的字节数为%d,内容为:%v", n, string(p[:n]))
	}
}

func testInterFace() {
	// 实例化file
	f := new(file)
	// 声明DataWrite接口
	var writer DataWriter
	// *file类型
	writer = f
	// 使用DataWriter接口进行写入
	writer.WriteData("data")
	fmt.Println("\n============")
	var x Sayer
	var y Mover

	var a = dog{name: "旺财"}
	x = a
	y = a
	x.say()
	y.move()

	var b = car{brand: "保时捷"}
	y = b
	y.move()

}

func emptyInterface() {
	//定义一个空接口x
	var x interface{}
	s := "go语言真好玩"
	x = s
	fmt.Printf("type:%T,value:%v \n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T,value:%v \n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T,value:%v \n", x, x)

	//空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "李白"
	studentInfo["age"] = 19
	studentInfo["married"] = false
	fmt.Println(studentInfo)

	var xx interface{}
	xx = "go语言有点奇怪"
	v, ok := xx.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}

// DataWriter 数据写入器
type DataWriter interface {
	// WriteData 写入数据
	WriteData(Data interface{}) error
	// CanWrite 能否写入
	CanWrite() bool
}

// 文件结构，实现DataWriter
type file struct {
}

// WriteData 实现WriteData方法
func (d *file) WriteData(data interface{}) error {
	// 模拟写入数据
	fmt.Println("WriteData", data)
	return nil
}

// CanWrite 实现CanWrite方法
func (d *file) CanWrite() bool {
	// 模拟写入数据
	fmt.Println("CanWrite....")
	return false
}

// Sayer Sayer接口
type Sayer interface {
	say()
}

// Mover mover接口
type Mover interface {
	move()
}

// 一个类型可以同时实现多个接口，接口间独立，不知道对方实现
type dog struct {
	name string
}
type car struct {
	brand string
}

// dog实现say接口
func (d dog) say() {
	fmt.Printf("%s会叫汪汪汪 \n", d.name)

}

// dog实现move接口
func (d dog) move() {
	fmt.Printf("%s会叫跑 \n", d.name)
}

// car实现move接口
func (d car) move() {
	fmt.Printf("%s速度70迈 \n", d.brand)
}
