package main

import "fmt"

func main() {
	var x int = 20
	//x := 20

	if x > 5 {
		fmt.Println("a")
	} else if x < 5 && x > 0 {
		fmt.Println("b")
	} else {
		fmt.Println("c")
	}
}
