package tm

import "time"

func UnixNano() int64 {
	return time.Now().UnixNano()
}

func UnixMilli() int64 {
	return time.Now().UnixNano() / 1000000
}
