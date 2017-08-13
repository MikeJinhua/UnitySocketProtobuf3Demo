package internal


var(
	rooms = make(map[uint]*Room)
	lastRoomID = uint(0)
)

func CreateRooms(roomID uint) *Room{
	room := newRoom(roomID)
	rooms[roomID] = room
	return  room
}


func GetRoom(roomID uint) *Room{
	room, ok := rooms[roomID]
	if(!ok){
		   room = CreateRooms(roomID)
	}
	return  room
}

