package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 运行时反射，识别 interface{} 变量底层的具体变量和具体值
	//reflect.Type 表示interface{}的具体类型
	//reflect.TypeOf()返回reflect.Type
	//reflect.Kind() 反射类型
	var aa int64 = 123
	reflectType(aa)
	var b string = "学习go"
	reflectType(b)

	book := book{
		"学习go",
		"学习java",
		8,
	}
	reflectNumField(book)

	// 反射第一定律： 反射可以将  接口类型变量  转化为  反射类型变量
	var a interface{} = 3.14
	fmt.Printf("类型为%T,值为%v", a, a)
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	fmt.Printf("接口变量到反射对象：%T \n", t)
	fmt.Printf("接口变量到反射对象：%T \n", v)

	// 反射第二定律： 将 反射类型对象  转换到 接口对象
	i := v.Interface()
	fmt.Printf("从反射对象到接口变量：对象类型%T,值为%v \n", i, i)
	// 使用类型断言进行转换
	x := v.Interface().(float64)
	fmt.Printf("x的类型为%T,值为%v \n", x, x)

	// 反射第三定律： 如果要修改 反射类型对象  其值必须是 可写的
	var c = 3.14
	y := reflect.ValueOf(&c).Elem()
	fmt.Println("是否可写：", y.CanSet())
	y.SetFloat(2.1)
	fmt.Println(y)

}
func reflectNumField(x interface{}) {
	if reflect.ValueOf(x).Kind() == reflect.Struct {
		v := reflect.ValueOf(x)
		fmt.Println("Number of field", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v \n", i, v.Field(i), v.Field(i))
		}
	}
}

type book struct {
	name   string
	target string
	spend  int
}

func reflectType1(x interface{}) {
	typeX := reflect.TypeOf(x)
	fmt.Println(typeX)
}
func reflectType(x interface{}) {
	typeX := reflect.TypeOf(x)
	valueX := reflect.ValueOf(x)
	fmt.Println(typeX)
	fmt.Println(valueX)
}
