package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/timer"
	"github.com/name5566/leaf/go"
	"server/msg"
	"github.com/name5566/leaf/log"
	"time"
	"github.com/name5566/leaf/util"
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

	skeleton.Go(func() {
		err := playerBaseInfo.initValue(playerID)
		if err != nil {
			log.Error("init acc %v data error: %v", playerID, err)
			playerBaseInfo = nil
			player.WriteMsg(&msg.LoginFaild{Code: msg.LoginFaild_InnerError})
			player.Close()
			return
		}
	}, func() {
		// network closed
		if player.state == userLogout {
			player.logout(playerID)
			return
		}

		// db error
		player.state = userGame
		if playerBaseInfo == nil {
			return
		}

		// ok
		player.playerBaseInfo = playerBaseInfo
		playerID2Player[playerID] = player
		//player.UserData().(*AgentInfo).userID = userData.UserID
		player.onLogin()
		player.autoSaveDB()
	})

}

func CreatePlayer(playerID uint) error {
	err := CreatePlayerBaseInfo(playerID)
	return err
}

func (player *Player) isOffline() bool {
	return player.state == userLogout
}

func (player *Player) logout(playerID uint) {

}

func (player *Player) autoSaveDB() {
	const duration = 5 * time.Minute
	// save
	player.saveDBTimer = skeleton.AfterFunc(duration, func() {
		data := util.DeepClone(player.playerBaseInfo)
		player.Go(func() {
			err := data.(*PlayerBaseInfo).saveValue()
			if err != nil {
				log.Error("save user %v data error: %v", player.playerBaseInfo.PlayerID, err)
			}

		}, func() {
			player.autoSaveDB()
		})
	})
}

func (player *Player) onLogin() {

}

func (player *Player) onLogout() {

}
