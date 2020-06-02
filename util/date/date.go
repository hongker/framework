package date

import (
	"fmt"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
	DateFormat = "2006-01-02"
)

// GetLocalTime 返回东8时区的时间
func GetLocalTime() time.Time {
	return time.Now().In(GetLocalTimeZone())
}

// GetLocalTimeZone
func GetLocalTimeZone() *time.Location {
	return time.FixedZone("CST", 8*3600) // UTC+8
}

// GetLocalMicroTimeStampStr 返回ms
func GetLocalMicroTimeStampStr() string {
	return fmt.Sprintf("%.6f", float64(GetLocalTime().UnixNano())/1e9)
}

// FormatTimeInLocation 格式化时间戳
func TimestampInLocation(timestamp int64, loc *time.Location) time.Time  {
	return time.Unix(timestamp, 0).In(loc)
}

// ParseTimeInLocation 格式化时间
func ParseTimeInLocation(timeStr string, loc *time.Location) time.Time  {
	tmp, _ := time.ParseInLocation(TimeFormat, timeStr, loc)
	return tmp
}
