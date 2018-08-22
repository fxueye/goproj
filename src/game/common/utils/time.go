package utils

import (
	"strings"
	"time"
)

func WebTime(t time.Time) string {
	ftime := t.Format(time.RFC1123)
	if strings.HasSuffix(ftime, "UTC") {
		ftime = ftime[0:len(ftime)-3] + "GMT"
	}
	return ftime
}
func NowTimestamp() int64 {
	return time.Now().Unix()
}
func FormatTimestamp(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}
