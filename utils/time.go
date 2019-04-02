package utils

import "time"

//NowTimestampNano 当前时间 纳秒
func NowTimestampNano() int64 {
	return time.Now().UnixNano()
}
