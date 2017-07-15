package internal


type PlayerBaseInfo struct {
	PlayerID uint `gorm:"primary_key"`
	Name     string `gorm:"not null;unique"`
}

func (data *PlayerBaseInfo) initValue(playerID uint) error {


	return nil
}







