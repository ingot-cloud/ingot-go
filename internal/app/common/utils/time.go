package utils

import (
	"time"
)

// TimeToStr return string
func TimeToStr(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

// TimeToDateStr 日期
func TimeToDateStr(time time.Time) string {
	return time.Format("2006-01-02")
}

// StrToTime return time
func StrToTime(traget string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", traget)
}

// StrToDate return time
func StrToDate(target string) (time.Time, error) {
	return time.Parse("2006-01-02", target)
}

// DayDiff 时间间隔，单位天
func DayDiff(start, end string) (int64, error) {
	timeLayout := "2006-01-02"
	loc, _ := time.LoadLocation("Local")
	startUnix, err := time.ParseInLocation(timeLayout, start, loc)
	if err != nil {
		return 0, err
	}
	endUnix, err := time.ParseInLocation(timeLayout, end, loc)
	if err != nil {
		return 0, err
	}
	startTime := startUnix.Unix()
	endTime := endUnix.Unix()
	date := (endTime - startTime) / 86400
	return date, nil
}
