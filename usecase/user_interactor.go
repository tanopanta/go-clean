package usecase

import "github.com/tanopanta/go-clean/domain"

type UserInteractor struct {
	UserRepositry UserRepositry
}

func (interactor *UserInteractor) Add(u domain.User) (err error) {
	_, err = interactor.UserRepositry.Store(u)
	return
}

func (interactor *UserInteractor) Users() (users domain.Users, err error) {
	users, err = interactor.UserRepositry.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier int) (user domain.User, err error) {
	user, err = interactor.UserRepositry.FindById(identifier)
	return
}
