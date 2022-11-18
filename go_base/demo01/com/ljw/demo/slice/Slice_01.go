package main

import (
	"fmt"
)

func main() {
	//切片创建
	SliceCreate()
	//切片使用
	SliceUse()
	//切片循环
	SliceLoop()
	//多维切片
	MultidimensionalSlice()
	//切片作为参数传递
	SliceBigData(make([]int, 1e6))
}

func SliceCreate() {
	slice1 := make([]string, 5)
	fmt.Println("只指定长度，不指定容量：", slice1, len(slice1))

	slice2 := make([]string, 3, 5)
	fmt.Println("指定长度和容量：", slice2, len(slice2))

	slice3 := []string{"1", "2", "3"}
	fmt.Println("不使用make函数创建，[]中没有值代表创建切片。有值代表数组：", slice3, len(slice3))

	slice4 := []string{99: "100"}
	fmt.Println("初始化指定索引的值：", slice4[99])

	var nilSlice1 []int
	nilSlice2 := make([]int, 0)
	nilSlice3 := []int{}
	fmt.Println("nil 空切片 不会分配内存：", nilSlice1, nilSlice2, nilSlice3)
}

func SliceUse() {
	slice := []int{1, 2, 3, 4, 5}
	slice[1] = 10
	fmt.Println("切片初始化赋值：", slice)

	newSlice := slice[1:5]
	fmt.Println("使用同一个内存地址：", newSlice)

	newSlice[0] = 1
	fmt.Println("修改共享内存数据：", slice)

	newSlice = append(newSlice, 100)

	fmt.Println("扩展超过容量后使用新的内存数据：", newSlice)

	newSlice[0] = 1000
	fmt.Println("扩展超过容量后使用新的内存数据：", newSlice)
	fmt.Println("修改扩容后的数组，原数组不变", slice)

	newSlice2 := newSlice[1:2:2]
	fmt.Println("指定长度和容量", newSlice2, len(newSlice2), cap(newSlice2))

}

func SliceLoop() {
	slice := []int{1, 2, 3, 4, 5}
	for index, value := range slice {
		fmt.Println(index, value)
	}

	for _, value := range slice {
		fmt.Println(value)
	}

	for index := 2; index < len(slice); index++ {
		fmt.Println(index, slice[index])
	}
}

func MultidimensionalSlice() {
	slice := [][]int{{10}, {10, 200}}
	fmt.Println("多维数组：", slice)
	slice[0] = append(slice[0], 20)
	fmt.Println("多维数组赋值：", slice)
}

func SliceBigData(a []int) {
	fmt.Println(len(a))
}
