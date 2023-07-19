package main

import "fmt"

func main() {
	var a int = 21
	var b int = 10
	var c int
	c = a + b
	fmt.Println(c) // 31
	c = a - b
	fmt.Println(c) // 11
	c = a / b
	fmt.Println(c) // 2
	c = a % b
	fmt.Println(c) // 1
	a++
	fmt.Println(a) // 22
	a--
	a--
	fmt.Println(a) // 20
}
