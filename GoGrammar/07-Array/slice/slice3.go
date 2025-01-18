// slice 相关 api —— 追加
package main

import "fmt"

// append 会根据 cap 自动追加
func main() {
	nums := make([]int, 3, 5)                                                 // 初始大小，容量
	fmt.Printf("len = %d, cap = %d, nums = %v\n", len(nums), cap(nums), nums) // len = 3, cap = 5, nums = [0 0 0]

	nums = append(nums, 1)
	fmt.Printf("len = %d, cap = %d, nums = %v\n", len(nums), cap(nums), nums) // len = 4, cap = 5, nums = [0 0 0 1]

	nums = append(nums, 2)
	nums = append(nums, 3)                                                    // 此时已经没空间了，go语言底层会再开辟一块空间
	fmt.Printf("len = %d, cap = %d, nums = %v\n", len(nums), cap(nums), nums) // len = 6, cap = 10, nums = [0 0 0 1 2 3]

	fmt.Println("=============")

	nums1 := make([]int, 3)                                                      // len = cap
	fmt.Printf("len = %d, cap = %d, nums = %v\n", len(nums1), cap(nums1), nums1) // len = 3, cap = 3, nums = [0 0 0]

	nums1 = append(nums1, 1)
	fmt.Printf("len = %d, cap = %d, nums = %v\n", len(nums1), cap(nums1), nums1) // len = 4, cap = 6, nums = [0 0 0 1]
}
