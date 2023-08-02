package filelog

// go get github.com/yougg/log4go
import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yougg/log4go"
)

type FileLog struct {
	log log4go.Logger
}

func NewFileLog(logPath string, logName string) (*FileLog, error) {
	exist, err := pathExists(logPath)
	if err != nil {
		fmt.Println("[EROR] failed to path exist,err:",err)
		return nil, err
	}

	//如果文件夹不存在，则要新创建文件夹
	if !exist {
		err = os.MkdirAll(logPath, 0777)
		if err != nil {
			fmt.Println("[EROR] failed to new log dir,err:",err)
			return nil, err
		}
	}

	logName = filepath.Join(logPath, logName)

	logWriter := log4go.NewFileLogWriter(logName, false)
	logWriter.SetFormat("[%T] [%L] (%I) (%s) %M ")
	logWriter.SetRotate(true)
	logWriter.SetRotateSize(50 * 1024 * 1024)
	logWriter.SetRotateDaily(true)
	logWriter.SetMaxDays(10)

	log4go.EnableGoRoutineID = true
	log4go.FuncCallDepth = 3

	v := log4go.Logger{}
	v.AddFilter("file", log4go.DEBUG, logWriter)

	return &FileLog{log: v}, nil
}

// 日志级别 FINE DEBUG TRACE INFO WARNING ERROR CRITICAL （1--7）
func (p *FileLog) SetLogLevel(level int) {
	p.log["file"].Level = log4go.Level(level)
}

func (p *FileLog) Debug(args ...interface{}) {
	p.log.Debug(format(args...))
}

func (p *FileLog) Info(args ...interface{}) {
	p.log.Info(format(args...))
}

func (p *FileLog) Warn(args ...interface{}) {
	p.log.Warn(format(args...))
}

func (p *FileLog) Error(args ...interface{}) {
	p.log.Error(format(args...))
}

func (p *FileLog) Critical(args ...interface{}) {
	p.log.Critical(format(args...))
}

func (p *FileLog) DebugF(format string,args ...interface{}) {
	p.log.Debug(format,args...)
}

func (p *FileLog) InfoF(format string,args ...interface{}) {
	p.log.Info(format,args...)
}

func (p *FileLog) WarnF(format string,args ...interface{}) {
	p.log.Warn(format,args...)
}

func (p *FileLog) ErrorF(format string,args ...interface{}) {
	p.log.Error(format,args...)
}

func (p *FileLog) CriticalF(format string,args ...interface{}) {
	p.log.Critical(format,args...)
}

func format(args ...interface{}) (msg string) {
	msg = fmt.Sprintf(strings.Repeat(" %v", len(args)), args...)
	return
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
