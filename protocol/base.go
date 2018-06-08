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
	"github.com/SpecterTeam/go-raknet/binary"
)

type BasePacket struct {
	binary.RaknetStream
}

func (base *BasePacket) Encode(pk raknet.Packet) error {
	err := base.PutByte(pk.ID())
	if err != nil {
		return err
	}

	return nil
}

func (base *BasePacket) Decode(pk raknet.Packet) error {
	base.Skip(1) // for id

	return nil
}

func (base *BasePacket) Bytes() []byte {
	return base.RaknetStream.AllBytes()
}
