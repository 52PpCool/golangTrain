package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		select{
		case expression1:
			code
		case expression2:
			code
		default:
			code
		}
	*/
	SelectDemo01()

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go task1(ch1)
	go task2(ch2)
	go task3(ch3)
	
	select {
	case message1 := <-ch1:
		fmt.Println("ch1", message1)
	case message2 := <-ch2:
		fmt.Println("ch2", message2)
	case message3 := <-ch3:
		fmt.Println("ch3", message3)
	default:
		fmt.Println("No Data Received")
	}

	// 死锁情况
	Deadlock()
	EmptySelect()
	
	// Timeout
	Timeout()
	
	selectDemo02()
}

func selectDemo02() {
	c1:=make(chan string,2)
	c1<-"学习go"
	select {
	case c1 <- "学习java":
		fmt.Println("c1 Receive",<-c1)
		fmt.Println("c1 Receive",<-c1)
	default:
		fmt.Println("channel blocking")
	}
}

func Timeout() {
	ch1:=make(chan string,1)
	ch2:=make(chan string,1)
	ch3:=make(chan string,1)
	timeout:=make(chan bool,1)
	
	// 超时
	go makeTimeout(timeout,2)


	select {
	case message1 := <-ch1:
		fmt.Println("ch1", message1)
	case message2 := <-ch2:
		fmt.Println("ch2", message2)
	case message3 := <-ch3:
		fmt.Println("ch3", message3)
	case <-timeout:
		fmt.Println("Timeout")
	}
}

func makeTimeout(timeout chan bool, i int) {
	time.Sleep(time.Second*time.Duration(i))
	timeout<-true
}

func EmptySelect() {
	select {}
}

func Deadlock() {
	// 发生死锁   all goroutines are asleep - deadlock!
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 1)
	select {
	case message1 := <-ch1:
		fmt.Println("ch1 received:", message1)
	case message2 := <-ch2:
		fmt.Println("ch2 received:", message2)
	case message3 := <-ch3:
		fmt.Println("ch3 received:", message3)
	}
}

func task1(ch chan string) {
	time.Sleep(5 * time.Second)
	ch <- "学习基础"
}

func task2(ch chan string) {
	time.Sleep(7 * time.Second)
	ch <- "学习go"
}

func task3(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "学习算法"
}

func SelectDemo01() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 1)
	ch1 <- "学习基础"
	ch2 <- "学习go"
	ch3 <- "学习算法"

	select {
	case message1 := <-ch1:
		fmt.Println("ch1:", message1)
	case message2 := <-ch2:
		fmt.Println("ch1:", message2)
	case message3 := <-ch3:
		fmt.Println("ch1:", message3)
	default:
		fmt.Println("No Data Received")
	}
}
