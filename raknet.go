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
package raknet

import (
	"time"
)

const (

	// Version is version of go-raknet library
	Version = "1.0.1"

	// NetworkProtocol is a version of Raknet protocol
	NetworkProtocol = 8

	// MaxMTU is the maximum size of MTU
	MaxMTU = 1492

	// MinMTU is the minimum size of MTU
	MinMTU = 400

	// MaxChannel is the maximum size of order channel
	MaxChannels = 32

	// DefaultChannel is default channel
	DefaultChannel = 0

	// MaxSplitCount is the maximum size that can split
	MaxSplitCount = 128

	// MaxSplitsPerQueue is the maximum size of Queue
	MaxSplitsPerQueue = 4
)

// Magic is Raknet offline message data id
// using offline connection in Raknet
var Magic = []byte{0x00, 0xff, 0xff, 0x00, 0xfe, 0xfe, 0xfe, 0xfe, 0xfd, 0xfd, 0xfd, 0xfd, 0x12, 0x34, 0x56, 0x78}

// MaxPacketsPerSecond is the maximum size that can send per second
var MaxPacketsPerSecond = 500

var (
	// SendInterval
	SendInterval                           = 50 * time.Millisecond
	RecoverySendInterval                   = SendInterval
	PingSendInterval                       = 2500 * time.Millisecond
	DetectionSendInterval                  = PingSendInterval * 2
	SessionTimeout                         = DetectionSendInterval * 5
	MaxPacketsPerSecondBlock               = 1000 * 300 * time.Millisecond
)