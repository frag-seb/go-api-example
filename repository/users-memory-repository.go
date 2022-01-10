package repository

import (
	"demo/entity"
	"github.com/google/uuid"
	"log"
)

var (
	users []entity.User
)

func init() {
	users = []entity.User{entity.User{Id: uuid.NewString(), Firstname: "Firstname 1", Lastname: "Lastname 1"}}
}

type repo struct{}

//NewUserMemoryRepository
func NewUserMemoryRepository() UserRepository {
	return &repo{}
}

func (*repo) Save(user *entity.User) (*entity.User, error) {
	//ctx := context.Background()
	users = append(users, *user)

	return user, nil
}

func (*repo) FindAll() ([]entity.User, error) {
	log.Println("FindAll memory")
	return users, nil
}
