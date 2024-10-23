package util

import (
	"fmt"
	"time"
)

// 输入年份和季度,返回该年份下该季度的范围
func GetSeasonRange(year int, season int) (time.Time, time.Time, error) {
	switch season {
	case 1:
		return time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC),
			time.Date(year, time.March, 31, 23, 59, 59, 0, time.UTC), nil
	case 2:
		return time.Date(year, time.April, 1, 0, 0, 0, 0, time.UTC),
			time.Date(year, time.June, 30, 23, 59, 59, 0, time.UTC), nil
	case 3:
		return time.Date(year, time.July, 1, 0, 0, 0, 0, time.UTC),
			time.Date(year, time.September, 30, 23, 59, 59, 0, time.UTC), nil
	case 4:
		return time.Date(year, time.October, 1, 0, 0, 0, 0, time.UTC),
			time.Date(year, time.December, 31, 23, 59, 59, 0, time.UTC), nil
	default:
		return time.Time{}, time.Time{}, fmt.Errorf("无效的季度: %d", season)
	}
}

// 返回这一年时间范围
func GetYearRange(year int) (time.Time, time.Time, error) {
	nextYear := time.Now().AddDate(1, 0, 0).Year()
	if year < 1900 || year > nextYear {
		return time.Time{}, time.Time{}, fmt.Errorf("年份不合理，必须在 1900 到当前年份的后一年之间: %d", nextYear)
	}

	return time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(year, time.December, 31, 23, 59, 59, 0, time.UTC), nil
}
