package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type request interface {
	Build(r *http.Request) error
	Validate() error
}

func ParseRequest(r *http.Request, req request) error {
	err := req.Build(r); if err != nil {
		return err
	}
	return req.Validate()
}

func JSONReturn(w http.ResponseWriter, statusCode int, jsonObject interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(jsonObject)
	if err != nil {
		fmt.Printf("could not encode json: %v \n", err.Error())
	}
}

func JSONError(w http.ResponseWriter, statusCode int, err error) {
	JSONReturn(w, statusCode, err)
}


func Empty(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
}