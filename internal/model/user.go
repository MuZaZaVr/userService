package model

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	RoleID   int    `json:"role_id"`
}
