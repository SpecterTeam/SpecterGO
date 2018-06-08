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

import (
	"github.com/SpecterTeam/go-raknet"
	"github.com/SpecterTeam/go-raknet/identifier"
)

type ConnectedPing struct {
	BasePacket
	Time int64
}

func (pk ConnectedPing) ID() byte {
	return IDConnectedPing
}

func (pk *ConnectedPing) Encode() error {
	err := pk.BasePacket.Encode(pk)
	if err != nil {
		return err
	}

	err = pk.PutLong(pk.Time)
	if err != nil {
		return err
	}

	return nil
}

func (pk *ConnectedPing) Decode() error {
	err := pk.BasePacket.Decode(pk)
	if err != nil {
		return err
	}

	pk.Time, err = pk.Long()
	if err != nil {
		return err
	}

	return nil
}

func (pk *ConnectedPing) New() raknet.Packet {
	return new(ConnectedPing)
}

type ConnectedPong struct {
	BasePacket
	Time int64
}

func (pk ConnectedPong) ID() byte {
	return IDConnectedPong
}

func (pk *ConnectedPong) Encode() error {
	err := pk.BasePacket.Encode(pk)
	if err != nil {
		return err
	}

	err = pk.PutLong(pk.Time)
	if err != nil {
		return err
	}

	return nil
}

func (pk *ConnectedPong) Decode() error {
	err := pk.BasePacket.Decode(pk)
	if err != nil {
		return err
	}

	pk.Time, err = pk.Long()
	if err != nil {
		return err
	}

	return nil
}

func (pk *ConnectedPong) New() raknet.Packet {
	return new(ConnectedPong)
}

type UnconnectedPing struct {
	BasePacket
	Timestamp       int64
	Magic           bool
	PingID          int64
	ConnectionMagic []byte
	Connection      *raknet.ConnectionType
}

func (pk UnconnectedPing) ID() byte {
	return IDUnconnectedPing
}

func (pk *UnconnectedPing) Encode() error {
	err := pk.BasePacket.Encode(pk)
	if err != nil {
		return err
	}

	err = pk.PutLong(pk.Timestamp)
	if err != nil {
		return err
	}

	err = pk.PutMagic()
	if err != nil {
		return err
	}

	err = pk.PutLong(pk.PingID)
	if err != nil {
		return err
	}

	err = pk.Put(raknet.ConnctionTypeMagic)
	if err != nil {
		return err
	}

	err = pk.PutConnectionType(pk.Connection)
	if err != nil {
		return err
	}

	return nil
}

func (pk *UnconnectedPing) Decode() error {
	err := pk.BasePacket.Decode(pk)
	if err != nil {
		return err
	}

	pk.Timestamp, err = pk.Long()
	if err != nil {
		return err
	}

	pk.Magic = pk.CheckMagic()

	pk.PingID, err = pk.Long()
	if err != nil {
		return err
	}

	if pk.Len() < len(raknet.ConnctionTypeMagic) {
		pk.Connection = &raknet.ConnectionVanilla

		return nil
	}

	pk.ConnectionMagic = pk.Get(len(raknet.ConnctionTypeMagic))

	pk.Connection, err = pk.ConnectionType()
	if err != nil {
		return err
	}

	return nil
}

func (pk *UnconnectedPing) New() raknet.Packet {
	return new(UnconnectedPing)
}

type UnconnectedPingOpenConnections struct {
	UnconnectedPing
}

func (pk UnconnectedPingOpenConnections) ID() byte {
	return IDUnconnectedPingOpenConnections
}

func (pk *UnconnectedPingOpenConnections) New() raknet.Packet {
	return new(UnconnectedPingOpenConnections)
}

type UnconnectedPong struct {
	BasePacket
	Timestamp       int64
	PongID          int64
	Magic           bool
	Identifier      identifier.Identifier
	ConnectionMagic []byte
	Connection      *raknet.ConnectionType
}

func (pk UnconnectedPong) ID() byte {
	return IDUnconnectedPong
}

func (pk *UnconnectedPong) Encode() error {
	err := pk.BasePacket.Encode(pk)
	if err != nil {
		return err
	}

	err = pk.PutLong(pk.Timestamp)
	if err != nil {
		return err
	}

	err = pk.PutLong(pk.PongID)
	if err != nil {
		return err
	}

	err = pk.PutMagic()
	if err != nil {
		return err
	}

	err = pk.PutString(pk.Identifier.Build())
	if err != nil {
		return err
	}

	err = pk.Put(raknet.ConnctionTypeMagic)
	if err != nil {
		return err
	}

	err = pk.PutConnectionType(pk.Connection)
	if err != nil {
		return err
	}

	return nil
}

func (pk *UnconnectedPong) Decode() error {
	err := pk.BasePacket.Decode(pk)
	if err != nil {
		return err
	}

	pk.Timestamp, err = pk.Long()
	if err != nil {
		return err
	}

	pk.PongID, err = pk.Long()
	if err != nil {
		return err
	}

	pk.Magic = pk.CheckMagic()

	id, err := pk.String()
	if err != nil {
		return err
	}

	if pk.Len() >= len(raknet.ConnctionTypeMagic) {
		pk.ConnectionMagic = pk.Get(len(raknet.ConnctionTypeMagic))

		pk.Connection, err = pk.ConnectionType()
		if err != nil {
			return err
		}
	} else {
		pk.Connection = &raknet.ConnectionVanilla
	}

	pk.Identifier = identifier.Base{
		Identifier: id,
		Connection: pk.Connection,
	}

	return nil
}

func (pk *UnconnectedPong) New() raknet.Packet {
	return new(UnconnectedPong)
}
