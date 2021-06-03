package handler

import (
	"encoding/json"
	"fmt"
	"github.com/MuZaZaVr/notesService/internal/model"
	"github.com/MuZaZaVr/notesService/internal/service"
	"github.com/MuZaZaVr/notesService/pkg/middleware"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

/* Login user */
type loginRequest struct {
	model.LoginUserRequest
}

func (req *loginRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.LoginUserRequest)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	return nil
}

func (req *loginRequest) Validate() error {
	if req.Login == "" {
		return fmt.Errorf("login can not be nil")
	}

	if req.Password == "" {
		return fmt.Errorf("password can not be nil")
	}

	return nil
}


func (u *userRouter) loginUser(w http.ResponseWriter, r *http.Request) {
	var req loginRequest

	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, http.StatusBadRequest, err)
	}

	fmt.Printf("Login | Request: %v \n", req)
}

/* Register user */
type registerRequest struct {
	model.RegisterUserRequest
}

func (req *registerRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.RegisterUserRequest)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	return nil
}

func (req *registerRequest) Validate() error {
	if req.Login == "" {
		return fmt.Errorf("login can not be nil")
	}

	if req.Password == "" {
		return fmt.Errorf("password can not be nil")
	}

	return nil
}


func (u *userRouter) registerUser(w http.ResponseWriter, r *http.Request) {
	var req registerRequest

	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, http.StatusBadRequest, err)
		return
	}

	exist, err := u.services.User.IsExist(req.Login)
	if err != nil {
		middleware.JSONError(w, http.StatusInternalServerError, err)
		return
	}
	if exist {
		middleware.JSONReturn(w, http.StatusFound, "This user already exists")
		return
	}

	id, err := u.services.User.Create(req.RegisterUserRequest)
	if err != nil {
		 middleware.JSONError(w, http.StatusInternalServerError, err)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, strconv.Itoa(id))
}
