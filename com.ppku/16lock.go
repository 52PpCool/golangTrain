package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	// sync.Cond
	s1 := []string{"张三"}
	s2 := []string{"李四"}
	s3 := []string{"王五"}

	var m sync.Mutex
	cond := sync.NewCond(&m) //带锁的条件变量

	// Broadcast 唤醒所有因等待条件阻塞的goroutine
	// Signal 唤醒一个因等待条件阻塞的goroutine
	// Wait 等待解锁并挂起goroutine，在稍后恢复执行后，Wait返回锁定c.L,只有当被Broadcast或Signal唤醒，wait才返回。

	go listen("学习java", s1, cond)
	go listen("学习go", s2, cond)
	go listen("学习基础", s3, cond)
	
	go broadcast("秒杀开始",cond)
	
	ch:=make(chan os.Signal,1)
	signal.Notify(ch,os.Interrupt)
	<-ch
}
func listen(name string, s []string, c *sync.Cond) {
	c.L.Lock()
	c.Wait()
	fmt.Println(name, "报名：", s)
	c.L.Unlock()
}

func broadcast(event string, c *sync.Cond) {
	time.Sleep(time.Second)
	c.L.Lock()
	fmt.Println(event)
	c.Broadcast()
	c.L.Unlock()
}

type BankV3 struct {
	balance int          //余额
	rwMutex sync.RWMutex //读写锁，适用读多写少
	// 读锁 RLock() RUnlock()
	// 写锁 Lock()  Unlock()
}

func (b *BankV3) Deposit(amount int) { // 存款
	b.rwMutex.Lock()         //加锁
	defer b.rwMutex.Unlock() // 解锁
	b.balance += amount
}

func (b *BankV3) Balance() (balance int) {
	b.rwMutex.RLock()
	defer b.rwMutex.RUnlock()
	return b.balance
}

type BankV2 struct {
	balance int        //余额
	m       sync.Mutex // 互斥锁
}

func (b *BankV2) Deposit(amount int) { // 存款
	b.m.Lock()         //加锁
	defer b.m.Unlock() // 解锁
	b.balance += amount
}

func (b *BankV2) Balance() (balance int) {
	return b.balance
}

type Bank struct {
	balance int // 余额
}

func (b *Bank) Deposit(amount int) { // 存款
	b.balance += amount
}

func (b *Bank) Balance() (balance int) {
	return b.balance
}
