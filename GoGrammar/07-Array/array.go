// array 静态数组
package main

import "fmt"

// 静态数组的传参
func test(arr [10]int) { // 类型写死
	// 这里持有arr的副本，修改不影响arr
	fmt.Println(arr)
}

func main() {
	// 定长数组
	var arr1 [10]int

	// 简单遍历
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	arr2 := [10]int{1, 2, 3, 4} // 剩下的没赋值都是0
	// range 遍历，类似 rust 的 enumerate()，返回 index 和 value
	for i, v := range arr2 {
		fmt.Printf("arr2[%d] = %d\n", i, v)
	}

	// 数组类型查看
	fmt.Printf("type of arr2 is %T\n", arr2) // [10]int

	test(arr2) // [1 2 3 4 0 0 0 0 0 0]
}
