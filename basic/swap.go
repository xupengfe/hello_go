package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello hi", "world wow")
	fmt.Println("a:", a, "| b:", b)
}
