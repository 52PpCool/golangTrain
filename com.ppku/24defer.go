package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	/*
			defer特性:
			1. 关键字 defer ⽤于注册延迟调⽤。
			2. 这些调⽤直到 return 前才被执。因此，可以⽤来做资源清理。
			3. 多个defer语句，按先进后出的⽅式执⾏。
			4. defer语句中的变量，在defer声明时就决定了。

		defer的⽤途：
		1. 关闭⽂件句柄
		2. 锁资源释放
		3. 数据库连接释放
	*/

	//Go 语⾔中所有的 函数调⽤都是传值的
	//调⽤ defer 关键字会 ⽴刻拷⻉函数中引⽤的外部参数 ，包括start 和time.Since中的
	//Now
	//defer的函数在 压栈的时候也会保存参数的值，并⾮在执⾏时取值
	start := time.Now()
	log.Printf("开始时间为：%v", start)
	defer log.Printf("时间差：%v", time.Since(start)) // Now()此时已经 copy进去了
	//不受这3秒睡眠的影响
	time.Sleep(3 * time.Second)
	log.Printf("函数结束")

	//	解决上述问题：使⽤defer fun()
	starts := time.Now()
	log.Printf("开始时间为：%v", starts)
	defer func() {
		log.Printf("开始调⽤defer")
		log.Printf("时间差：%v", time.Since(starts))
		log.Printf("结束调⽤defer")
	}()
	time.Sleep(3 * time.Second)
	log.Printf("函数结束")

	var whatever = [5]int{1, 2, 3, 4, 5}
	for i, _ := range whatever {
		//函数正常执⾏,由于闭包⽤到的变量 i 在执⾏的时候已经变成4,所以输出全都是4.
		defer func() { fmt.Println(i) }()
	}

	// 解决
	var what = [5]int{1, 2, 3, 4, 5}
	for i, _ := range what {
		i := i
		defer func() { fmt.Println(i) }()
	}
}
