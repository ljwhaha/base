package main

import (
	"fmt"
)

func main() {
	//创建MAP,不能使用切片当做key , 可以使用切片作为值
	mapCreate()
	//使用map
	mapUse()
	//map循环
	mapLoop()
	//删除map
	removeMap()
}

func mapCreate() {
	map1 := make(map[string]int)
	fmt.Println("make函数创建map：", map1)

	map2 := map[string]int{"1": 1}
	fmt.Println("映射字面量创建map：", map2)
}

func mapUse() {
	map1 := make(map[string]int)
	map1["1"] = 1
	map1["2"] = 1
	fmt.Println("map 赋值：", map1)

	var map2 map[string]int
	fmt.Println("创建一个nil map，用于返回空值：", map2)

	value, exists := map1["1"]
	if exists {
		fmt.Println("获取map的值：", value)
	}

	value2 := map1["1"]
	if value != 0 {
		fmt.Println("获取map的值：", value2)
	}
}

func mapLoop() {
	map1 := make(map[string]int)
	map1["1"] = 1
	map1["2"] = 1
	for key, value := range map1 {
		fmt.Println(key, value)
	}
}

func removeMap() {
	map1 := make(map[string]int)
	map1["1"] = 1
	map1["2"] = 1
	reoveMap(map1, "2")
	fmt.Println(map1)
}

func reoveMap(a map[string]int, key string) {
	delete(a, key)
}
