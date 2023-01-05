package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//分配一个逻辑处理器
	runtime.GOMAXPROCS(1)

	//设置等待程序
	var wg sync.WaitGroup

	wg.Add(2)

	fmt.Println("Satrt...........")

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("waiting..............")

	//等待
	wg.Wait()
	fmt.Println("")
	fmt.Println("end..............")

}
