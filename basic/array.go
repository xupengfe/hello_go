package main

import "fmt"

func main() {
	var arr1 [4]int   // 元素自动初始化为零
	fmt.Println(arr1) // [0 0 0 0]

	arr2 := [4]int{1, 2} // 其他未初始化的元素为零
	fmt.Println(arr2)    // [1 2 0 0]

	arr3 := [4]int{5, 3: 10} // 可指定索引位置初始化
	fmt.Println(arr3)        // [5 0 0 10]

	arr4 := [...]int{1, 2, 3, 3, 6} // 编译器按照初始化值数量确定数组长度
	fmt.Println(arr4)               // [1 2 3 3 6]

	t := len(arr4) // 内置函数len(数组名称)表示数组的长度
	fmt.Println(t) // 5
}
