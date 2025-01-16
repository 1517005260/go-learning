package main

import "fmt"

// const 定义枚举类型
/* const (
	BEIJING   = 0
	SHANGHAI  = 1
	GUANGZHOU = 2
	SHENZHEN  = 3
)
*/
// 如上写法可以改为：
const (
	BEIJING  = iota // 第一行iota=0,后面的均++
	SHANGHAI        // SHANGHAI = iota , iota =1
	GUANGZHOU
	SHENZHEN
)

const (
	a = iota * 10
	b
)

const (
	c, d = iota + 1, iota + 2 // c=1, d=2, iota=0
	e, f                      // e=2, f=3, iota=1
	g, h = iota * 2, iota * 3 // iota=2, g=4, h=6
)

func testConst() { // 同一文件下只能有一个main函数
	// 常量：只读
	const length int = 10
	fmt.Println("len =", length)

	fmt.Println("b =", b) // b = 10
}
