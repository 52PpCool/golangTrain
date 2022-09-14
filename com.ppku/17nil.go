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
	*/
	FindFile()

	a := 100
	b := -100
	i, err := area(a, b)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(i)
}

type areaError struct {
	err    string //错误信息
	length int    // 长度
	width  int    // 宽度
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
