/**
 *     SpecterGO  Copyright (C) 2018  SpecterTeam
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package utils

import (
	"os"
	"fmt"
)

type Logger struct {
}

var (
	LogFile string
)

const (
	Prefix      = "[SpecterGO] "
	ErrorPrefix = "[Error] "
	NewLine     = "\n"
)

func NewLogger() Logger {
	l := Logger{}
	os.Mkdir(GetServerPath() + "/logs", 0777)
	SetLogFile(GetServerPath() + "logs/log-" + GetTimeString() + ".txt")
	logger = l
	return l
}

func SetLogFile(path string) {
	LogFile = path
	if !FileExists(path) {
		f, err := os.Create(path)
		defer f.Close()
		if err != nil {
			HandleError(err)
		}
	}
}

func (l *Logger) Log(log string) {
	f,err := os.Open(LogFile)
	defer f.Close()
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