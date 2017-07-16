package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"server/msg"
	"github.com/name5566/leaf/log"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("UserLogin", rpcUserLogin)
	skeleton.RegisterChanRPC("CreatePlayer", rpcCreatePlayer)
}

func rpcNewAgent(args []interface{}) {
	fmt.Println("rpcNewAgent")
	a := args[0].(gate.Agent)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	fmt.Println("rpcCloseAgent")
	a := args[0].(gate.Agent)
	_ = a
}

func rpcUserLogin(args []interface{}) {
	agent := args[0].(gate.Agent)
	playerID := args[1].(uint)
	// network closed
	if agent.UserData() == nil {
		return
	}

	oldUser := playerID2Player[playerID]
	if oldUser != nil {
		m := &msg.LoginFaild{Code: msg.LoginFaild_LoginRepeat}
		agent.WriteMsg(m)
		oldUser.WriteMsg(m)
		agent.Close()
		oldUser.Close()
		log.Debug("acc %v login repeated", playerID)
		return
	}
	log.Debug("acc %v login", playerID)

	// login
	newPlayer := new(Player)
	newPlayer.Agent = agent
	newPlayer.LinearContext = skeleton.NewLinearContext()
	newPlayer.state = userLogin
	newPlayer.UserData().(*PlayerBaseInfo).PlayerID = playerID
	playerID2Player[playerID] = newPlayer
	newPlayer.login(playerID)
}

func rpcCreatePlayer(args []interface{}) {
	agent := args[0].(gate.Agent)
	playerID := args[1].(uint)
	err := CreatePlayer(playerID)
	if nil != err {
		m := &msg.LoginFaild{Code: msg.LoginFaild_InnerError}
		agent.WriteMsg(m)
	}

}
