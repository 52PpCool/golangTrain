package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	/*
		三种获取field
		filed：=reflect.TypeOf(obj).FieldByName("Name")
		filed：=reflect.ValueOf(obj).Type().Field(i)
		filed：=reflect.ValueOf(&obj).Elem().Type().Field(i)

		获取Tag
		tag:=filed.Tag

		获取键值
		keyValue:=tag.Get("key") 没有获取到对应的内容后，会返回空字符串
		keyValue:=tag.Lookup("key")
	*/
	book1 := books{
		Name:   "学习java",
		Target: "学习go",
		Spend:  5,
	}
	data, err := json.Marshal(book1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", data)
	fmt.Println("------------------------")

	book2 := books{
		Name:   "从0到Go语言微服务架构师",
		Target: "全面掌握Go语言微服务落地，代码级一次性解决微服务和分布式。",
		Spend:  9,
	}
	data2, err := json.Marshal(book2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data2)
	fmt.Println("------------------------")

	b := reflect.TypeOf(book{})
	name, _ := b.FieldByName("Name")
	tag := name.Tag
	fmt.Println("Name Tag:", tag)

	keyValue, _ := tag.Lookup("json")
	fmt.Println("key:json,value:", keyValue)

}

type books struct {
	Name   string `json:"name"`
	Target string `json:"target"`
	Spend  int    `json:"speed,omitempty"`
}
