package security

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type contextKey string

const UserIDKey contextKey = "userID"
const IsAdminKey contextKey = "isAdmin"

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"code": 401, "error_code": "no_authorization", "msg": "Missing token"}`, http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, `{"code":401,"error_code":"no_authorization","msg":"Invalid token format"}`, http.StatusUnauthorized)
			return
		}

		claims, err := VerifyAdminToken(tokenString)
		if err != nil {
			fmt.Println(err)
			http.Error(w, `{"code":401,"error_code":"no_authorization","msg":"Invalid token, not an admin"}`, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), IsAdminKey, claims.IsAdmin)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"code": 401, "error_code": "no_authorization", "msg": "Missing token"}`, http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, `{"code":401,"error_code":"no_authorization","msg":"Invalid token format"}`, http.StatusUnauthorized)
			return
		}

		claims, err := VerifyToken(tokenString)
		if err != nil {
			http.Error(w, `{"code":401,"error_code":"no_authorization","msg":"Invalid token"}`, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}