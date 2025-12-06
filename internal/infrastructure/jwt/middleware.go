package jwt

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const UserContextKey contextKey = "user"

type UserContext struct {
	UserID   int
	Username string
}

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Отсутствует токен авторизации", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Неверный формат токена", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		claims, err := ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Недействительный токен: "+err.Error(), http.StatusUnauthorized)
			return
		}

		userCtx := &UserContext{
			UserID:   claims.UserID,
			Username: claims.Username,
		}
		ctx := context.WithValue(r.Context(), UserContextKey, userCtx)

		next(w, r.WithContext(ctx))
	}
}

func GetUserFromContext(ctx context.Context) (*UserContext, bool) {
	user, ok := ctx.Value(UserContextKey).(*UserContext)
	return user, ok
}
