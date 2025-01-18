// slice: 动态数组
package main

import "fmt"

// 动态数组传参
func printArray(arr []int) {
	for _, v := range arr { // 匿名 i，仅关心v
		fmt.Println("val: ", v)
	}

	// 持有原数组指针
	arr[0] = 100
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(arr)        // [1 2 3]
	fmt.Printf("%T\n", arr) // []int
	printArray(arr)

	fmt.Println("====After function===")
	fmt.Println(arr) // [100 2 3]
}
