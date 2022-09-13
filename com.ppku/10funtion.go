package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
			函数可见性：
			首字母大写 public
			首字母小写 private

		func name(parameterList) (resultList){
		        // 函数体
		}

		func (parameterList) (resultList){
			// 函数体
		}
	*/
	fmt.Println("56+1=", sum(56, 1))
	printBook()
	fmt.Println("56-1=", sub(56, 1))

	fmt.Println("===============")
	fmt.Println(show("学java", "学go", "学分布式"))

	fmt.Println("===============")
	printType(57, 54.15, "学习")

	fmt.Println("===============")
	book, err := showBookInfo("从头到尾一人成神", "Ppku")
	fmt.Printf("bookInfo=%s,err=%v \n", book, err)

	book1, err1 := showBookInfo("学成失败！", "")
	fmt.Printf("bookInfo=%s,err=%v \n", book1, err1)
}

func showBookInfo(bookName string, autherName string) (string, error) {
	if bookName == "" {
		return "", errors.New("图书名称不能为空！")
	}
	if autherName == "" {
		return "", errors.New("作者名称不能为空！")
	}
	return "图书名称:" + bookName + "作者：" + autherName, nil
}

func showBookInfo2(bookName string, autherName string) (info string, err error) {
	if bookName == "" {
		err = errors.New("图书名称不能为空！")
		return
	}
	if autherName == "" {
		errors.New("作者名称不能为空！")
		return
	}
	info = "图书名称:" + bookName + "作者：" + autherName
	return
}

func printType(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case string:
			fmt.Println("string")
		case int:
			fmt.Println("int")
		case float64:
			fmt.Println("float64")
		default:
			fmt.Println("unknow type")
		}
	}
}

func show(args ...string) int {
	sum := 0
	for _, item := range args {
		fmt.Println(item)
		sum += 1
	}
	return sum
}

func sub(x int, y int) int {
	return x - y
}

func printBook() {
	fmt.Println("《go入门到精通》")
}

func sum(x int, y int) int {
	return x + y
}
