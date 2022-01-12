package service

import (
	"demo/entity"
	"demo/repositories"
	"errors"
	"log"

	"github.com/google/uuid"
)

type UserService interface {
	Validate(user *entity.User) error
	Create(user *entity.User) (*entity.User, error)
	FindAll() ([]entity.User, error)
	Find(userId string) (*entity.User, error)
	Delete(userId string) error
}

type service struct{}

var (
	repo repositories.UserRepository = repositories.NewUserSQLiteRepository()
)

//NewUserService
func NewUserService() UserService {
	return &service{}
}

func (*service) Validate(user *entity.User) error {
	if user == nil {
		return errors.New("User is empty")
	}

	if user.Firstname == "" {
		return errors.New("User firstname is empty")
	}

	return nil
}

func (*service) Create(user *entity.User) (*entity.User, error) {
	uuidResult, err := uuid.NewRandom()

	if err != nil {
		log.Fatalln(err)
	}

	user.ID = uuidResult.String()

	return repo.Save(user)
}

func (*service) FindAll() ([]entity.User, error) {
	return repo.FindAll()
}

func (*service) Find(userId string) (*entity.User, error) {
	return repo.Find(userId)
}

func (*service) Delete(userId string) error {
	return repo.Delete(userId)
}
