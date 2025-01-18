package main

// 导入的包必须被使用（即一定要调用test方法，否则报错）
import (
	// 像rust一样加 _ 可以向编译器声明仅导包不调用
	_ "go-learning/GoGrammar/04-Import/lib1"   // 调用init()方法
	xyz "go-learning/GoGrammar/04-Import/lib2" // 可以在前面加别名
)

func main() {
	// lib1.Lib1_test() // 调用test方法
	// lib2.Lib2_test()

	xyz.Lib2_test()
}

/*
加 _ 后输出：
lib1 init
lib2 init
lib2 test
*/
