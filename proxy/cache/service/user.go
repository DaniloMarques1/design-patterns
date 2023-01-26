package service

import (
	"errors"
	"time"
)

type User struct {
	Id   string
	Name string
}

type UserService struct {
	users []*User
}

func NewUserService() *UserService {
	us := &UserService{}
	us.users = make([]*User, 0)
	return us
}

func (s *UserService) Add(user *User) {
	s.users = append(s.users, user)
}

func (s *UserService) FindById(id string) (*User, error) {
	time.Sleep(3 * time.Second)
	for _, user := range s.users {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, errors.New("No user found")
}
