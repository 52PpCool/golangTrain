package main

import "fmt"

func main() {
	// 异常表示不应该出现问题的地方出现了问题 panic
	/*
			两个合理的用例：
			1. 发生了一个不能恢复的错误，程序不能继续运行
			2. 发生了一个编程上的错误。假如其他人使用nil作为入参调用一个接受指针的方法

			func panic(v interface{})
			当发生panic时在执行完所有延迟函数后才抛出异常

			recover() 内建函数 ，用于重新获得panic协程的控制
			func recover() interface{}

			recover 必须在 defer 函数中才能生效，在其他作用域下，它是不工作的。
			在延迟函数内调用 recover ，可以取到 panic 的错误信息，
			并且停止 panic 续发事件，程序运行恢复正常。
		只有在相同的协程中调用 recover 才管用， recover 不能恢复一个不同协程的 panic 。
	*/
	//testDefer()

	// 越界打印错误
	outOfArray(10)
	// 后续可以继续运行
	fmt.Println("main")

}
func testDefer() {
	defer fmt.Println("defer ...")
	panic("panic")
}

func outOfArray(s int) {
	defer func() {
		// recover()  可以将捕获到的panic信息打印
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var array [5]int
	array[s] = 1
}
