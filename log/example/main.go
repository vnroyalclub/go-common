package main

import (
	"github.com/vnroyalclub/go-common/log"
	"time"
)

func main() {
	str := "######### hello world"
	log.Debug(str)
	log.Info(str)
	log.Warn(str)
	log.Error(str)
	log.Critical(str)

	str = "********** hello world"
	log.SetLogLevel(5) //warming 级别
	log.Debug(str)
	log.Info(str)
	log.Warn(str)
	log.Error(str)
	log.Critical(str)

	//example中，这是为了等待日志输出完整
	time.Sleep(time.Second)
}
