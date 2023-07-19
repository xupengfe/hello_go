package main

import "fmt"

// 定义接口 包含公共方法
type user interface {
	talking()
}

// 定义一个struct类
type memo struct {
}

// 实现接口user中方法talking
func (m *memo) talking() {
	fmt.Println("Talking...")
}

func main() {
	mm := memo{}
	mm.talking()
}
