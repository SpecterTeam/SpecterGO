package SpecterGO

import (
	raknetServer "github.com/SpecterTeam/go-raknet/server"
	"github.com/satori/go.uuid"
	"github.com/beito123/binary"
	"github.com/SpecterTeam/go-raknet/identifier"
	"github.com/SpecterTeam/go-raknet"
)

type Players map[int]Player

type RakNetInterface struct {
	players Players
	server  *raknetServer.Server
}

func (r *RakNetInterface) Server() *raknetServer.Server {
	return r.server
}

func (r *RakNetInterface) SetServer(server *raknetServer.Server) {
	r.server = server
}

func (r *RakNetInterface) Players() Players {
	return r.players
}

func (r *RakNetInterface) SetPlayers(players Players) {
	r.players = players
}

func NewRakNetInterface() RakNetInterface {

	uid,_ := uuid.NewV4()

	id := identifier.Minecraft{
		Connection:        raknet.ConnectionGoRaknet,
		ServerName:        GetServer().Motd(),
		ServerProtocol:    raknet.NetworkProtocol,
		VersionTag:        "1.0.0",
		OnlinePlayersCount: 0,
		MaxPlayersCount:    GetServer().MaxPlayers(),
		GUID:              binary.ReadLong(uid.Bytes()[0:8]),
		WorldName:         "world",
		Gamemode:          "0",
		Legacy:            false,
	}
	ser := &raknetServer.Server{
		Logger:              GetLogger(),
		MaxConnections:      10,
		MTU:                 1472,
		Identifier:          id,
		UUID:                uid,
		BroadcastingEnabled: true,
	}
	ser.Start("0.0.0.0", GetServer().Port())
	return RakNetInterface{players: make(Players), server: ser}
}

func (r *RakNetInterface) Minecraft() *identifier.Minecraft {
	return &r.Server().Identifier
}

func (r *RakNetInterface) Shutdown() {
	r.Server().Shutdown()
}