package repositories

import "demo/entity"

type UserRepository interface {
	Save(user *entity.User) (*entity.User, error)
	FindAll() ([]entity.User, error)
	Find(userId string) (*entity.User, error)
	Delete(userId string) error
}
