package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func multipleBranch() {
	/*
		if 条件一{
			逻辑一
		}else if 条件二{
			逻辑二
		}else {
			逻辑三
		}
	*/
	score := 88
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else {
		fmt.Println("不及格")
	}
}

func singleBrand() {
	score := 88
	if score >= 60 {
		fmt.Println("通过")
	}
}

func doubleBrand() {
	score := 88
	if score >= 60 {
		fmt.Println("通过")
	} else {
		fmt.Println("不通过")
	}
}

// 高级条件分支判断
func advanceBrand() {
	if score := 88; score >= 60 {
		fmt.Println("及格")
	}
}

// 特殊写法
func specialBrand() {
	if a := 10; a > 5 {
		fmt.Println(a)
		return
	}
}

func choose() {

	/*
		switch 表达式{
		case 表达值1：
			逻辑一
		case 表达值2：
			逻辑二
		default:
			默认逻辑
		}
	*/
	grade := "B"
	switch grade {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
	case "C":
		fmt.Println("及格")
	case "D":
		fmt.Println("不及格")
	default:
		fmt.Println("输入有误")
	}
}

func chooseMouth() {

	/*
		switch 表达式{
		case 表达值1,2,3：
			逻辑一
		case 表达值4,5,6：
			逻辑二
		default:
			默认逻辑
		}
	*/
	month := 5
	switch month {
	case 1, 2, 3, 4:
		fmt.Println("第一季度")
	case 5, 6, 7, 8:
		fmt.Println("良好")
	case 9, 10, 11, 12:
		fmt.Println("及格")
	default:
		fmt.Println("输入有误")
	}
}

// 高级选择条件写法
func advanceChoose() {
	switch score := 10; score {
	case 1, 2, 3, 4:
		fmt.Println("第一季度")
	case 5, 6, 7, 8:
		fmt.Println("良好")
	case 9, 10, 11, 12:
		fmt.Println("及格")
	default:
		fmt.Println("输入有误")

	}
}

/*
	for range

for key, value := range coll  value始终为集合的值拷贝,只读
可以遍历数组，切片，字符串，map及管道
数组、切⽚、字符串返回索引和值。
map 返回键和值。
channel只返回管道内的值
*/
func forResult(args ...int) bool {
	for _, v := range args {
		if v > 10 {
			return false
		}
	}
	return true
}

// 无表达式的switch
func naCase() {
	score := 98
	switch {
	case score >= 90 && score <= 100:
		fmt.Println("优秀")
	case score >= 80 && score < 90:
		fmt.Println("良好")
	default:
		fmt.Println("还可以")
	}
}

// fallthrough  强制执行后面的case代码
func through() {
	score := 98
	switch {
	case score >= 90 && score <= 100:
		fmt.Println("优秀")
	case score >= 80 && score < 90:
		fmt.Println("良好")
	default:
		fmt.Println("还可以")
	}
}

// 循环语句
func loop() {
	/*
		// 三个表达式
			for i := 10; i <= 100; i++ {
			}

			// 一个表达式
			a := 10
			for a < 20 {
			}

			// forrange
			for range表达式 {

			}

			// 无表达式 break退出循环
			for {
				break
			}
			for ; ; {
				continue
			}
	*/
}

// 结束当次循环
func continueLoop() {
	for n := 1; n <= 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

// 结束循环方法
func breakLoop() {
	for n := 10; n < 12; n++ {
		if n == 11 {
			break
		}
		fmt.Println(n)
	}
}

// 一个条件表达式
func singleLoop() {
	num := 0
	for num < 4 {
		fmt.Println(num)
		num++
	}
}

// 单for循环
func forTest() {
	sum := 0
	for {
		sum++
		if sum > 100 {
			//break是跳出循环
			break
		}
	}
}

func advance() {
	step := 2
	//初值可以省略，但是;必须有，但是这样写step的作⽤域就⽐较⼤了，脱离了for循环
	for ; step > 0; step-- {
		fmt.Println(step)
	}
}

// 简化
func simple() {
	step := 2
	for step > 0 {
		step--
		fmt.Println(step)
	}
}

// 关键字 defer 用于注册延迟调用。
// 这些调用直到 return 前才被执。因此，可以用来做资源清理。 类似于java finally
var s = "学习go"

func showS() string {
	defer func() {
		s = "学习"
	}()
	fmt.Println(s)
	return s
}
func deferTest() { // 3   2   1  先进后出
	name := "1"
	defer fmt.Println(name)

	name = "2"
	defer fmt.Println(name)

	name = "3"
	fmt.Println(name)
}

func visitUrl() {
	url := "http://www.baidu.com"
	r, err := http.Get(url)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}
	if r.StatusCode != http.StatusOK {
		panic(r.StatusCode)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("请求成功  ：%s", body)
	n := rand.Intn(10)
	fmt.Printf("休眠 %d 秒", n)
	time.Sleep(time.Duration(n) * time.Second)
}

func download(url string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
		return
	}
	// 自定义header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", "www.go-edu.cn")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	resp, err := client.Do(req)
	if err != nil {
		panic(err.Error())
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
		return
	}
	fmt.Println(string(body))

	n := rand.Intn(10)
	fmt.Printf("休眠%d秒", n)
	time.Sleep(time.Duration(n) * time.Second)
}

func length(s string) int {
	println("call length.")
	return len(s)
}

/*
return 结束整个方法
break  结束 for、switch 和 select 的代码块  后⾯添加 标签 ，表示退出某个标签对应的代码块
continue 结束当前循环，开始下⼀次的循环迭代过程。仅限在 for 循环内使⽤，在后添加标签时，表示开始标签对应的循环
标签 要求必须定义
在对应的 for 、 switch 和 select 的代码块上
*/
func breakTest() {
OuterLoop:
	for a := 0; a < 20; a++ {
		for i := 0; i < 5; i++ {
			switch i {
			case 2:
				fmt.Println(a, i)
				break OuterLoop // 表示结束整个标签代码
			}
			fmt.Println(a, i)
		}
	}
}

func continueTest() {
OuterLoop:
	for a := 0; a < 20; a++ {
		for i := 0; i < 5; i++ {
			switch i {
			case 2:
				fmt.Println(a, i)
				continue OuterLoop // 表示开始 标签对应的循环
			}
			fmt.Println(a, i)
		}
	}
}

func main() {
	//url := "http://www.go-edu.cn/"
	//for {
	//	download(url)
	//}

	// goto 流程跳转 统一异常处理 多层循环退出
	//goto,goto语句与标签之间不能有变量声明，否则编译报错。
	fmt.Println("学go")
	goto label
	fmt.Println("学java")
label:
	fmt.Println("《学习使我快乐》")

	s := "abcd"
	// 这样写会多次调佣length函数
	for i := 0; i < length(s); i++ { // for i,n:= 0,length(s); i <n; i++
		println(i, s[i])
	}
}
