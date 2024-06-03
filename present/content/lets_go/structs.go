package main

import (
	"fmt"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (user *User) FullName() string {
	return user.FirstName + " " + user.LastName
}

func main() {
	user := User{FirstName: "Arnaud", LastName: "Lasnier"}
	p := &user
	p.FirstName, p.LastName = "Rob", "Pike"
	fmt.Println(user.FullName())
}
