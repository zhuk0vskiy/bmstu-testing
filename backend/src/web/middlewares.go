package web

import (
	"fmt"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

func ValidateUserRoleJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			ErrorResponse(w, fmt.Errorf("getting claims from JWT: %w", err).Error(), http.StatusBadRequest)
			return
		}

		role, ok := claims["role"]
		if !ok {
			ErrorResponse(w, fmt.Errorf("получение 'role' claim`а из JWT").Error(), http.StatusBadRequest)
			return
		}

		if role != "client" && role != "admin" {
			ErrorResponse(w, fmt.Errorf("вам нужно авторизоваться, прежде чем сделать это").Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ValidateAdminRoleJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			ErrorResponse(w, fmt.Errorf("getting claims from JWT: %w", err).Error(), http.StatusBadRequest)
			return
		}

		role, ok := claims["role"]
		if !ok {
			ErrorResponse(w, fmt.Errorf("получение 'role' claim`а из JWT").Error(), http.StatusBadRequest)
			return
		}

		if role != "admin" {
			ErrorResponse(w, fmt.Errorf("только администраторы могут делать это").Error(), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
