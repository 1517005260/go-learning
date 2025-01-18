package lib1 // 一般包名和文件名一致

import "fmt"

func Lib1_test() { // 函数名必须要大写，才能被外部调用。小写的是私有方法
	fmt.Println("lib1 test")
}

// 建议将 init() 放在文件的最后，确保所有依赖都已经加载
func init() {
	fmt.Println("lib1 init")
}
