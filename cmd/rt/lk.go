package main

import (
	"evaluate-go/pkg/rt"
	"time"
)

func main() {
	rt.TLock()

	time.Sleep(1 * time.Minute)
}
