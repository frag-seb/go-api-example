package repositories

import (
	"demo/connection"
	"demo/entity"
	"errors"
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

func (*sqliteRepo) Find(userId string) (*entity.User, error) {
	var user entity.User

	if err := connection.DB.First(&user, "id = ?", userId).Error; err != nil {
		return &user, nil
	}

	return &user, nil
}

func (*sqliteRepo) Delete(userId string) error {
	if err := connection.DB.Delete(&entity.User{}, "id = ?", userId).Error; err != nil {
		return errors.New("User could still be deleted")
	}

	return nil
}
