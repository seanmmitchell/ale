package tfuncs

import (
	"fmt"
	"time"
)

func TimeToString(now time.Time) string {
	return fmt.Sprintf("%02d-%02d-%04d %02d:%02d:%02d.%03d", int(now.Month()), now.Day(), now.Year(), now.Hour(), now.Minute(), now.Second(), GetMilliseconds(now.Nanosecond()))
}

func GetMilliseconds(nanoseconds int) int {
	for nanoseconds >= 1000 {
		nanoseconds = nanoseconds / 1000
	}
	return nanoseconds
}
