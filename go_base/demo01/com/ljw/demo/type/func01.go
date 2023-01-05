package main

import "fmt"

func Notify(u user) {
	fmt.Println("send email to ", u.name, " email :", u.email)
}

func ChangeEmail(u user, email string) {
	u.email = email
}
