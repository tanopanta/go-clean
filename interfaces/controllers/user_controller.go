package controllers

import (
	"strconv"

	"github.com/tanopanta/go-clean/domain"
	"github.com/tanopanta/go-clean/interfaces/database"
	"github.com/tanopanta/go-clean/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SQLHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepositry: &database.UserRepositry{
				SQLHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	u := domain.User{}
	c.Bind(&u)
	err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, u)
}

func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
}

func (controller *UserController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
}
