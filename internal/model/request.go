package model

type (
	RegisterUserRequest struct {
		Login    string
		Password string
	}

	LoginUserRequest struct {
		Login    string
		Password string
	}
)
