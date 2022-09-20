package main

import "fmt"

func main() {
	// 异常表示不应该出现问题的地方出现了问题 panic
	//惯例是:导致关键流程出现不可修复性错误的使⽤ panic，其他使⽤ error
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


						异常的使⽤场景简单描述：Go中可以抛出⼀个panic的异常，然后在defer中通过
						recover捕获这个异常，然后正常处理。

					panic：
					1. 内置函数
					2. 假如函数F中书写了panic语句，会终⽌其后要执⾏的代码，在panic所在函数F
					内如果存在要执⾏的defer函数列表，按照defer的逆序执⾏
					3. 返回函数F的调⽤者G，在G中，调⽤函数F语句之后的代码不会执⾏，假如函数
					G中存在要执⾏的defer函数列表，按照defer的逆序执⾏
					4. 直到goroutine整个退出，并报告错误

				recover：
				1. 内置函数
				2. ⽤来捕获panic，从⽽影响应⽤的⾏为

			golang 的错误处理流程：当⼀个函数在执⾏过程中出现了异常或遇到
			panic()，正常语句就会⽴即终⽌，然后执⾏ defer 语句，再报告异常信息，最
			后退出 goroutine。如果在 defer 中使⽤了 recover() 函数,则会捕获错误信息，
			使该错误信息终⽌报告。


		注意:
		1. 利⽤recover处理panic指令，defer 必须放在 panic 之前定义，另外 recover 只
		有在 defer 调⽤的函数中才有效。否则当panic时，recover⽆法捕获到panic，
		⽆法防⽌panic扩散。
		2. recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后
		的那个点。
		3. 多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调⽤。

	*/
	//testDefer()

	// 越界打印错误
	outOfArray(10)
	// 后续可以继续运行
	fmt.Println("main")

	//延迟调⽤中引发的错误，可被后续延迟调⽤捕获，但仅最后⼀个错误可被捕获
	defer func() {
		// defer panic 会打印
		fmt.Println(recover())
	}()
	defer func() {
		panic("defer panic")
	}()
	panic("test panic")

	//需要保护代码段，可将代码块重构成匿名函数，如此可确保后续代码被执
	test(2, 1)
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
func test(x, y int) {
	var z int
	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		panic("test panic")
		z = x / y
		return
	}()
	fmt.Printf("x / y = %d\n", z)
}
