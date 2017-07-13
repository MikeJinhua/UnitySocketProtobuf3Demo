package internal
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"fmt"
	"sync"
)


var (
	  	mysqlDB *gorm.DB
		mute = new (sync.Mutex)
)

func OpenDB()  {
	fmt.Println("mysqldb->open db")
	mysqlDB1, err := gorm.Open("mysql", "mike:123456@tcp(localhost:3306)/gorm?parseTime=true")
	mysqlDB = mysqlDB1
	if err != nil {
		panic("connect db error")
	}
}

func InitDBTables()  {
	mysqlDB.AutoMigrate(&PlayerBaseInfo{})
}

func  Lock()  {
	mute.Lock()
}

func Unlock()  {
	mute.Unlock()
}

func  GetMysqlDBAndLock()  *gorm.DB{
	Lock()
	return MysqlDB()
}

func  MysqlDB() *gorm.DB {
	return mysqlDB
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
