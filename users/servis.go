package users

import (
	"fmt"
)

type Service interface {
	ListUsers() []UserT
	GetUser(id int) (UserT, error)
	DeleteUser(userId int) error
	UpsertUser(id int, user UserT) (UserT,error)
}

var Users = map[int]UserT{}

type service struct {}

func NewUsersService() Service {
	return &service{}
}

func (s *service ) ListUsers() []UserT {
	users := make([]UserT, 0, len(Users))

	for _, user := range Users {
		users = append(users, user)
	}

	return users
}

func (s *service) GetUser(id int) (UserT, error)  {
	if val, ok := Users[id]; ok {
		return val, nil
	}
	return UserT{}, fmt.Errorf("Пользователь не найден")
}

func (s *service) DeleteUser(id int) error {
	if _, ok := Users[id]; ok {
		delete(Users, id)
		return nil
	}
	return fmt.Errorf("Пользователь не найден")
}

func (s *service) UpsertUser(id int, user UserT) (UserT,error) {
	user.ID = id

	if val, ok := Users[id]; ok {
		val = user
		Users[id] = val

		return user, nil
	}

	Users[id] = user

	return user, nil
}