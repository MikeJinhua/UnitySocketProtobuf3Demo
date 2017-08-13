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
	skeleton *module.Skeleton
	chanRPC   *chanrpc.Server
	roomID uint
}

func   newRoom(roomID uint) *Room{
	fmt.Println("newRoom",roomID)
	room := Room{skeleton:base.NewSkeleton(),chanRPC:chanrpc.NewServer(conf.ChanRPCLen)}

	room.registerChanRPC()
	room.roomID = roomID
	room.skeleton.GoLen = 1111111111
	return  &room
}

func JoinRoom(agent gate.Agent)  {
	roomPlayers[uint(1)] = agent
}