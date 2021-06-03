package service

import (
	"github.com/MuZaZaVr/notesService/internal/model"
	"github.com/MuZaZaVr/notesService/internal/repository"
	"github.com/MuZaZaVr/notesService/pkg/auth"
)

type Service struct {
	User User
	Role UserRole
}

type Depends struct {
	Repos *repository.Repositories
	TokenManager auth.TokenManager
}

func NewServices(deps Depends) *Service {
	return &Service{
		User: newUserService(deps.Repos.UserRepo, deps.TokenManager),
		Role: newUserRoleService(deps.Repos.UserRole),
	}
}

type User interface {
	Create(req model.RegisterUserRequest) (int, error)
	FindByLogin(login string) (*model.User, error)
	FindByCredentials(req model.LoginUserRequest) (string, error)

	IsExist(login string) (bool, error)
}

type UserRole interface {
	FindAllUsers() ([]model.User, error)
}
