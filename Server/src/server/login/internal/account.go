package internal

import (
	"fmt"
	"server/mysql"
)

type Account struct {
	PlayerID  uint `gorm:"primary_key"`
	AccountID string `gorm:"not null;unique"`
	Password string
}

func  getAccountByAccountID(accountID string) *Account{

	var account Account
	db := mysql.MysqlDB()
	err := db.Where("Account = ?", accountID). Limit(1).Find(&account).Error
	if nil != err {
		fmt.Println(err)
		return nil
	}
	fmt.Println("password:",account.Password)
	return &account
}


func creatAccountByAccountIDAndPassword(accountID string, password string) *Account{
	db := mysql.MysqlDB()
	var account = Account{AccountID:accountID,Password:password}
	err := db.Create(&account).Error
	if nil != err {
		return  nil
	}


	return  &account
}