package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Nickname  *string `json:"nickname"`
}

func main() {
	userPayload := []byte(`{"firstName": "Arnaud", "lastName": "Lasnier"}`)
	var user User
	json.Unmarshal(userPayload, &user)
	length := len(*user.Nickname) // nil pointer dereference
	fmt.Println(length)
}
