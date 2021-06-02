package repository

import (
	"database/sql"
	"github.com/MuZaZaVr/notesService/internal/model"
)

type Repositories struct {
	UserRepo User
	UserRole UserRole
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		UserRepo: NewUserRepo(db),
		UserRole: newUserRoleRepo(db),
	}
}

type User interface {
	Create(user model.User) (int, error)
	FindByLogin(login string) (*model.User, error)
	FindByCredentials(login, password string) (*model.User, error)
}

type UserRole interface {
	FindAllUsers() ([]model.User, error)
}