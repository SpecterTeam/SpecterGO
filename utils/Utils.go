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
	"time"
	"fmt"
	"os"
	"path/filepath"
	"bytes"
	"encoding/gob"
)

const (
	DirectorySeparator = "/"
)

var (
	logger Logger
)

func HandleError(err error){
	logger.Error(err.Error())
}

func GetTimeString() string {
	sec := IntToString(time.Now().Second())
	min := IntToString(time.Now().Minute())
	hour := IntToString(time.Now().Hour())

	return hour + ":" + min + ":" + sec
}

func IntToString(int int) string {
	return fmt.Sprintf("%d", int)
}

func InterfaceToString(i interface{}) string {
	return i.(string)
}

func InterfaceToInt(i interface{}) int {
	return i.(int)
}

func ArrayToMap(array []string) map[int]string {
	m := make(map[int]string)
	for i,s := range array {
		m[i] = s
	}

	return m
}

func MapToArray(m map[int]string) []string {
	array := make([]string, 0)
	for i,s := range m {
		array[i] = s
	}

	return array
}

func GetServerPath() string {
	ex,_ := os.Executable()

	return filepath.Dir(ex) + DirectorySeparator
}

func FileExists(file string) bool {
	if _,err := os.Stat(file); err != nil {
		return false
	} else {
		return true
	}
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
