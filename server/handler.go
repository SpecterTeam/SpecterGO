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
package server

import (
	"net"

	"github.com/SpecterTeam/go-raknet"
)

// Handler handles packets, connections and more from Raknet server
type Handler interface {

	// StartServer is called on the server is started
	StartServer()

	// CloseServer is called on the server is closed
	CloseServer()

	// HandlePing is called on a ping packet is received
	HandlePing(addr net.Addr)

	// OpenPreConn is called before a new session is created
	OpenPreConn(addr net.Addr)

	// OpenConn is called on a new session is created
	OpenConn(uid int64, addr net.Addr)

	// ClosePreConn is called before a session is closed
	ClosePreConn(uid int64)

	// CloseConn is called on a session is closed
	CloseConn(uid int64)

	// HandleSendPacket handles a packet sent from the server to a client
	HandleSendPacket(addr net.Addr, pk raknet.Packet)

	// HandleRawPacket handles a raw packet no processed in Raknet server
	HandleRawPacket(addr net.Addr, pk raknet.Packet)

	// HandlePacket handles a message packet
	HandlePacket(uid int64, pk raknet.Packet)

	// HandleUnknownPacket handles a unknown packet
	HandleUnknownPacket(uid int64, pk raknet.Packet)

}
