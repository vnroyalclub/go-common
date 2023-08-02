package log

// go get github.com/yougg/log4go

import (
	"fmt"
	"strings"

	"github.com/yougg/log4go"
)

var (
	logger = make(log4go.Logger)
)

func init() {
	//读入日志配置
	logWriter := log4go.NewConsoleLogWriter()
	logWriter.SetFormat("[%T] [%L] (%I) (%s) %M ")
	log4go.EnableGoRoutineID = true
	log4go.FuncCallDepth = 3
	logger.AddFilter("stdout", log4go.DEBUG, logWriter)
}

// 日志级别 FINE DEBUG TRACE INFO WARNING ERROR CRITICAL （1--7）
func SetLogLevel(level int) {
	logger["stdout"].Level = log4go.Level(level)
}

func Debug(args ...interface{}) {
	logger.Debug(format(args...))
}

func Info(args ...interface{}) {
	logger.Info(format(args...))
}

func Warn(args ...interface{}) {
	logger.Warn(format(args...))
}

func Error(args ...interface{}) {
	logger.Error(format(args...))
}

func Critical(args ...interface{}) {
	logger.Critical(format(args...))
}

func format(args ...interface{}) (msg string) {
	msg = fmt.Sprintf(strings.Repeat(" %v", len(args)), args...)
	return
}
