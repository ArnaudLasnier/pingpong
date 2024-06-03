package main

import (
	"fmt"
	"time"
)

func winPointEverySecond(player string) {
	for i := 1; true; i++ {
		fmt.Println(player+":", i)
		time.Sleep(time.Second)
	}
}

// START OMIT
func main() {
	winPointEverySecond("Arnaud")
	winPointEverySecond("Jules")
}

// END OMIT
