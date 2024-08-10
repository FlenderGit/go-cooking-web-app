package auth

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userKey contextKey = "user"

type UserJWT struct {
	ID    int
	Login string
}

type CustomClaims struct {
	jwt.RegisteredClaims
}

func IsAuthenticated(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		testLogConsole()

		// Get the token from the request
		token, err := getTokenFromRequest(r)

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Convert the claims to a CustomClaims struct
		claims, err := ParseToken(token.Raw)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Set the context
		ctx := context.WithValue(r.Context(), userKey, claims)
		r = r.WithContext(ctx)

		log.Printf("User %s is authenticated", claims.Login)

		next.ServeHTTP(w, r)
	}
}

func ParseToken(token string) (UserJWT, error) {
	claims := &CustomClaims{}
	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !t.Valid {
		return UserJWT{}, errors.New("invalid token")
	}

	id, err := strconv.Atoi(claims.ID)
	if err != nil {
		return UserJWT{}, errors.New("invalid token")
	}

	return UserJWT{
		ID:    id,
		Login: claims.Subject,
	}, nil
}

func GenerateToken(id int, username string) (string, error) {

	if id <= 0 {
		return "", errors.New("invalid id")
	}

	if username == "" {
		return "", errors.New("invalid username")
	}

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Second)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "source de confiance",
		Subject:   username,
		ID:        strconv.Itoa(id),
		Audience:  []string{"somebody_else"},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}

func getTokenFromRequest(r *http.Request) (*jwt.Token, error) {

	auth := r.Header.Get("Authorization")
	if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
		return nil, errors.New("no authorization header")
	}

	encoded_token := strings.TrimPrefix(auth, "Bearer ")
	//encoded_token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzb3VyY2UgZGUgY29uZmlhbmNlIiwic3ViIjoidGVzdCIsImF1ZCI6WyJzb21lYm9keV9lbHNlIl0sImV4cCI6MTcyMzI5NjU0NiwibmJmIjoxNzIzMjk2NTIyLCJpYXQiOjE3MjMyOTY1MjIsImp0aSI6IjEifQ.OLci5bhqshHuUZG635eS17lDJ31KJy8C1eYHwTosZq8"
	token, err := jwt.ParseWithClaims(encoded_token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

func GetUserFromRequest(r *http.Request) (UserJWT, error) {

	user, ok := r.Context().Value(userKey).(UserJWT)
	if !ok {
		return UserJWT{}, errors.New("no user in context")
	}

	return user, nil
}

func testLogConsole() {
	token, err := GenerateToken(1, "test")
	if err != nil {
		log.Printf("Error")
	}
	log.Printf("Token generated : " + token)

}
