package service

import (
	"demo/entity"
	"demo/repository"
	"errors"

	"github.com/google/uuid"
)

type UserService interface {
	Validate(user *entity.User) error
	Create(user *entity.User) (*entity.User, error)
	FindAll() ([]entity.User, error)
}

type service struct{}

var (
	repo repository.UserRepository = repository.NewUserMemoryRepository()
)

//NewUserService
func NewUserService() UserService {
	return &service{}
}

func (*service) Validate(user *entity.User) error {
	if user == nil {
		err := errors.New("User is empty")
		return err
	}

	if user.Firstname == "" {
		err := errors.New("User firstname is empty")
		return err
	}

	return nil
}

func (*service) Create(user *entity.User) (*entity.User, error) {
	user.Id = uuid.NewString()

	return repo.Save(user)
}

func (*service) FindAll() ([]entity.User, error) {
	return repo.FindAll()
}
