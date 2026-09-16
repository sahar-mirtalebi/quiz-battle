package timestamp

import "time"

func Now() uint64 {
	return uint64(time.Now().UnixMicro())
}

func Add(d time.Duration) int64 {
	return time.Now().Add(d).UnixMicro()
}

func ToTime(value int64) time.Time {
	return time.UnixMicro(value)
}
