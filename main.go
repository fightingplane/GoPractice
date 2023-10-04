package main

import (
	"fmt"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//输入a和b并得出他们的积
func multiple(a, b int64) (result int64) {
	result = a * b
	return
}

func sum(start, end int) (result int) {
	i := start
	result = 0
	for i <= end {
		result += i
		i++
	}
	return result
}

func main() { // main函数，是程序执行的入口

	// fmt.Println("Hello World! Hello GoLang") // 在终端打印 Hello World!

	// fmt.Println("The time is ", time.Now())

	// tmp := sum(1, 3)
	// fmt.Println(tmp)
	// fmt.Println(split(17))
	var start, end int64
	fmt.Println("请输入数字a")
	_, err := fmt.Scanf("%d", &start)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("请输入数字b")
	_, err2 := fmt.Scanf("%d", &end)
	if err2 != nil {
		fmt.Println(err2)
	}

	result := multiple(start, end)
	fmt.Printf("数字%v x %v的积为%v\n", start, end, result)
}
