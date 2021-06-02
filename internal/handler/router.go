package handler

import (
	"github.com/MuZaZaVr/notesService/internal/service"
	"github.com/gorilla/mux"
)

const userPath = "/user"

type API struct {
	*mux.Router
}

func NewHandler(services *service.Service) *API {
	api := API{
		mux.NewRouter(),
	}

	api.PathPrefix(userPath).Handler(newUserRouter(services))

	return &api
}