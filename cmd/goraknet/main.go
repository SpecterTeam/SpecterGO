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
package main

import (
	"bufio"
	"context"
	"net"
	"os"

	"github.com/beito123/binary"
	"github.com/satori/go.uuid"
	"github.com/SpecterTeam/go-raknet"
	"github.com/SpecterTeam/go-raknet/identifier"
	"github.com/SpecterTeam/go-raknet/server"
	"github.com/SpecterTeam/SpecterGO/utils"
)

func main() {
	logger := utils.NewLogger()

	uid, _ := uuid.NewV4()

	id := identifier.Minecraft{
		Connection:        raknet.ConnectionGoRaknet,
		ServerName:        "SpecterGO-Server",
		ServerProtocol:    raknet.NetworkProtocol,
		VersionTag:        "1.0.0",
		OnlinePlayerCount: 0,
		MaxPlayerCount:    10,
		GUID:              binary.ReadLong(uid.Bytes()[0:8]),
		WorldName:         "world",
		Gamemode:          "0",
		Legacy:            false,
	}

	ser := &server.Server{
		Logger:              &logger,
		MaxConnections:      10,
		MTU:                 1472,
		Identifier:          id,
		UUID:                uid,
		BroadcastingEnabled: true,
	}

	ctx, cancel := context.WithCancel(context.Background())

	logger.Info("Starting the server...")

	addr := &net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: 19180,
	}

	go ser.ListenAndServe(ctx, addr)

	logger.Info("Enter to stop the server")

	bufio.NewScanner(os.Stdin).Scan() // wait input anything

	cancel()

	logger.Info("Stopping the server...")
}
