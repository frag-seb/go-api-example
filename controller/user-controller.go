package controller

import (
	"demo/entity"
	"demo/service"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type UserController interface {
	GetUsers(response http.ResponseWriter, request *http.Request)
	GetUser(response http.ResponseWriter, request *http.Request)
	PostUsers(response http.ResponseWriter, request *http.Request)
	DeleteUser(response http.ResponseWriter, request *http.Request)
}
type controller struct{}

var (
	userService service.UserService = service.NewUserService()
)

//NewUserController
func NewUserController() UserController {
	return &controller{}
}

func (*controller) GetUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	users, err := userService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errors.New("Error user not found"))
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(users)
}

func (*controller) GetUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)
	user_id := vars["id"]

	if user_id == "" {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errors.New("Error user not found"))
		return
	}

	user, err := userService.Find(user_id)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errors.New("Error user not found"))
		return
	}

	if user == nil || user.ID == "" {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errors.New("Error user not found"))
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(user)
}

func (*controller) PostUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user entity.User
	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.New("Error unmarshalling the reques"))
		return
	}

	err1 := userService.Validate(&user)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.New("Error saving the user"))
		return
	}

	userService.Create(&user)

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(user)
}

func (*controller) DeleteUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)
	user_id := vars["id"]

	fmt.Println(user_id)

	if user_id == "" {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errors.New("Error user not found"))
		return
	}

	err := userService.Delete(user_id)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errors.New("Error user not found"))
		return
	}

	response.WriteHeader(http.StatusAccepted)
	json.NewEncoder(response).Encode(nil)
}
