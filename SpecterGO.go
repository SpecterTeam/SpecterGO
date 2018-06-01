package SpecterGO

import (
	"github.com/SpecterTeam/SpecterGO/utils"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

const (
	Name = "SpecterGO"
	Version = "0.1-ALPHA"
)

var (
	logger = utils.NewLogger()
	server Server
)

func GetLogger() utils.Logger {
	return logger
}

func SetServer(s Server) {
	server = s
}

func GetServer() Server {
	return server
}

func Start() {
	InitConfig()
	s := NewServer()
	SetServer(s)
}

func InitConfig() {
	if utils.FileExists(utils.GetServerPath() + "server.yml") {
		bytes, err := ioutil.ReadFile(utils.GetServerPath() + "server.yml")
		if err != nil {
			utils.HandleError(err)
		} else {
			yml := new(interface{})
			yaml.Unmarshal(bytes, &yml)
			//TODO: utils/Config.go
		}
	}
}