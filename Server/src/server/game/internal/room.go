package internal

import (
	"github.com/name5566/leaf/module"
	"github.com/name5566/leaf/chanrpc"
	"server/base"
	"github.com/liangdas/mqant/gate"
	"fmt"
	"server/conf"
)
var(
	roomPlayers = make(map[uint]gate.Agent)
)

type Room struct{

	roomChanRPC   *chanrpc.Server
	roomID uint
}

func   newRoom(roomID uint) *Room{
	fmt.Println("newRoom",roomID)
	room := Room{roomChanRPC:chanrpc.NewServer(conf.ChanRPCLen)}

	room.registerChanRPC()
	room.roomID = roomID
	return  &room
}

func JoinRoom(agent gate.Agent)  {
	roomPlayers[uint(1)] = agent
}