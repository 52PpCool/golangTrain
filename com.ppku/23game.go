package main

import (
	"bufio"
	"fmt"
	"os"
)

var level = 1
var ex = 0

func main() {
	// 需求：能升级打怪
	fmt.Println("请输入你的角色名字：")
	name := Input()
	fmt.Printf("角色创建成功！你好，%s.你当前的等级为%d \n", name, level)
	fmt.Printf("你遇到一个怪物，请选择战斗还是逃跑？\n\t1.战斗\n\t2.逃跑 \n")
	for {
		choose := Input()
		switch choose {
		case "1":
			ex += 10
			fmt.Println("攻击！你成功了，你杀死了怪物！")
			completeLevel()
			fmt.Printf("你现在的经验为%d,等级为%d", ex, level)
		case "2":
			fmt.Println("逃跑成功！！等级不变")
		case "exit":
			fmt.Println("你退出了游戏")
			os.Exit(1)
		default:
			fmt.Println("请重新输入！")
		}
	}
}

func completeLevel() {
	if ex < 20 {
		level = 1
	} else if ex < 40 {
		level = 2
	} else if ex < 60 {
		level = 3
	} else {
		level = 4
		fmt.Println("你满级了！")
	}
}
func Input() string {
	read := bufio.NewReader(os.Stdin)
	input, err := read.ReadString('\n')
	if err != nil {
		panic(err)
	}
	// 删除最后的\n
	inputs := input[:len(input)-2]
	return inputs
}
