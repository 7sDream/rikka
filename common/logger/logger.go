package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	Prefix string
	il     *log.Logger
	wl     *log.Logger
	el     *log.Logger
	pl     *log.Logger
	fl     *log.Logger
}

const (
	LevelInfo = iota
	LevelWarn
	LevelError
)

var l = NewLogger("[Logger]")

var currentLevel = LevelInfo

func NewLogger(prefix string) *Logger {
	return &Logger{
		Prefix: prefix,
		il:     log.New(os.Stdout, fmt.Sprintf("[%s] %s ", "INFO", prefix), log.LstdFlags),
		wl:     log.New(os.Stdout, fmt.Sprintf("[%s] %s ", "WARN", prefix), log.LstdFlags),
		el:     log.New(os.Stdout, fmt.Sprintf("[%s] %s ", "ERROR", prefix), log.LstdFlags),
		pl:     log.New(os.Stdout, fmt.Sprintf("[%s] %s ", "PANIC", prefix), log.LstdFlags),
		fl:     log.New(os.Stdout, fmt.Sprintf("[%s] %s ", "FATAL", prefix), log.LstdFlags),
	}
}

func (this *Logger) Info(data ...interface{}) {
	if currentLevel <= LevelInfo {
		this.il.Println(data)
	}
}

func (this *Logger) Warn(data ...interface{}) {
	if currentLevel <= LevelWarn {
		this.wl.Println(data)
	}
}

func (this *Logger) Error(data ...interface{}) {
	if currentLevel <= LevelError {
		this.el.Println(data)
	}
}

func (this *Logger) Panic(data ...interface{}) {
	this.pl.Panicln(data)
}

func (this *Logger) Fatal(data ...interface{}) {
	this.fl.Fatalln(data)
}

func (this *Logger) SubLogger(prefix string) (subLogger *Logger) {
	subLogger = NewLogger(fmt.Sprintf("%s %s", this.Prefix, prefix))
	return subLogger
}

func SetLevel(level int) {
	if LevelInfo <= level && level < LevelError {
		currentLevel = level
	} else {
		l.Warn("Set logger level", level, "failed, accepted range is", LevelInfo, "to", LevelError)
	}
}
