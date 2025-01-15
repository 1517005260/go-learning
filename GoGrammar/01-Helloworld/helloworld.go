package main // 声明可以被 go run helloworld.go 执行

import (
	"fmt"
	"time"
)

func main() { // golang 的大括号不能换行
	fmt.Println("Hello World")

	time.Sleep(1 * time.Second)
}

// 除了 go run之外，还可以手动先 go build 生成可执行文件，然后再执行
