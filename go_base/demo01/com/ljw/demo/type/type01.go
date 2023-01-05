package main

import (
	_ "demo01/com/ljw/demo/type"
	"fmt"
)

func main() {
	createUser()
	createAdmin()

	lisa := user{
		name:       "Lisa",
		email:      "lisa@qq.com",
		ext:        123,
		privileged: true,
	}

}

func createUser() {
	var bill user
	fmt.Println(bill)

	lisa := user{
		name:       "Lisa",
		email:      "lisa@qq.com",
		ext:        123,
		privileged: true,
	}

	fmt.Println("创建一个用户lisa：", lisa)

	kk := user{"kk", "kk@qq.com", 123, true}
	fmt.Println("创建一个用户KK：", kk)

}

func createAdmin() {
	fred := admin{person: user{
		name:       "jane",
		email:      "jane@qq.com",
		ext:        123,
		privileged: true,
	}, level: "1"}

	fmt.Println("创建一个admin : ", fred)
}

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

type admin struct {
	person user
	level  string
}

type student user
