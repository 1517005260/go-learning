package main

import "fmt"

func printMap(myMap map[string]string) {
	// 引用传递
	for k, v := range myMap {
		fmt.Println("key:", k, "val:", v)
	}
}

func main() {
	cityMap := make(map[string]string)

	// add
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "New York"

	// 遍历
	for key, value := range cityMap {
		fmt.Println("key:", key, "value:", value)
	}

	// delete
	delete(cityMap, "China")
	fmt.Println(cityMap)

	// modify
	cityMap["USA"] = "abc"
	fmt.Println(cityMap)

	printMap(cityMap)
}
