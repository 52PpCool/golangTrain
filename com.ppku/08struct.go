package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 创建结构体 字段名
	lesson := Lesson{
		name:   "go",
		target: "基础练习",
		spend:  "早上",
	}
	fmt.Println(lesson)

	lesson1 := Lesson3{
		name:   "s",
		target: "b",
		spend:  5,
		author: Author{
			name: "Gist",
			wx:   "work",
		},
	}
	fmt.Println(lesson1)

	// 字段匿名
	lesson2 := Lesson3{
		"I",
		"L",
		5,
		Author{
			"gist",
			"work",
		},
	}
	fmt.Println(lesson2)

	// 匿名结构体
	lesson3 := struct {
		name, target string
		spend        int
	}{"学go",
		"架构",
		25,
	}
	fmt.Println(lesson3)
	fmt.Println("===================")

	// 结构体的零值
	var lesson4 = Lesson{}
	fmt.Println(lesson4)

	lesson4.spend = "a"
	lesson4.target = "b"
	lesson4.name = "c"
	fmt.Println(lesson4)
	fmt.Println("===================")

	var lesson5 = Lesson{
		"学go",
		"学java",
		"15",
	}
	fmt.Println(lesson5)
	fmt.Println(lesson5.spend)
	fmt.Println(lesson5.name)
	fmt.Println(lesson5.target)
	fmt.Println("===================")

	// 指针创建
	lesson6 := &Lesson{
		"go",
		"java",
		"15",
	}
	// c语言 lesson6->name  go语言 .
	fmt.Println(lesson6.name)
	fmt.Println((*lesson6).name)
	fmt.Println("====================")

	lesson7 := Lesson02{
		"a",
		6,
	}
	fmt.Println(lesson7)
	fmt.Println(lesson7.string)
	fmt.Println(lesson7.int)
	fmt.Println("====================")

	lesson8 := Lesson{
		name:   "从0到架构师",
		target: "全面掌握",
		spend:  "49",
	}
	lesson9 := Lesson{
		name:   "从0到架构师",
		target: "全面掌握",
		spend:  "49",
	}
	fmt.Println(lesson8 == lesson9)
	fmt.Println(reflect.DeepEqual(lesson9, lesson8))
	fmt.Println("---------------------")
}

type Author struct { // type name struct

	name string
	wx   string
}
type Lesson struct {
	name   string
	target string
	spend  string
}

type Lesson01 struct {
	name, target, spend string
}
type Lesson02 struct {
	string
	int
}

type Lesson3 struct {
	name, target string
	spend        int
	author       Author
}
