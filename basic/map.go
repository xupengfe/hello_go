package main

import "fmt"

func main() {
	// 定义 变量strMap
	var strMap map[int]string
	// 进行初始化
	strMap = make(map[int]string)

	// 给map 赋值
	for i := 0; i < 5; i++ {
		//fmt.Println("i:", i)
		strMap[i] = "Test_MAP" + fmt.Sprint(i)
	}

	// 打印出map值
	for k, v := range strMap {
		fmt.Println(k, ":", v)
	}

	// 打印出map 长度
	fmt.Println("Lenth of the strMap is:", len(strMap))
}
