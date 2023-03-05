package user

import (
	"errors"
)

type Service interface {
	Save(user User) User
	FindAll() (users []User)
	FindById(Id int) (user User, err error)
	FindByEmail(Email string) (user User, err error)
	Delete(user User)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

// Save is function for create data user
func (s *service) Save(user User) User {
	user = s.repository.Create(user)
	return user
}

// FindAll is function for get all data user
func (s *service) FindAll() (users []User) {
	users = s.repository.FindAll()
	return
}

// FindById is function for get detail data user
func (s *service) FindById(Id int) (user User, err error) {
	user = s.repository.FindById(Id)
	if user.Id == 0 {
		return user, errors.New("user not found")
	}
	return user, nil
}

// FindByEmail is function for get detail data user
func (s *service) FindByEmail(Email string) (user User, err error) {
	user = s.repository.FindByEmail(Email)
	if user.Id == 0 {
		return user, errors.New("user not found")
	}
	return user, nil
}

// Delete is function for delete data user
func (s *service) Delete(user User) {
	s.repository.Delete(user)

	return
}
