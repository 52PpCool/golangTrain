package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
			通道的声明
			var chanName chan channelType
		 	var ch chan string
			ch = make(chan string)
			ch:=make(chan string)

			发送数据
			channelName <-data
			接受数据
			value:= <-channelName


		 关闭通道
		 close(channelName)
	*/

	//创建一个通道
	ch := make(chan string)
	go Println(ch)
	// 从通道中接收数据
	rec := <-ch
	fmt.Println(rec)
	time.Sleep(time.Second)
	close(ch)

	c := make(chan int, 3)
	fmt.Println("初始化前")
	fmt.Println("cap:", cap(c))
	fmt.Println("len:", len(c))
	c <- 1
	c <- 2
	c <- 3
	fmt.Println("传入三个参数后")
	fmt.Println("cap:", cap(c))
	fmt.Println("len:", len(c))
	receive := <-c
	fmt.Println(receive)
	fmt.Println("读取一个参数后")
	fmt.Println("cap:", cap(c))
	fmt.Println("len:", len(c))

	/*
		无缓冲通道
		c:=make(chan int,0)
		c:=make(chan int)

		缓冲通道
		c:=make(chan int,3)
	*/

	i := make(chan int)
	go func() {
		fmt.Println("send 1")
		i <- 1
	}()

	go func() {
		n := <-i
		fmt.Println("receive", n)
	}()
	time.Sleep(1 * time.Second)

	/*
		只读通道 <-chan
		只写通道 chan <-
	*/

	// 创建一个双向通道  遍历通道
	var ch2 = make(chan int, 5)
	go loopPrint(ch2)
	for v := range ch2 {
		fmt.Println(v)

	}

	ch3 := make(chan bool, 1)
	var x int
	for i := 0; i < 10000; i++ {
		go increment(ch3, &x)
	}
	time.Sleep(time.Second)
	fmt.Println("X=", x)

	ch4 := make(chan bool)
	ch4 <- true //fatal error: all goroutines are asleep - deadlock!

	ch5 := make(chan bool)
	go funcReceive(ch5)
	ch5 <- true
	fmt.Println(<-ch5) // fatal error: all goroutines are asleep - deadlock!
	time.Sleep(time.Second)

	ch6 := make(chan bool, 1)
	ch6 <- true
	ch6 <- false
	fmt.Println(<-ch6) //fatal error: all goroutines are asleep - deadlock!

	/*
		waitgroup  等待一组任务完成，再执行其他任务

		add() 初始值0，累加子协程的数量
		Done() 子协程完成后计数器减一，defer调用
		wait() 阻塞当前协程，知道计数器归零
	*/

	isDone := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		isDone <- true
	}()

	<-isDone
	fmt.Println("main grouptime finished")

	// 初始化sync.WaitGroup
	var wg sync.WaitGroup
	// 子协程数量
	wg.Add(3)
	// 开启子协程
	go DoTask(1, &wg)
	go DoTask(2, &wg)
	go DoTask(3, &wg)
	// 阻塞完成
	wg.Wait()
}

func DoTask(task int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Printf("task:%d:%d \n", task, i)
	}
}

func funcReceive(ch5 chan bool) {
	fmt.Println(<-ch5)
}

func increment(ch3 chan bool, x *int) {
	ch3 <- true
	*x = *x + 1 //不是原子操作，避免多协程进行操作，使用容量为1的通道，达到锁的效果
	<-ch3
}

func loopPrint(ch2 chan int) {
	for i := 0; i < 10; i++ {
		ch2 <- i
	}
	close(ch2)
}

func Println(ch chan string) {
	ch <- "学习go"
}
