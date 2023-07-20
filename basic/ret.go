package main

import "fmt"

func split(sum int) (x int) {
	x = sum * 4 / 9
	y := sum - x
	fmt.Println("x:", fmt.Sprint(x), "y:", fmt.Sprint(y))
	// It's better to write as "return x"
	return
}

func main() {
	fmt.Println(split(17))
}
