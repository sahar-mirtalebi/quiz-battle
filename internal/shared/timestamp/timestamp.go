package timestamp

import "time"

func Now() uint64 {
	return uint64(time.Now().UnixMicro())
}
