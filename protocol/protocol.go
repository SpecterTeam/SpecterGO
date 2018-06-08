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
package protocol

import "github.com/SpecterTeam/go-raknet"

type Protocol struct {
	packets []raknet.Packet
}

func (protocol *Protocol) RegisterPackets() {
	protocol.packets = make([]raknet.Packet, 0xff)

	protocol.packets[IDConnectedPing] = &ConnectedPing{}
	protocol.packets[IDUnconnectedPing] = &UnconnectedPing{}
	protocol.packets[IDUnconnectedPingOpenConnections] = &UnconnectedPingOpenConnections{}
	protocol.packets[IDOpenConnectionRequest1] = &OpenConnectionRequestOne{}
	protocol.packets[IDOpenConnectionReply1] = &OpenConnectionResponseOne{}
	protocol.packets[IDOpenConnectionRequest2] = &OpenConnectionRequestTwo{}
	protocol.packets[IDOpenConnectionReply2] = &OpenConnectionResponseTwo{}
	protocol.packets[IDConnectionRequest] = &ConnectionRequest{}
	protocol.packets[IDConnectionRequestAccepted] = &ConnectionRequestAccepted{}
	protocol.packets[IDAlreadyConnected] = &AlreadyConnected{}
	protocol.packets[IDNewIncomingConnection] = &NewIncomingConnection{}
	protocol.packets[IDNoFreeIncomingConnections] = &NoFreeIncomingConnections{}
	protocol.packets[IDDisconnectionNotification] = &DisconnectionNotification{}
	protocol.packets[IDConnectionBanned] = &ConnectionBanned{}
	protocol.packets[IDIncompatibleProtocolVersion] = &IncompatibleProtocol{}
	protocol.packets[IDUnconnectedPong] = &UnconnectedPong{}
	protocol.packets[IDCustom0] = NewCustomPacket(IDCustom0)
	protocol.packets[IDCustom1] = NewCustomPacket(IDCustom1)
	protocol.packets[IDCustom2] = NewCustomPacket(IDCustom2)
	protocol.packets[IDCustom3] = NewCustomPacket(IDCustom3)
	protocol.packets[IDCustom4] = NewCustomPacket(IDCustom4)
	protocol.packets[IDCustom5] = NewCustomPacket(IDCustom5)
	protocol.packets[IDCustom6] = NewCustomPacket(IDCustom6)
	protocol.packets[IDCustom7] = NewCustomPacket(IDCustom7)
	protocol.packets[IDCustom8] = NewCustomPacket(IDCustom8)
	protocol.packets[IDCustom9] = NewCustomPacket(IDCustom9)
	protocol.packets[IDCustomA] = NewCustomPacket(IDCustomA)
	protocol.packets[IDCustomB] = NewCustomPacket(IDCustomB)
	protocol.packets[IDCustomC] = NewCustomPacket(IDCustomC)
	protocol.packets[IDCustomD] = NewCustomPacket(IDCustomD)
	protocol.packets[IDCustomE] = NewCustomPacket(IDCustomE)
	protocol.packets[IDCustomF] = NewCustomPacket(IDCustomF)

}

func (protocol *Protocol) Packet(id byte) (pk raknet.Packet, ok bool) {
	pk = protocol.packets[id]
	if pk == nil {
		return nil, false
	}

	return pk, true
}

func (protocol *Protocol) Packets() []raknet.Packet {
	return protocol.packets
}
