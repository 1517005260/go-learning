// slice 相关 api —— 截取
package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	s1 := s[0:2]    // [0,2)
	fmt.Println(s1) // [1 2]

	// 其他截取方式和python类似
	s[0] = 100
	fmt.Println(s1) // [100 2]
	// 所以切片和原动态数组都是指向同一片空间的

	// copy： 创建副本
	s2 := make([]int, 3)
	copy(s2, s)
	fmt.Println(s2) // [100 2 3]
	s[0] = 50
	fmt.Println(s2) // [100 2 3]
}
