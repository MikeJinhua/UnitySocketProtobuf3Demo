package internal

import "fmt"

func (room *Room) registerChanRPC() {
	room.chanRPC.Register("StartBattle", room.rpcStartBattle)

}


func  (room *Room)rpcStartBattle(args []interface{}) {
	fmt.Println("rpcStartBattle")
	//fighting now.
}