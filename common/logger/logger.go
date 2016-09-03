package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	Prefix string
	il     *log.Logger
	el     *log.Logger
	pl     *log.Logger
	fl     *log.Logger
}

func NewLogger(prefix string) *Logger {
	return &Logger{
		Prefix: prefix,
		il:     log.New(os.Stdout, fmt.Sprintf("[%s] %s ", "INFO", prefix), log.LstdFlags),
		el:     log.New(os.Stdout, fmt.Sprintf("[%s] %s ", "ERROR", prefix), log.LstdFlags),
		pl:     log.New(os.Stdout, fmt.Sprintf("[%s] %s ", "PANIC", prefix), log.LstdFlags),
		fl:     log.New(os.Stdout, fmt.Sprintf("[%s] %s ", "FATAL", prefix), log.LstdFlags),
	}
}

func (this *Logger) Info(data ...interface{}) {
	this.il.Println(data)
}

func (this *Logger) Error(data ...interface{}) {
	this.el.Println(data)
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
