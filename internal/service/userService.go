package service

import (
	"github.com/MuZaZaVr/notesService/internal/model"
	"github.com/MuZaZaVr/notesService/internal/repository"
	"github.com/pkg/errors"
)

type UserService struct {
	repository.User
}

func newUserService(userRepo repository.User) *UserService {
	return &UserService{userRepo}
}


func (u UserService) Create(req model.RegisterUserRequest) (int, error) {
	user := model.User {
		Login: req.Login,
		Password: req.Password,
	}
	id, err := u.User.Create(user)
	if err != nil {
		return 0, errors.Wrap(err, "Can't create user")
	}

	return id, nil
}

func (u UserService) FindByLogin (login string) (*model.User, error) {
	user, err := u.User.FindByLogin(login)
	if err != nil {
		return nil, errors.Wrap(err, "Can't find user with login")
	}
	return user, nil
}

func (u UserService) FindByCredentials(req model.LoginUserRequest) (*model.User, error) {
	user := model.User{
		Login: req.Login,
		Password: req.Password,
	}
	newUser, err := u.User.FindByCredentials(user)
	if err != nil {
		return nil, errors.Wrap(err, "Can't find users with credentials")
	}

	return newUser, nil
}

