package main

import (
	"fmt"
	"os"
)

func main() {
	/*
			type error interface{
			Error() string
			}
		错误表示可能出现问题的地方出现了问题 Error

	*/
	FindFile()

	a := 100
	b := -100
	area, err := area(a, b)
	if err != nil {
		// 断言 *areError 类型
		if err, ok := err.(*areaError); ok {
			// 如果错误类型是areError
			// 如果长度为负数 打印错误长度具体值
			if err.lengthNegative() {
				fmt.Printf("error: 长度 %d 小于0 \n", err.length)
			}
			// 如果宽度为负数，打印具体值
			if err.widthNegative() {
				fmt.Printf("error: 宽度 %d 小于0 \n", err.width)
			}
			return
		}
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("area: ", area)
}

type areaError struct {
	err    string //错误信息
	length int    // 长度
	width  int    // 宽度
}

func (e *areaError) Error() string {
	return e.err
}

// 长度为负返回true
func (e *areaError) lengthNegative() bool {
	return e.length < 0
}

// 宽度为负返回true
func (e *areaError) widthNegative() bool {
	return e.width < 0
}

func area(length, width int) (int, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err += "width is less than zero"
		} else {
			err += "and width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaError{err, length, width}
	}
	return length * width, nil
}

func FindFile() {
	file, err := os.Open("/a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file.Name(), "opened successfully")
}
