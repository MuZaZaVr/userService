package repository

import (
	"database/sql"
	"github.com/MuZaZaVr/notesService/internal/model"
	"github.com/MuZaZaVr/notesService/internal/model/dto"
)

type UserRoleRepo struct {
	db *sql.DB
}

func newUserRoleRepo(db *sql.DB) *UserRoleRepo {
	return &UserRoleRepo{db: db}
}


func (u *UserRoleRepo) FindAllUsers() ([]model.User, error) {
	var user model.User
	var users []model.User

	rows, err := u.db.Query("SELECT u.id, u.login, u.password, u.roleID FROM users u INNER JOIN user_role ur ON u.roleID = ur.id WHERE u.roleID = $1", dto.USER)
	if err != nil {
		return nil, err
	}

	for i := 0; rows.Next(); i++ {
		err := rows.Scan(&user.ID, &user.Login, &user.Password, &user.RoleID)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
