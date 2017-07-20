package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"server/conf"
	"server/game"
	"server/gate"
	"server/login"
	"server/gamedata"
	"fmt"
	"server/mysql"
)

func main() {
	mysql.OpenDB()
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	gamedata.LoadTables()
	testData := gamedata.GetDataByID(2)
	fmt.Println(testData.Name)

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)

}

func  InitDBTable()  {

}