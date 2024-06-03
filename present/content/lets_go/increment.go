package main

import "fmt"

func incrementIntCopy(x int) {
	x++
}

func incrementIntThroughPtr(p *int) {
	(*p)++
}

func main() {
	// START OMIT
	a, b := 0, 0
	p := &b
	incrementIntCopy(a)       // takes int variable `x` and performs `x++`
	incrementIntThroughPtr(p) // takes *int variable `p` and performs `(*p)++`
	fmt.Println(a, b)
	// END OMIT
}
