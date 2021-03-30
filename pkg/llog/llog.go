package llog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	DefaultCallerDepth = 2
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

type LLog struct {
	file     *os.File
	filepath string      //日志文件路径
	debug    bool        //是否开启Debug ， 如果开启debug则打印debug日志
	logger   *log.Logger //日志对象
}

func New(filepath string, debug bool) *LLog {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(fmt.Sprintf("打开日志文件 %v 时发生错误,错误原因:%v", filepath, err))
	}
	logger := log.New(f, "", log.LstdFlags)
	return &LLog{filepath: filepath, file: f, debug: debug, logger: logger}
}

func (l *LLog) setPrefix(level Level) {
	var logPrefix string
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	l.logger.SetPrefix(logPrefix)
}

func (l *LLog) Info(v ...interface{}) {
	l.setPrefix(INFO)
	l.logger.Println(v)
}

func (l *LLog) Debug(v ...interface{}) {
	if l.debug {
		l.setPrefix(DEBUG)
		l.logger.Println(v)
	}
}

func (l *LLog) Warn(v ...interface{}) {
	l.setPrefix(WARNING)
	l.logger.Println(v)
}

func (l *LLog) Error(v ...interface{}) {
	l.setPrefix(ERROR)
	l.logger.Println(v)
}

func (l *LLog) Faltal(v ...interface{}) {
	l.setPrefix(FATAL)
	l.logger.Fatal(v)
}

func (l *LLog) Close() {
	l.file.Close()
}
