package internal

import (
	"github.com/name5566/leaf/module"
	"server/base"
	"server/mysql"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}


func (m *Module) OnInit() {
	m.Skeleton = skeleton
	InitGameTables()
	room := newRoom()
	///test call room rpc
	room.chanRPC.Go("StartBattle", 1, 1)
}

func InitGameTables() {
	db:=mysql.MysqlDB()
	db.AutoMigrate(&PlayerBaseInfo{})
}

func (m *Module) OnDestroy() {

}
