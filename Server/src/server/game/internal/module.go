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
	InitRooms()

}

func InitGameTables() {
	db:=mysql.MysqlDB()
	db.AutoMigrate(&PlayerBaseInfo{})
}

func InitRooms(){
	for i := 0; i < 10; i++ {
		CreateRooms(uint(i))
	}
}

func (m *Module) OnDestroy() {

}
