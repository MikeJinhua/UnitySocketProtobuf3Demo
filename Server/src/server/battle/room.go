package battle

import (
	"github.com/name5566/leaf/module"
	"github.com/name5566/leaf/chanrpc"
)

type Room struct{
	skeleton *module.Skeleton
	ChanRPC   *chanrpc.Server
}
