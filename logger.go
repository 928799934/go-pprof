package pprof

import "log"

var logger *log.Logger

// SetErrorLog 设置输出
func SetErrorLog(l *log.Logger) {
	logger = l
}

func logf(format string, args ...interface{}) {
	if logger != nil {
		logger.Printf(format, args...)
		return
	}
	log.Printf(format, args...)
}
