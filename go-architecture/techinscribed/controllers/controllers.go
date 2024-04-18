package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"techinscribed-course/models"
)

type BaseHandler struct {
	userRepo models.UserRepository
}

type UserReqBody struct {
	ID int `json:"id"`
}

func NewBaseHandler(userRepo models.UserRepository) *BaseHandler {
	return &BaseHandler{
		userRepo: userRepo,
	}
}

func (h *BaseHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	if user, err := h.userRepo.FindById(1); err != nil {
		fmt.Println("Error:", user)
	}

	w.Write([]byte("Hello World!"))
}

func (h *BaseHandler) FindById(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	json.NewEncoder(w).Encode("method not supported")
	// 	return
	// }

	fmt.Println("id", r.PathValue("id"), r.URL.Path)
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id not valid"))
		return
	}

	// var body UserReqBody
	// err := json.NewDecoder(r.Body).Decode(&body)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// }

	user, err := h.userRepo.FindById(id)
	if err != nil {
		fmt.Println("Error:", user, err)
	}

	// w.Write([]byte(user.Name))
	json.NewEncoder(w).Encode(user)
}
