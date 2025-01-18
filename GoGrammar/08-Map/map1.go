package main

import "fmt"

func main() {
	// map 的声明
	var map1 map[string]string // [key]value
	if map1 == nil {
		fmt.Println("map1 is nil")
	}

	// make 分配空间
	map1 = make(map[string]string, 10)

	map1["one"] = "java"
	map1["two"] = "python"
	map1["three"] = "golang"
	fmt.Println(map1) // map[one:java three:golang two:python]

	map2 := make(map[string]int)
	map2["one"] = 1
	map2["two"] = 2
	fmt.Println(map2) // map[one:1 two:2]

	map3 := map[string]int{
		"abc": 1,
		"def": 2,
	}
	fmt.Println(map3) // map[abc:1 def:2]
}
