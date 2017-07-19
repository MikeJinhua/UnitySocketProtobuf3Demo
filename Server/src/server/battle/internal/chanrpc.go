package internal

import "fmt"

func init() {
	skeleton.RegisterChanRPC("BattleTestRpc", BattleTestRpc)
}


func BattleTestRpc(args []interface{}) {
	fmt.Println("BattleTestRpc")
	//
}

