package main

import "fmt"

// readme 修改为指针版本
func changeValue(p *int) {
	*p = 200
}

// 经典指针使用：交换两数
func mySwap(b *int, c *int) {
	var tmp int
	tmp = *b
	*b = *c
	*c = tmp
}

func main() {
	var a int = 100
	fmt.Println("a.address :", &a)

	changeValue(&a)
	fmt.Println("a :", a)

	var b, c int = 100, 200
	mySwap(&b, &c)
	fmt.Println("b :", b, "c :", c)
}
