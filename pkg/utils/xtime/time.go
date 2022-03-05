package xtime

import "time"

var TS TimeFormat = "2006-01-02 15:04:05"

type TimeFormat string

const (
	DateFormat         = "2006-01-02"
	UnixTimeUnitOffset = uint64(time.Millisecond / time.Nanosecond)
)

func (ts TimeFormat) Format(t time.Time) string {
	return t.Format(string(ts))
}
