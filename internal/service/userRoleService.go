package service

import (
	"github.com/MuZaZaVr/notesService/internal/model"
	"github.com/MuZaZaVr/notesService/internal/repository"
	"github.com/pkg/errors"
)

type UserRoleService struct {
	repository.UserRole
}

func newUserRoleService(repo repository.UserRole) *UserRoleService {
	return &UserRoleService{repo}
}


func (u UserRoleService) FindAllUsers() ([]model.User, error) {
	users, err := u.UserRole.FindAllUsers()
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't find all users")
	}

	return users, nil
}
