package main

import "fmt"

func main() {
	SliceCreate()
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
