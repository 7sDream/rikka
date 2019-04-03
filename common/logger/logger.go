package logger

import (
	"fmt"
	"log"
	"os"
)

// Logger is a util class to print log in different level like DEBUG.
// It has 4 level: DEBUG, INFO, WARN, ERROR,
// and 2 exception logger: PANIC and FATAL
type Logger struct {
	prefix     string
	trueLogger *log.Logger
}

// Logger levels
const (
	LevelDebug int = iota
	LevelInfo
	LevelWarn
	LevelError
)

var (
	l = NewLogger("[Logger]")

	currentLevel = LevelDebug
)

// NewLogger create a new top level logger based on prefix.
func NewLogger(prefix string) *Logger {
	flag := log.Ldate | log.Ltime | log.Lmicroseconds
	return &Logger{
		prefix:     prefix,
		trueLogger: log.New(os.Stdout, "", flag),
	}
}

// Debug print log message as DEBUG level.
// if you set log level higher than LevelDebug, no message will be print.
func (logger *Logger) Debug(data ...interface{}) {
	if currentLevel <= LevelDebug {
		logger.trueLogger.Print("[DEBUG] ", logger.prefix, " ", fmt.Sprintln(data...))
	}
}

// Info print log message as INFO level.
// If you set log level higher than LevelInfo, no message will be print.
func (logger *Logger) Info(data ...interface{}) {
	if currentLevel <= LevelInfo {
		logger.trueLogger.Print("[ Info] ", logger.prefix, " ", fmt.Sprintln(data...))
	}
}

// Warn print log message as WARN level.
// If you set log level higher than LevelWarn, no message will be print.
func (logger *Logger) Warn(data ...interface{}) {
	if currentLevel <= LevelWarn {
		logger.trueLogger.Print("[ WARN] ", logger.prefix, " ", fmt.Sprintln(data...))
	}
}

// Error print log message as ERROR level. 
// This function do not create panic or fatal, it just print error message.
// If you want get a runtime panic or fatal, use Logger.Panic or Logger.Fatal instead.
func (logger *Logger) Error(data ...interface{}) {
	if currentLevel <= LevelError {
		logger.trueLogger.Print("[ERROR] ", logger.prefix, " ", fmt.Sprintln(data...))
	}
}

// Panic print log message, and create a panic use the message.
func (logger *Logger) Panic(data ...interface{}) {
	logger.trueLogger.Panic("[PANIC]", logger.prefix, " ", fmt.Sprint(data...))
}

// Fatal print log message, and create a fatal use the message.
func (logger *Logger) Fatal(data ...interface{}) {
	logger.trueLogger.Fatal("FATAL", logger.prefix, " ", fmt.Sprint(data...))
}

// SubLogger create a new logger based on the logger.
// Prefix string of new logger will be concat of old and provided argument.
func (logger *Logger) SubLogger(prefix string) (subLogger *Logger) {
	return NewLogger(fmt.Sprintf("%s %s", logger.prefix, prefix))
}

// SetLevel set the minimum level that message will be print out
func SetLevel(level int) {
	if LevelDebug <= level && level < LevelError {
		currentLevel = level
	} else {
		l.Error("Set logger level", level, "failed, accepted range is", LevelInfo, "to", LevelError)
	}
}
