package internal

import (
	"fmt"
	"server/mysql"
)

type PlayerBaseInfo struct {
	PlayerID uint `gorm:"primary_key"`
	Name     string `gorm:"not null"`
}

func (playerInfo *PlayerBaseInfo) initValue(playerID uint) error {
	mysql := mysql.MysqlDB()

	err := mysql.Where("PlayerID = ?", playerID). Limit(1).Find(&playerInfo).Error
	if nil != err {
		fmt.Println(err)
		return fmt.Errorf("get   PlayerBaseInfo id error: %v", err)
	}

	return nil
}


func (playerInfo *PlayerBaseInfo) saveValue() error {
	mysql := mysql.MysqlDB()
	err :=  mysql.Save(&playerInfo).Error

	if nil != err {
		fmt.Println(err)
		playerInfo = nil
		return fmt.Errorf("get   PlayerBaseInfo id error: %v", err)
	}

	return nil
}

func  CreatePlayerBaseInfo(playerID uint)  error{
	playerInfo := new (PlayerBaseInfo)

	mysql := mysql.MysqlDB()
	err :=  mysql.Save(&playerInfo).Error

	if nil != err {
		fmt.Println(err)
		return fmt.Errorf("get   PlayerBaseInfo id error: %v", err)
	}

	return nil
}





