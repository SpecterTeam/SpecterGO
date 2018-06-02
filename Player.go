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

const (
	Survival = iota
	Creative
	Adventure
	Spectator
	View = Spectator
)

type Gamemode int

type Player struct {
	protocol int
	ip string
	port int
	loggedIn bool
	spawned bool
	userName string
	displayName string
	xuid string
	playedBefore bool
	gameMode Gamemode
}

func NewPlayer() Player {
	p := Player{}
	p.SetProtocol(-1)
	p.SetLoggedIn(false)
	p.SetSpawned(false)
	p.SetUserName("")
	p.SetDisplayName("")
	p.SetXuid("")
	p.SetPlayedBefore(false)
	p.SetGameMode(Survival)

	return p
}

func (p *Player) GameMode() Gamemode {
	return p.gameMode
}

func (p *Player) SetGameMode(gameMode Gamemode) {
	p.gameMode = gameMode
}

func (p *Player) PlayedBefore() bool {
	return p.playedBefore
}

func (p *Player) SetPlayedBefore(playedBefore bool) {
	p.playedBefore = playedBefore
}

func (p *Player) IsAuthenticated() bool {
	return p.Xuid() != "";
}

func (p *Player) Xuid() string {
	return p.xuid
}

func (p *Player) SetXuid(xuid string) {
	p.xuid = xuid
}

func (p *Player) DisplayName() string {
	return p.displayName
}

func (p *Player) SetDisplayName(displayName string) {
	p.displayName = displayName
}

func (p *Player) UserName() string {
	return p.userName
}

func (p *Player) SetUserName(userName string) {
	p.userName = userName
}

func (p *Player) Spawned() bool {
	return p.spawned
}

func (p *Player) SetSpawned(spawned bool) {
	p.spawned = spawned
}

func (p *Player) LoggedIn() bool {
	return p.loggedIn
}

func (p *Player) SetLoggedIn(loggedIn bool) {
	p.loggedIn = loggedIn
}

func (p *Player) Port() int {
	return p.port
}

func (p *Player) SetPort(port int) {
	p.port = port
}

func (p *Player) Ip() string {
	return p.ip
}

func (p *Player) SetIp(ip string) {
	p.ip = ip
}

func (p *Player) Protocol() int {
	return p.protocol
}

func (p *Player) SetProtocol(protocol int) {
	p.protocol = protocol
}
