package repositories

import "demo/entity"

type UserRepository interface {
	Save(user *entity.User) (*entity.User, error)
	FindAll() ([]entity.User, error)
	Find(user_id string) (*entity.User, error)
	Delete(user_id string) error
}
