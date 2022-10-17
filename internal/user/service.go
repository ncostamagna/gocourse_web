package user

import "log"

type Service interface {
	Create(firstName, lastName, email, phone string) error
}

type service struct {
	log  *log.Logger
	repo Repository
}

func NewService(log *log.Logger, repo Repository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}

func (s service) Create(firstName, lastName, email, phone string) error {
	s.log.Println("Create user service")
	user := User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}

	if err := s.repo.Create(&user); err != nil {
		return err
	}

	return nil
}
