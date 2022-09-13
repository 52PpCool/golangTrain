package main

import "fmt"

func main() {
	/*  方法
	func (接受者 结构体类型)  方法名(参数列表) 返回值列表{
		逻辑体
	}
	*/
	lesson := Lesson1{
		"a",
		"b",
		5,
	}
	lesson.ShowInfo()
	lesson.AddTime(2)
	fmt.Println(lesson)
}

type Lesson1 struct {
	name   string //课程名称
	target string //学习目标
	spend  int    //学习时间
}

func (l Lesson1) ShowInfo() {
	fmt.Println(l)
}

func (l *Lesson1) AddTime(n int) Lesson1 {
	l.spend += n
	return *l
}
