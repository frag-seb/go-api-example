package repositories

import (
	"demo/connection"
	"demo/entity"
)

type sqliteRepo struct{}

//NewUserSQLiteRepository
func NewUserSQLiteRepository() UserRepository {
	return &sqliteRepo{}
}

func (*sqliteRepo) Save(user *entity.User) (*entity.User, error) {
	connection.DB.Create(&user)

	return user, nil
}

func (*sqliteRepo) FindAll() ([]entity.User, error) {
	var users []entity.User
	if err := connection.DB.Find(&users).Error; err != nil {
		return users, nil
	}

	return users, nil
}
