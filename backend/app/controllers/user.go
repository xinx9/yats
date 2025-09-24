package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	userGo "yats/app/models"
)

type UserHandler struct {
	storage userGo.UserStorage
}

func NewUserHandler(storage userGo.UserStorage) *UserHandler {
	return &UserHandler{storage: storage}
}

func (u UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(u.storage.List())
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (u UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idNum, _ := strconv.Atoi(id)
	user := u.storage.Get(idNum)
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (u UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	data := &userGo.User{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser := new(userGo.User)
	newUser.Name = data.Name
	newUser.ID = int(len(u.storage.List()) + 1)
	u.storage.Create(*newUser)

	render.Status(r, http.StatusCreated)
}

// func (u userData) Bind(r *http.Request) error {
// 	if u.User == nil {
// 		return errors.New("missing required User field")
// 	}

// 	return nil
// }

func (u UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var user userGo.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	idnum, _ := strconv.Atoi(id)
	updatedUser := u.storage.Update(idnum, user)
	if updatedUser == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(updatedUser)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (u UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idnum, _ := strconv.Atoi(id)
	deletedUser := u.storage.Delete(idnum)
	if deletedUser == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(deletedUser)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
