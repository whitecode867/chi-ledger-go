package helpers

import (
	"chi-domain-go/models/utils"
	"time"
)

func GetCurrentMillisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func TimestampToTime(timestamp int64) time.Time {
	return utils.TimestampToTime(timestamp)
}
