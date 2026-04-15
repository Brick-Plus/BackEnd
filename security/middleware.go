package security

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

type contextKey string

const UserIDKey contextKey = "userID"
const IsAdminKey contextKey = "isAdmin"

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"Error": err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), IsAdminKey, claims.IsAdmin)
		ctx = context.WithValue(ctx, UserIDKey, claims.UserId) // Ajouter aussi l'UserId pour les admins
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"Error": err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserId)
		ctx = context.WithValue(ctx, IsAdminKey, claims.IsAdmin) // Injecter aussi le rôle admin
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
