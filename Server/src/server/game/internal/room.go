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

func   newRoom() *Room{
	room := Room{skeleton:base.NewSkeleton(),chanRPC:skeleton.ChanRPCServer}
	room.registerChanRPC()
	return  &room
}