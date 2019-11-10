package utils

import (
	"encoding/json"
	"time"
)

func MergeData(source interface{}, output interface{}) {
	bytes, _ := json.Marshal(source)
	json.Unmarshal(bytes, &output)
}

func Stringify(data interface{}) string {
	if bytes, err := json.Marshal(data); err == nil {
		return string(bytes)
	}
	return ""
}

func TimestampToTime(timestamp int64) time.Time {
	return time.Unix(0, timestamp*int64(time.Millisecond))
}
