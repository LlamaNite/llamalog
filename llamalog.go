package llamalog

import (
	"fmt"
	"strings"
	"time"
)

type Logger struct {
	packageName string
}

type LogType string

var (
	Error LogType = SeqRed + "[ERR]" + SeqReset
	Warn  LogType = SeqYellow + "[WRN]" + SeqReset
	Info  LogType = SeqGreen + "[INF]" + SeqReset
)

func NewLogger(packageDir ...string) *Logger {
	return &Logger{Blue("[" + strings.Join(packageDir, " > ") + "] >>>")}
}

func (log *Logger) Log(logType LogType, format string, a ...interface{}) {
	fmt.Println(logType, Magenta(time.Now().Format("[15:04]")), log.packageName, fmt.Sprintf(format, a...))
}

func (log *Logger) Error(format string, a ...interface{}) {
	log.Log(Error, format, a...)
}

func (log *Logger) Warn(format string, a ...interface{}) {
	log.Log(Warn, format, a...)
}
func (log *Logger) Info(format string, a ...interface{}) {
	log.Log(Info, format, a...)
}
