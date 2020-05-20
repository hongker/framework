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

// GetLocalDateStr 返回日期，如： 2019-12-30
func GetLocalDateStr() string {
	return GetLocalTime().Format(DateFormat)
}

// GetLocalTimeStr 返回时间，如:2019-12-30 22:00:00
func GetLocalTimeStr() string {
	return GetLocalTime().Format(TimeFormat)
}

// GetLocalTimeStamp 返回时间戳
func GetLocalTimeStamp() int64 {
	return GetLocalTime().Unix()
}

// GetLocalMicroTimeStampStr 返回ms
func GetLocalMicroTimeStampStr() string {
	return fmt.Sprintf("%.6f", float64(GetLocalTime().UnixNano())/1e9)
}
