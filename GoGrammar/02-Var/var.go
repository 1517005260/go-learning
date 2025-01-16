package main

import "fmt"

// 声明全局变量仅能var
var aa = 100

func main() {
	var a int             // 默认为0, 且后续必须使用
	fmt.Println("a =", a) // 输出 a = 0，即会自动用 “空格a” 补充
	fmt.Printf("%T\n", a) // 输出类型

	var b int = 100
	fmt.Println("b =", b)
	fmt.Printf("%T\n", b)

	var c = 11.11
	fmt.Println("c =", c)
	fmt.Printf("%T\n", c)

	/*输出：
	a = 0
	int
	b = 100
	int
	c = 11.11
	float64
	*/

	// 更常用的：
	d := 100 // 只能声明局部变量
	fmt.Println("d =", d)
	fmt.Printf("%T\n", d)

	fmt.Println("aa =", aa)

	// 多个变量声明
	var x, y = 'a', 3.14
	fmt.Println("x =", x, ", y =", y)
	// 或者：
	var (
		xx int     = 100
		yy float32 = 3.14
	)
	fmt.Println("xx =", xx, ", yy =", yy)

	testConst()
}
