package main
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

func main1() {

	Test()
}


type Player struct {
	ID    uint
	//初始化之后自己写函数读取BaseInfo和CardsInfo吧，保存信息也一样。
	//本来想用ORM的级联发现各种不习惯。还是自己写代码控制下逻辑好了，也多不了几行代码。
	playerBaseInfo PlayerBaseInfo
	cards_map map[uint]PlayersCards // key is cardID , value is PlayersCards info

	//...其他信息
}

type PlayerBaseInfo struct {
	ID    uint `gorm:"primary_key"`
	Name string `gorm:"not null;unique"`
}

type PlayersCards struct {
	ID uint
	PlayerID uint `gorm:"unique_index:idx_name_code"`
	CardID uint `gorm:"unique_index:idx_name_code"`
	Amount uint`gorm:"default:'0'"`
	Level uint `gorm:"default:'1'"`
}

type Card struct {
	ID uint `gorm:"primary_key"`
	Name string `gorm:"not null"` // e.a. Knights of the Round Table
}



func Test() {
	db, err := gorm.Open("mysql", "mike:xxxxx@tcp(localhost:3306)/gorm?parseTime=true")
	defer db.Close()
	if err != nil {
		panic("connect db error")
	}
	defer db.Close()
	db.DropTableIfExists(&Player{}, &Card{}, &PlayersCards{},&PlayerBaseInfo{})
	////
	db.AutoMigrate(&PlayerBaseInfo{})
	db.AutoMigrate(&Card{})
	db.AutoMigrate(&PlayersCards{})

	playersCard1 := PlayersCards{PlayerID:1,CardID:1,Amount:20}
	err =  db.Save(&playersCard1).Error
	if nil != err {
		fmt.Println("create 1 error:",err)
	}

	playersCard2 := PlayersCards{PlayerID:1,CardID:2,Amount:65}
	err =  db.Save(&playersCard2).Error
	if nil != err {
		fmt.Println("create 2 error:",err)
	}

	playersCard1.Amount = 0
	err =  db.Save(playersCard1).Error
	if nil != err {
		fmt.Println("Update 1 error:",err)
	}

	playersCard3 := PlayersCards{PlayerID:2,CardID:3,Amount:615}
	err =  db.Save(&playersCard3).Error
	if nil != err {
		fmt.Println("create 3 error:",err)
	}

	player1 := PlayerBaseInfo{ID:2,Name:"Mike11221"}
	err =  db.Save(&player1).Error
	if nil != err {
		fmt.Println("create 3 error:",err)
	}


	//err2 := db.Create(&Player{
	//	Name: "Mike",
	//}).Error
	//if nil != err2 {
	//	fmt.Println("already exist:",err2)
	//}


	//query player Mike
	var player Player
	err = db.Where("Name = ?", "Mike"). Limit(1).Find(&player).Error
	if nil != err {
		fmt.Println(err)
	}
}

