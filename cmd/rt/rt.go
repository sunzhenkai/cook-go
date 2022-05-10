package main

import (
	"cooking-go/pkg/rt"
	"time"
)

func main() {
	go rt.Task()
	go rt.Task()
	println("YES")

	time.Sleep(3 * time.Second)
	println("DONE")
}
