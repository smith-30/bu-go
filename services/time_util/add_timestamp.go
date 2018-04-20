package time_util

import (
	"time"
)

const YMDHis = "20060102150405"

var Now = func() time.Time {
	return time.Now()
}

func AddTS(str string) string {
	t := Now()
	return t.Format(YMDHis) + "_" + str
}
