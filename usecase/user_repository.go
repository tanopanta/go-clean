package usecase

import "github.com/tanopanta/go-clean/domain"

type UserRepositry interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
}
