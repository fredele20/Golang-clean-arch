package controller

import (
	"Demo/entity"
	"Demo/errors"
	"Demo/service/users"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

type UserController interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	service users.Service
}

func NewUserController(s users.Service) UserController{
	return &controller{
		service: s,
	}
}

func (c *controller) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.MessageError{Message: "Error unmarshalling the request"})
		return
	}

	err = c.service.Validate(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.MessageError{Message:err.Error()})
		return
	}

	_, err = c.service.UserAlreadyExist(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.MessageError{Message:err.Error()})
		return
	}
	//err := users.EncryptPassword(user.Password)
	_, err = c.service.EncryptPassword(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.MessageError{Message:"could not hash password"})
		return
	}

	result, err := c.service.Register(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.MessageError{Message:err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func (c *controller) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	//var password string

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.MessageError{Message:"error unmarshalling data"})
	}

	_, err := c.service.ValidatePassword(user.Password, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.MessageError{Message:err.Error()})
		return
	}

	result, err := c.service.Login(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.MessageError{Message:err.Error()})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"email": user.Email,
		"isAdmin": user.IsAdmin,
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.MessageError{Message:"error while generating token"})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{} {
		"token": tokenString,
		"user": result,
	})
	return
}
