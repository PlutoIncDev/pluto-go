package logging

import "log"

type LogLevel int

const (
	DebugLogLevel LogLevel = iota
	InfoLogLevel
	WarningLogLevel
	ErrorLogLevel
)

var currentLogLevel = InfoLogLevel

func SetLogLevel(level LogLevel) {
	currentLogLevel = level
}

func Debug(v ...interface{}) {
	if currentLogLevel > DebugLogLevel {
		return
	}
	log.Println("[DEBUG]:", v)
}

func Info(v ...interface{}) {
	if currentLogLevel > InfoLogLevel {
		return
	}
	log.Println(" [INFO]:", v)
}

func Warning(v ...interface{}) {
	if currentLogLevel > WarningLogLevel {
		return
	}
	log.Println(" [WARN]:", v)
}

func Error(v ...interface{}) {
	if currentLogLevel > ErrorLogLevel {
		return
	}
	log.Println("[ERROR]:", v)
}

func Fatal(v ...interface{}) {
	log.Fatalln("[FATAL]:", v)
}