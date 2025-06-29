package middleware

import (
	"fmt"
	"net/http"
)

type AuthMiddleware struct{}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("AuthMiddleware start")
		fmt.Println(r)
		fmt.Println(r.Header)
		next(w, r)
		fmt.Println("AuthMiddleware end")
	}
}
