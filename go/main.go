package main

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

	head := BuildDoubleLinkList([]int{3, 22, 5, 45, 23, 8})
	head.Print()
}
