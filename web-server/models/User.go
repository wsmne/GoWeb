package models

type User struct {
	ID       int
	UserName string
	UserPW   string
	Type     int
}

func GetUserByID(id string) (user User) {
	Db.Debug().First(&user, id)
	return user
}
func CreateUser(user User) error {
	err := Db.Model(&user).Create(user).Error
	return err
}
