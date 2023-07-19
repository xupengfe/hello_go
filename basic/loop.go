package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		if i == 4 {
			continue
		} else if i == 5 {
			break
		}
		fmt.Println(i)
	}
}
