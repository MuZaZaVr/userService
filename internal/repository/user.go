package repository

import (
	"database/sql"
	"github.com/MuZaZaVr/notesService/internal/model"
	"github.com/MuZaZaVr/notesService/internal/model/dto"
	"strings"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}


func (u UserRepo) Create(user model.User) (int, error) {
	var id int

	rows, err := u.db.Query("INSERT INTO users (login, password, roleID) VALUES ($1, $2, $3) RETURNING id",
		user.Login, user.Password, dto.USER)
	if err != nil {
		return 0, err
	}

	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}

	return id, rows.Err()
}

func (u UserRepo) FindByLogin(login string) (*model.User, error) {
	var user model.User

	rows, err := u.db.Query("SELECT * FROM users WHERE login = $1", login)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Login, &user.Password, &user.RoleID)
		if err != nil {
			return nil, err
		}
	}

	return &user, rows.Err()
}

func (u UserRepo) FindByCredentials(user model.User) (*model.User, error) {
	var newUser model.User

	rows, err := u.db.Query("SELECT * FROM users WHERE login = $1 AND password = $2", user.Login, user.Password)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err := rows.Scan(&newUser.ID, &newUser.Login, &newUser.Password, &newUser.RoleID)
		if err != nil {
			return nil, err
		}
	}

	return &newUser, rows.Err()
}

func (u UserRepo) IsExist(login string) (bool, error) {
	userFromDB, err := u.FindByLogin(login)
	if err != nil && strings.Contains(err.Error(), "sql: Rows are closed") {
		return false, err
	} else if userFromDB != nil && userFromDB.ID != 0 {
		return true, nil
	}

	return false, nil
}