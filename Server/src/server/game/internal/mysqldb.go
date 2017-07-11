package internal
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"fmt"
	"sync"
	"container/heap"
)


var (
	  	mysqlDB *gorm.DB
		mute = new (sync.Mutex)
)

func openDB()  {
	mute.Lock()
	defer mute.Unlock()
	mysqlDB, err := gorm.Open("mysql", "mike:123456@tcp(localhost:3306)/gorm?parseTime=true")
	if err != nil {
		panic("connect db error")
	}
}

func  Lock()  {
	mute.Lock()
}

func Unlock()  {
	mute.Unlock()
}

func  GetMysqlDBAndLock()  {
	Lock()
	return MysqlDB()
}

func  MysqlDB() *gorm.DB {
	return mysqlDB
}



type PlayerBaseInfo struct {
	ID    uint `gorm:"primary_key"`
	Name string `gorm:"not null;unique"`
}

func Select()  {
	mute.Lock()
	defer mute.Unlock()
	var playerBaseInfo PlayerBaseInfo
	err := mysqlDB.Where("Name = ?", "Mike"). Limit(1).Find(&playerBaseInfo).Error
	if nil != err {
		fmt.Println(err)
	}
}



func  GetOpenDB() (*gorm.DB) {
	return  mysqlDB
}
