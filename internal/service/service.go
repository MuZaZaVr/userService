package service

import (
	"github.com/MuZaZaVr/notesService/internal/model"
	"github.com/MuZaZaVr/notesService/internal/repository"
)

type Service struct {
	User User
	Role UserRole
}

type Depends struct {
	Repos *repository.Repositories
}

func NewServices(deps Depends) *Service {
	return &Service{
		User: newUserService(deps.Repos.UserRepo),
		Role: newUserRoleService(deps.Repos.UserRole),
	}
}

type User interface {
	Create(req model.RegisterUserRequest) (int, error)
	FindByLogin(login string) (*model.User, error)
	FindByCredentials(req model.LoginUserRequest) (*model.User, error)
}

type UserRole interface {
	FindAllUsers() ([]model.User, error)
}
