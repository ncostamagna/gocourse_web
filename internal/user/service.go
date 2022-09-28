package user

import "log"

type Service interface {
	Create(firstName, lastName, email, phone string) error
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s service) Create(firstName, lastName, email, phone string) error {
	log.Println("Create user service")
	return nil
}
