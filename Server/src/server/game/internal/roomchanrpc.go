package internal

import "fmt"

func (room *Room) registerChanRPC() {
	room.roomChanRPC.Register("EnterRoom", room.rpcEnterRoom)
	room.roomChanRPC.Register("StartBattle", room.rpcStartBattle)

}

func  (room *Room)rpcEnterRoom(args []interface{}) {
	fmt.Println("rpcStartBattle")
	//fighting now.
}

func  (room *Room)rpcStartBattle(args []interface{}) {
	fmt.Println("rpcStartBattle")
	//fighting now.
}