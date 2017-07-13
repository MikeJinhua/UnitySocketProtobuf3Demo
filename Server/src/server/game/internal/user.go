package internal


type Player struct {
	ID    uint
	//初始化之后自己写函数读取BaseInfo和CardsInfo吧，保存信息也一样。
	//本来想用ORM的级联发现各种不习惯。还是自己写代码控制下逻辑好了，也多不了几行代码。
	playerBaseInfo *PlayerBaseInfo

	//...其他信息
}

func (user *Player) login(accID string) {
	playerBaseInfo:= new(PlayerBaseInfo)
	user.playerBaseInfo = playerBaseInfo

}

type PlayerBaseInfo struct {
	ID    uint `gorm:"primary_key"`
	Name string `gorm:"not null;unique"`
}