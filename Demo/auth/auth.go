package auth

import (
	"Demo/entity"
	"Demo/errors"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func IsLoggedIn (next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		tokenString := r.Header.Get("Authorization")
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
			// Validating expected algorithm
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte("secret"), nil
		})
		if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errors.MessageError{Message:"Unauthorized, please login."})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func IsAdmin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var user entity.User
		tokenString := r.Header.Get("Authorization")
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
			// Validating expected algorithm
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte("secret"), nil
		})
		if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
			if user.IsAdmin == false {
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(errors.MessageError{Message:"Forbidden, access denied."})
				return
			}
			return
		}
		next.ServeHTTP(w, r)
	})
}