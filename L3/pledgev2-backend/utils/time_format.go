package utils

import (
	"strconv"
	"time"
)

const (
	TimeOffset    = 8 * 3600
	HalfDayOffset = 12 * 3600
)

func GetCurrentTimestampBySecond() int64 {
	return time.Now().Unix()
}

// time.Unix(sec, nsec)
func UnixSecondToTime(second int64) time.Time {
	return time.Unix(second, 0)
}

func UnixNanoSecondToTime(second int64) time.Time {
	return time.Unix(0, second)
}

func GetCurrentTimestampByNano() int64 {
	return time.Now().UnixNano()
}

func GetCurrentTimestampByMill() int64 {
	return GetCurrentTimestampByNano() / 1e6
}

// 时间戳秒
func GetCurDayZeroTimestamp() int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timeStr)
	return t.Unix() - TimeOffset //时区
}

// 当天中午12点
func GetCurDayHalfTimestamp() int64 {
	return GetCurDayZeroTimestamp() + HalfDayOffset
}

func GetCurDayHalfTimeFormat() string {
	return time.Unix(GetCurDayZeroTimestamp(), 0).Format("2006-01-02_15-04-05")
}

// 时间字符串转时间戳
func GetTimeStampByFormat(datetime string) string {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	timestamp := tmp.Unix()
	return strconv.FormatInt(timestamp, 10)
}

// 将格式化的时间字符串转成时间戳
func TimeStringFormatTimeUnix(timeFormat string, timeSrc string) int64 {
	tm, _ := time.Parse(timeFormat, timeSrc)
	return tm.Unix()
}

// 获取当前时间
func GetCurDateTimeFormat() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
