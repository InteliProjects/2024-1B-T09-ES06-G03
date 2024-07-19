// em /service/auth/jwt.go

package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
  "strconv"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/ceo/config"
	"github.com/dgrijalva/jwt-go"
)

var JWTSecret = config.Envs.JWTSecret

func JWTMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenString := extractToken(r)
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return []byte(JWTSecret), nil
        })

        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            userID := extractUserID(claims["userID"])
            if userID == -1 {
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }
            ctx := context.WithValue(r.Context(), "userID", userID)
            next.ServeHTTP(w, r.WithContext(ctx))
        } else {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
        }
    })
}

func extractUserID(v interface{}) int {
    switch id := v.(type) {
    case float64:
        return int(id)
    case string:
        if intID, err := strconv.Atoi(id); err == nil {
            return intID
        }
    }
    return -1
}


func extractToken(r *http.Request) string {
    bearerToken := r.Header.Get("Authorization")
    // Normalmente, os tokens JWT são enviados como 'Bearer {token}'
    return strings.TrimPrefix(bearerToken, "Bearer ")
}

