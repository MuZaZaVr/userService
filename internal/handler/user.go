package handler

import (
	"github.com/MuZaZaVr/notesService/internal/model"
	"github.com/MuZaZaVr/notesService/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

type userRouter struct {
	*mux.Router
	services *service.Service
}

func newUserRouter(services *service.Service) userRouter {
	router := mux.NewRouter().PathPrefix(userPath).Subrouter()
	handler := userRouter{
		Router:   router,
		services: services,
	}

	router.Path("/login").Methods(http.MethodPost).HandlerFunc(handler.loginUser)
	router.Path("/registration").Methods(http.MethodPost).HandlerFunc(handler.registerUser)

	return handler
}

type loginRequest struct {
	model.LoginUserRequest
}

func (u *userRouter) loginUser(w http.ResponseWriter, r *http.Request) {

}


type registerRequest struct {
	model.RegisterUserRequest
}

func (u *userRouter) registerUser(w http.ResponseWriter, r *http.Request) {

}
