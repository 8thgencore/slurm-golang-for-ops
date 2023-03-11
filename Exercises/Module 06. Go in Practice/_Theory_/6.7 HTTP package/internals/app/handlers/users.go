package handlers

import (
	"6_7/example/internals/app/models"
	"6_7/example/internals/app/processors"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

type UsersHandler struct {
	processor *processors.UsersProcessor
}

func NewUsersHandler(processor *processors.UsersProcessor) *UsersHandler {
	handler := new(UsersHandler)
	handler.processor = processor
	return handler
}

func (handler *UsersHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		WrapError(w, err)
		return
	}

	err = handler.processor.CreateUser(newUser)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{} {
		"result" : "OK",
		"data" : "",
	}

	WrapOK(w, m)
}

func (handler *UsersHandler) List(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	list, err := handler.processor.ListUsers(strings.Trim(vars.Get("name"), "\""))

	if err != nil {
		WrapError(w, err)
	}

	var m = map[string]interface{} {
		"result" : "OK",
		"data" : list,
	}

	WrapOK(w, m)
}

func (handler *UsersHandler) Find(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //переменные, обьявленные в ресурсах попадают в Vars и могут быть адресованы
	if vars["id"] == "" {
		WrapError(w, errors.New("missing id"))
		return
	}

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		WrapError(w, err)
		return
	}

	user, err := handler.processor.FindUser(id)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{} {
		"result" : "OK",
		"data" : user,
	}

	WrapOK(w, m)
}
