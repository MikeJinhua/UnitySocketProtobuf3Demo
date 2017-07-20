package battle

import "fmt"

func (room *Room) RegisterChanRPC() {
	room.chanRPC.Register("StartBattle", room.rpcStartBattle)

}


func  (room *Room)rpcStartBattle(args []interface{}) {
	fmt.Println("rpcStartBattle")
	//fighting now.
}