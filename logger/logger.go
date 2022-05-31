package logger

import (
	"io"
	"log"
)

type Logger struct {
	Error *log.Logger
	Warn  *log.Logger
	Info  *log.Logger
	Debug *log.Logger
}

func New(out, errOut io.Writer) *Logger {
	flag := log.Ldate | log.Ltime | log.Lshortfile

	return &Logger{
		Error: log.New(errOut, "ERROR ", flag),
		Warn:  log.New(out, "WARN ", flag),
		Info:  log.New(out, "INFO ", flag),
		Debug: log.New(out, "DEBUG ", flag),
	}
}
