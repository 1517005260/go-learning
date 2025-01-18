package main

import "fmt"

// 简单的函数定义，以及返回值定义
func f1(a string, b int) int {
	fmt.Println("a:", a, "b:", b)
	return 10
}

// 函数可以有多个返回值
// 返回多个匿名返回值
func f2(a string, b int) (int, int) {
	fmt.Println("a:", a, "b:", b)
	return 555, 666
}

// 返回多个有名返回值
func f3(a string, b int) (r1 int, r2 int) {
	fmt.Println("a:", a, "b:", b)
	fmt.Println("r1:", r1) // r1=0， 被定义后，不赋值，默认均为0
	r1 = 1000              // 已被定义过的变量不需要冒号
	r2 = 2000
	return
}

// 返回相同类型
func f4(a string, b int) (r1, r2 int) {
	fmt.Println("a:", a, "b:", b)
	r1 = 1000
	r2 = 2000
	return
}

func main() {
	c := f1("xxx", 555)
	fmt.Println("c:", c)

	res1, res2 := f2("hello", 999)
	fmt.Println("res1:", res1, "res2:", res2)

	res3, res4 := f3("hello", 100)
	fmt.Println("res3:", res3, "res4:", res4)

	res3, res4 = f4("hello", 100) // 已被定义过的变量不需要冒号
	fmt.Println("res3:", res3, "res4:", res4)
}
