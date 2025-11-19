package llog

import "C"
import (
	"fmt"
	"log"
	"os"
	"time"

	"gitee.com/liumou_site/logger"
)

type flogfunc func(*log.Logger)

type clogfunc func(*logger.LocalLogger)

type Log struct {
	LogPath  string
	LogLevel string
	CLog     *logger.LocalLogger
	FLog     *log.Logger
	ToDay    string
}

func (l *Log) fileLog(f flogfunc) {
	logFile := l.fileIO()
	defer logFile.Close()

	l.FLog.SetOutput(logFile)
	f(l.FLog)
}

func (l *Log) consoleLog(f clogfunc) {
	f(l.CLog)
}

func (l *Log) fileIO() *os.File {
	date := time.Now().Format("20060102")

	if date != l.ToDay {
		l.ToDay = date
	}

	logFile, _ := os.OpenFile(fmt.Sprintf("%s/%s.log", l.LogPath, l.ToDay),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666)

	return logFile
}

func (l *Log) MyINFO(info_ string) {

	l.fileLog(func(l *log.Logger) {
		l.Printf(fmt.Sprintf("[INFO] %v", info_))
	})

	l.consoleLog(func(l *logger.LocalLogger) {
		l.Info(info_)
	})
}

func (l *Log) MyDEBUG(debug_ string) {

	l.fileLog(func(l *log.Logger) {
		l.Printf(fmt.Sprintf("[DEBUG] %v\n", debug_))
	})

	l.consoleLog(func(l *logger.LocalLogger) {
		l.Debug(debug_)
	})
}

func (l *Log) MyWarning(warning_ string) {

	l.fileLog(func(l *log.Logger) {
		l.Printf(fmt.Sprintf("[WARNING] %v\n", warning_))
	})

	l.consoleLog(func(l *logger.LocalLogger) {
		l.Warn(warning_)
	})
}

func (l *Log) MyError(error_ string) {

	l.fileLog(func(l *log.Logger) {
		l.Printf(fmt.Sprintf("[ERROR] %v\n", error_))
	})

	l.consoleLog(func(l *logger.LocalLogger) {
		l.Error(error_)
	})
}
