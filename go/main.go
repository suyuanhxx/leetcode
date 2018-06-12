package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

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

	//head := BuildDoubleLinkList([]int{3, 22, 5, 45, 23, 8})
	//head.Print()

	baseurl := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxf3d2819a2bc62b29&secret=2ebca847f109b84ce0476788d076aca6";
	resp, e := http.Get(baseurl)
	if e != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(body, string(body))

}
