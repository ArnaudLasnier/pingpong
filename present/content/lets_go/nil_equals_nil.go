package main

import (
	"fmt"
)

func main() {
	fmt.Println(nil == nil) // compile time error: cannot compare two untyped nil
}
