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

package SpecterGO

import (
	"github.com/SpecterTeam/SpecterGO/utils"
	"os"
	"time"
	"fmt"
	"math"
	"errors"
	"github.com/SpecterTeam/SpecterGO/command"
)

const (
	pluginPath = "plugins/"
	worldsPath = "worlds/"
	playersPath = "players/"
	tickDelay = time.Second / 20
)

type Server struct {
	running bool
	path string
}

var (
	ticks = make(map[int]float64)
	lastTick = 0
	TPS = 20
	ServerConfig utils.Config
)


func NewServer(path string) Server {
	s := Server{}
	ServerConfig = utils.NewConfig(utils.GetServerPath() + "server.yml", utils.TypeYaml, map[string]interface{}{
		"motd": Name + " Server",
		"server-port": 19132,
		"max-players": 100,
		"gamemode": 0,

	})
	InitDirectories(path)
	return s
}

func InitDirectories(path string) {
	os.Mkdir(path + pluginPath, 0777)
	os.Mkdir(path + worldsPath, 0777)
	os.Mkdir(path + playersPath, 0777)
}

func (s *Server) Start() {
	if s.Running() == true {
		err := errors.New("server is already running")
		logger.Error(err.Error())
	} else {
		s.SetRunning(true)
		command.InitCommands()
		cr := command.CommandReader{}
		cr.ReadConsole()
		s.InitTicker()
	}
}

func (s *Server) Shutdown() {
	s.SetRunning(false)
}

func (s *Server) Path() string {
	return s.path
}

func (s *Server) SetPath(path string) {
	s.path = path
}

func (s *Server) Running() bool {
	return s.running
}

func (s *Server) SetRunning(running bool) {
	s.running = running
}

func (s *Server) InitTicker() {
	for i := 1; i <= 20; i++ {
		ticks[i] = 20.0
	}
	s.Tick()
}

func (s *Server) Tick() {
	for range time.NewTicker(tickDelay).C {
		if s.Running() == false {
			break;
		}

		t := time.Now().Nanosecond()

		s.DoTitleTick()

		if lastTick == 20 {
			var all float64
			for _, tick := range ticks {
				all += tick
			}
			TPS = int(math.Round(all / 20))
			lastTick = 0
		}
		lastTick++
		ticks[lastTick] = float64(time.Now().Nanosecond() - t)
	}
}

func (s *Server) DoTitleTick() {
	fmt.Print("\x1b]0;" + "TPS: " + utils.IntToString(TPS) + "%\x07")
}