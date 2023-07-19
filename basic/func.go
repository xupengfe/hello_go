package main

import "fmt"

// 定义 Write函数 返回值有两个，一个为name，一个age为
func Write() (name string, age int) {
	return "Dong", 9
}

// 定义 Read函数
func Read(name string, age int) {
	fmt.Println(name, "已经", age, "岁了")
}

func main() {
	Read(Write())
}
