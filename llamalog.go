package llamalog

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Logger struct {
	packageName string
	writer      io.Writer
}

type LogType string

var (
	Error LogType = SeqRed + "[ERR]" + SeqReset
	Warn  LogType = SeqYellow + "[WRN]" + SeqReset
	Info  LogType = SeqGreen + "[INF]" + SeqReset
)

func NewLogger(packageDir ...string) *Logger {
	return NewLoggerFromWriter(os.Stderr, packageDir...)
}

func NewLoggerFromWriter(writer io.Writer, packageDir ...string) *Logger {
	return &Logger{packageName: Blue("[" + strings.Join(packageDir, " > ") + "] >>>"), writer: writer}
}

func (log *Logger) Log(logType LogType, format string, a ...interface{}) {
	fmt.Fprintln(log.writer, logType, Magenta(time.Now().Format("[15:04]")), log.packageName, fmt.Sprintf(format, a...))
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
