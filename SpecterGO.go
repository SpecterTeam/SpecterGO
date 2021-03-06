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
)

const (
	Name = "SpecterGO"
	Version = "0.1-ALPHA"
)

var (
	logger = utils.NewLogger()
	server Server
)

func GetLogger() *utils.Logger {
	return &logger
}

func SetServer(s Server) {
	server = s
}

func GetServer() *Server {
	return &server
}

func Load() {
	GetLogger().Info(Name + " version: " + Version)
	GetLogger().Info("Starting the server...")
	s := NewServer(utils.GetServerPath())
	SetServer(s)
	GetServer().Start()
}