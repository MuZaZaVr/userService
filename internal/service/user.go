package service

import (
	"github.com/MuZaZaVr/notesService/internal/model"
	"github.com/MuZaZaVr/notesService/internal/repository"
	"github.com/MuZaZaVr/notesService/pkg/auth"
	"github.com/pkg/errors"
	"strconv"
)

type UserService struct {
	repository.User
	auth.TokenManager
}

func newUserService(userRepo repository.User, tokenManager auth.TokenManager) *UserService {
	return &UserService{userRepo, tokenManager}
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

func (u UserService) FindByCredentials(req model.LoginUserRequest) (string, error) {
	user := model.User{
		Login: req.Login,
		Password: req.Password,
	}
	newUser, err := u.User.FindByCredentials(user)
	if err != nil {
		return "", errors.Wrap(err, "Can't find users with credentials")
	}

	if newUser.ID != 0 {
		newJWT, err := u.NewJWT(strconv.Itoa(newUser.ID))
		if err != nil {
			return "", errors.Wrap(err, "can't create a token")
		}
		return newJWT, nil
	}

	return "", nil
}

func (u UserService) IsExist(login string) (bool, error) {
	exist, err := u.User.IsExist(login)
	if err != nil {
		return false, errors.Wrap(err, "can not check user existence")
	}

	return exist, nil
}
