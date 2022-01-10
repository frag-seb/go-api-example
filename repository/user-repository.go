package repository

import "demo/entity"

type UserRepository interface {
	Save(user *entity.User) (*entity.User, error)
	FindAll() ([]entity.User, error)
	//Find
	//Delete
}
