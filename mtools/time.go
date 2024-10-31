package mtools

import "time"

// TimeUnix time类型转时间戳
func Time2Unix(t time.Time) int64 {
	if t.IsZero() {
		return 0
	}
	return t.Unix()
}

// GetZeroTime 获取0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// Get24Time 获取24点时间
func Get24Time(d time.Time) time.Time {
	return GetZeroTime(d).Add(time.Hour * 24)
}
