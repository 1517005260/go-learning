package main

import "fmt"

func main() {
	// slice 的多种声明方式：

	slice1 := []int{1, 2, 3}
	fmt.Println(slice1) // [1 2 3]

	var slice2 []int
	fmt.Println(slice2) // []
	// slice2[0] = 1 会报错 index out of range [0] with length 0，即越界错

	// 空切片的判断
	if slice2 == nil {
		fmt.Println("slice2 is nil")
	}

	// 为 [] 开辟空间
	slice2 = make([]int, 3)
	fmt.Println(slice2) // [0 0 0]
	slice2[0] = 100
	fmt.Println(slice2) // [100 0 0]

	var slice3 []int = make([]int, 10)
	fmt.Println(slice3) // [0 0 0 0 0 0 0 0 0 0]

	slice4 := make([]int, 3)
	fmt.Println(slice4) // [0 0 0]
}
