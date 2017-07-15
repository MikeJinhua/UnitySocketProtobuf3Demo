package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/timer"
	"github.com/name5566/leaf/go"
)

var (
	playerID2Player = make(map[uint]*Player)
)

const (
	userLogin  = iota
	userLogout
	userGame
)

type Player struct {
	gate.Agent
	*g.LinearContext
	state       int
	saveDBTimer *timer.Timer
	//初始化之后自己写函数读取BaseInfo和CardsInfo吧，保存信息也一样。
	//本来想用ORM的级联发现各种不习惯。还是自己写代码控制下逻辑好了，也多不了几行代码。
	playerBaseInfo *PlayerBaseInfo

	//...其他信息
}

func (player *Player) login(playerID uint) {
	playerBaseInfo := new(PlayerBaseInfo)
	player.playerBaseInfo = playerBaseInfo

}

func (player *Player) logout(playerID uint) {

}

func (player *Player) autoSaveDB() {

}
