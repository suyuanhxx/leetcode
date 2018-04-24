package main

import "fmt"

type People interface {
	getUserName() string
}

type User struct {
	Name string
}

func (user *User) getUserName() string {
	return user.Name
}

func main() {
	var p People
	p = &User{Name: "111"}
	fmt.Print(p.getUserName())
}
