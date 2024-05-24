package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"user_name";gorm:"size:256"`
	UserPW   string `json:"user_pw";gorm:"size:256"`
	Type     string `json:"type";gorm:"size:5"`
}

func GetUserByID(id uint) (user User, err error) {
	Db.AutoMigrate(&user)
	err = Db.Debug().First(&user, id).Error
	return user, err
}
func CreateUser(user User) error {
	Db.AutoMigrate(&user)
	err := Db.Model(&user).Create(user).Error
	return err
}
