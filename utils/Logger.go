package utils

import (
	"os"
	"fmt"
)

type Logger struct {
	LogFile string
}

const (
	Prefix      = "[SpecterGO] "
	ErrorPrefix = "[Error] "
	NewLine     = "\n"
)

func NewLogger() Logger {
	l := Logger{}
	os.Mkdir(GetServerPath() + "/logs", 0777)
	l.SetLogFile(GetServerPath() + "/logs/log-" + GetTimeString() + ".txt")

	return l
}

func(l* Logger) SetLogFile(path string) {
	l.LogFile = path
	if !FileExists(path) {
		f, err := os.Create(path)
		defer f.Close()
		if err != nil {
			HandleError(err)
		}
	}
}

func (l *Logger) GetLogFile() string {
	return l.LogFile
}

func (l *Logger) Log(log string) {
	f,err := os.Open(l.GetLogFile())
	if err != nil {
		HandleError(err)
	} else {
		f.WriteString(log + NewLine)
	}
}

func (l *Logger) Info(log string) {
	time := "[" + GetTimeString() + "] "
	fmt.Println(Prefix + time + log)
	l.Log(Prefix + time + log)
}

func (l *Logger) Error(log string) {
	time := "[" + GetTimeString() + "] "
	fmt.Println(Prefix + time + ErrorPrefix + log)
	l.Log(Prefix + ErrorPrefix + time + log)
}