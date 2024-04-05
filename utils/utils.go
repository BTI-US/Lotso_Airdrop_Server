package utils

import (
	"time"
)

var (
	TimeZoneName = "UTC"
)

func ZeroTime() (t time.Time) {
	t = time.Date(1970, 1, 1, 8, 0, 0, 0, time.Now().Location())
	return
}

func NextOddHourTime() (nextTime time.Time) {
	currentTime := time.Now()

	// 获取当前小时数
	currentHour := currentTime.Hour()

	// 计算下一个奇数小时
	var nextOddHour int
	if currentHour%2 == 0 {
		nextOddHour = currentHour + 1
	} else {
		nextOddHour = currentHour + 2
	}

	// 设置下一个奇数小时的时间
	nextTime = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), nextOddHour, 0, 0, 0, currentTime.Location())
	return
}
