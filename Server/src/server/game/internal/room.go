package internal

import (
	"fmt"
)

type Room struct{
	roomID uint
	roomPlayers map[uint]gate.Agent
}

func   newRoom(roomID uint) *Room{
	fmt.Println("newRoom",roomID)
	room := Room{roomID:roomID}

	return  &room
}

