package main

import (
	"encoding/json"
	"fmt"

	"github.com/aarondl/opt/null"
)

type User struct {
	FirstName string           `json:"firstName"`
	LastName  string           `json:"lastName"`
	Nickname  null.Val[string] `json:"nickname"`
}

func main() {
	userPayload := []byte(`{"firstName": "Arnaud", "lastName": "Lasnier"}`)
	var user User
	json.Unmarshal(userPayload, &user)
	nickname, ok := user.Nickname.Get()
	fmt.Printf("%q %v\n", nickname, ok)
}
