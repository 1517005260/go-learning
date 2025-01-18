// defer: 函数的亡语  LIFO
package main

import "fmt"

func deferFunc() {
	fmt.Println("defer func")
}

func returnFunc() int {
	fmt.Println("return func")
	return 0
}

func deferAndReturn() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	defer fmt.Println("This is the end 1")
	defer fmt.Println("This is the end 2")

	fmt.Println("hello 1")
	fmt.Println("hello 2")

	defer fmt.Println("This is the end 3")

	// defer 后于 return 执行
	deferAndReturn()
}

/*
	输出：
	hello 1
	hello 2
	return func
	defer func
	This is the end 3
	This is the end 2
	This is the end 1
*/
