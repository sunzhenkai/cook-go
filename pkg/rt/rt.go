package rt

import (
	"evaluate-go/pkg/tm"
	"fmt"
	"time"
)

func Task() {
	time.Sleep(1 * time.Second)
	fmt.Printf("one task done at %d\n", tm.UnixMilli())
}
