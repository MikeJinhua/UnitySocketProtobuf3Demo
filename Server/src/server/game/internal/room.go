package internal

import (
	"github.com/name5566/leaf/module"
	"github.com/name5566/leaf/chanrpc"
	"server/base"
)

type Room struct{
	skeleton *module.Skeleton
	chanRPC   *chanrpc.Server
}

func (room *Room) OnInit() {
	room.skeleton = base.NewSkeleton()
	room.chanRPC = room.skeleton.ChanRPCServer
	room.registerChanRPC()
}