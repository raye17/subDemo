package times

import "time"

func NowTime() string {
	t := time.Now()
	// 格式化时间
	return t.Format("2006-01-02 15:04:05")
}
