package tm

import "time"

func UnixNano() int64 {
	return time.Now().UnixNano()
}
