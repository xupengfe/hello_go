package main

import "fmt"

func main() {
	// 方式一
	a := make([]int, 5) // 初始化长度为5的slice,默认值为零
	for i := 0; i < 5; i++ {
		a = append(a, i)
	}
	a = append(a, 6)
	fmt.Println(a) // [0 0 0 0 0 0 1 2 3 4 6]

	// 方式二
	var b []int
	for i := 0; i < 7; i++ {
		b = append(b, i)
	}
	fmt.Println(b) // [0 1 2 3 4 5 6]
}
