package repositories

import (
	"demo/entity"
	"reflect"
)

var (
	users []entity.User
)

func init() {
	users = []entity.User{}
}

type repo struct{}

//NewUserMemoryRepository
func NewUserMemoryRepository() UserRepository {
	return &repo{}
}

func (*repo) Save(user *entity.User) (*entity.User, error) {
	users = append(users, *user)

	return user, nil
}

func (*repo) FindAll() ([]entity.User, error) {
	return users, nil
}

func (*repo) Find(userId string) (*entity.User, error) {
	var user entity.User

	for _, v := range users {
		if v.ID == userId {
			user = entity.User{
				ID:        v.ID,
				Firstname: v.Firstname,
				Lastname:  v.Lastname,
			}
		}
	}

	if user.ID == "" {
		return nil, nil
	}

	return &user, nil
}

func (*repo) Delete(userId string) error {

	for i, v := range users {
		if v.ID == userId {
			// Found it!
			users = append(users[:i], users[i+1:]...)
			break
		}
	}

	return nil
}

func sliceRemoveItem(slicep interface{}, i int) {
	v := reflect.ValueOf(slicep).Elem()
	v.Set(reflect.AppendSlice(v.Slice(0, i), v.Slice(i+1, v.Len())))
}
