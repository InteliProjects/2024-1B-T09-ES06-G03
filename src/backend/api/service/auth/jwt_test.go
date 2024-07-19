package auth

import (
	"strconv"
	"testing"
	"time"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/config"
	"github.com/golang-jwt/jwt"
)

func TestCreateJWT(t *testing.T) {
	// Define a userID for the test
	testUserID := 1

	// Create the JWT with the global configuration and testUserID
	tokenString, err := CreateJWT([]byte(config.Envs.JWTSecret), testUserID)
	if err != nil {
		t.Fatalf("Failed to create JWT: %v", err)
	}

	// Verify the token is not empty
	if tokenString == "" {
		t.Fatalf("Token string is empty")
	}

	// Parse the token to check its contents
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the expected signing method was used
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(config.Envs.JWTSecret), nil
	})
	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}

	// Verify the claims are correct
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if strconv.Itoa(testUserID) != claims["userID"].(string) {
			t.Errorf("Expected userID %d, got %s", testUserID, claims["userID"])
		}

		// Verify the token expires at the expected time
		if int64(claims["expiresAt"].(float64)) <= time.Now().Unix() {
			t.Errorf("Token expires too soon")
		}
	} else {
		t.Errorf("Failed to read claims")
	}
}
