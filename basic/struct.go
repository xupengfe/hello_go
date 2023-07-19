package main

import "fmt"

type user struct {
	name string
	age  int
}

// 结构体user Read方法
func (u *user) Read() string {
	return fmt.Sprintf("%s 已经 %d 岁了", u.name, u.age)
}

func main() {
	// 初始化
	u := &user{
		name: "Dong",
		age:  9,
	}
	fmt.Println(u.name, "已经", u.age, "岁了")
	// 调用结构体user的 Read方法
	fmt.Println(u.Read())
}
