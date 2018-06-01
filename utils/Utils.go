package utils

import (
	"time"
	"fmt"
	"github.com/SpecterTeam/SpecterGO"
	"os"
	"path/filepath"
)

func HandleError(err error){
	SpecterGO.GetLogger().Error(err.Error())
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

func GetServerPath() string {
	ex,_ := os.Executable()

	return filepath.Dir(ex)
}

func FileExists(file string) bool {
	if _,err := os.Stat(file); err != nil {
		return false
	} else {
		return true
	}
}
