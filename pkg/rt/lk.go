package rt

import (
	"fmt"
	"sync"
	"time"
)

func TLock() {
	var l sync.Mutex
	var cnt = 10000
	go Run(&cnt, &l)
	go Run(&cnt, &l)
	go Run(&cnt, &l)
	go Run(&cnt, &l)
	go Run(&cnt, &l)
	go Run(&cnt, &l)

	go PrintCnt(&cnt)
}

func PrintCnt(cnt *int) {
	for true {
		fmt.Printf("count number: %d\n", *cnt)
		time.Sleep(1 * time.Second)
	}
}

func Run(cnt *int, l *sync.Mutex) {
	for *cnt > 0 {
		l.Lock()
		*cnt--
		l.Unlock()
		RTask()
	}
}

func RTask() {
	time.Sleep(100 * time.Millisecond)
}
