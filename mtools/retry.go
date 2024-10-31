package mtools

import (
	"errors"
	"time"
)

// 默认重试周期0s, 1s, 2s, 4s, 8s
func DoRetry(maxRetryTime int, callFunc func() bool) (err error) {
	i := 1
	sleep := time.Second * 1
	for {
		if i > maxRetryTime {
			return errors.New("retry too many times")
		}

		if ok := callFunc(); !ok {
			time.Sleep(sleep)
			sleep = sleep * 2
			i++
		} else {
			break
		}
	}
	return
}
