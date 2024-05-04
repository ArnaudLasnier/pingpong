//go:build !dev

package main

import "github.com/ArnaudLasnier/pingpong/internal/startup"

func main() {
	app := startup.NewApplication()
	app.Start()
}
