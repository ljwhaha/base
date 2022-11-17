package main

import (
	"fmt"
)

func main() {
	fmt.Println("数组创建")
	ArrayCreate()
	fmt.Println("===========================")
	fmt.Println("数组复制")
	ArrayCopy()
	fmt.Println("===========================")
	fmt.Println("多维数组")
	MultidimensionalArray()
	fmt.Println("===========================")
	fmt.Println("大数据量数组传值")
	var array [1e6]int
	bigDataArray(array)
	fmt.Println("===========================")
	fmt.Println("大数据量数组传值优化，使用指针数组，减少内存开销，但是共享内存")
	bigDataPointerArray(&array)
}

func ArrayCreate() {
	var arrayEmpty [5]int
	fmt.Println("声明一个数组：", arrayEmpty)
	arrayInit := [5]int{1, 2, 3, 4, 5}
	fmt.Println("初始化一个数组，长度为初始化数值数量：", arrayInit)
	fmt.Println("索引下标为1的值：", arrayInit[1])
	arrayInit[1] = 10
	fmt.Println("设置下标为1的值为10：", arrayInit[1])
	pointerArray := [5]*int{0: new(int), 2: new(int)}
	fmt.Println("指针数组：可初始化执行索引下标的值：", pointerArray)
	*pointerArray[0] = 1
	fmt.Println("初始化指定下标后，可赋值，*标明具体指针的值", *pointerArray[0])
}

func ArrayCopy() {
	arrayInit := [5]int{1, 2, 3, 4, 5}
	var copyArray [5]int
	copyArray = arrayInit
	fmt.Println("复制后的数组：", copyArray)

	pointerArray := [2]*string{0: new(string), 1: new(string)}
	fmt.Println("初始化指针数组：", pointerArray)
	fmt.Println("初始化指针数组的值：", *pointerArray[0])
	var pointerArrayCopy [2]*string
	pointerArrayCopy = pointerArray
	fmt.Println("复制后的指针数组，只复制指针，不复制值：", pointerArrayCopy)
	*pointerArrayCopy[0] = "1"
	fmt.Println("修改复制后的指针数组的值，原数组的值也会发生变化：", *pointerArray[0])
}

/**
多维数组
*/
func MultidimensionalArray() {
	arrayInit := [2][2]int{{1, 1}, {2, 2}}
	fmt.Println("多维数组初始化：", arrayInit)
	var arrayInitCopy [2][2]int
	arrayInitCopy = arrayInit
	fmt.Println("相同纬度和类型的值进行复制：", arrayInitCopy)
	arrayInitCopy[0][0] = 10
	fmt.Println("修改数组的值：", arrayInitCopy)

	var array1 = arrayInitCopy[0]
	fmt.Println("多维数组的一个赋值给单纬数组：", array1)

}

/**
大数据量数组传值
*/
func bigDataArray(array [1e6]int) {
	fmt.Println("数组长度", len(array))
}

func bigDataPointerArray(array *[1e6]int) {
	fmt.Println("数组长度", len(array))
}
