package shared

import (
	"fmt"
	"os"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type JWTHandler struct{}

func (h JWTHandler) GenerateToken(owner string) (string, error) {
	tok, err := jwt.NewBuilder().
		IssuedAt(time.Now()).
		Subject(owner).
		Build()

	if err != nil {
		return "", fmt.Errorf("error building token: %w", err)
	}

	key, err := getJWTKey()
	if err != nil {
		return "", fmt.Errorf("error getting JWT key: %w", err)
	}

	signedToken, err := jwt.Sign(tok, jwt.WithKey(jwa.HS512, []byte(key)))
	if err != nil {
		return "", fmt.Errorf("error signing JWT: %w", err)
	}

	return string(signedToken), nil
}

func (h JWTHandler) ParseToken(token string) (jwt.Token, error) {
	key, err := getJWTKey()
	if err != nil {
		return nil, fmt.Errorf("error getting JWT key: %w", err)
	}

	verifiedToken, err := jwt.Parse([]byte(token), jwt.WithKey(jwa.HS512, []byte(key)), jwt.WithValidate(true))

	if err != nil {
		return nil, fmt.Errorf("error parsing JWT key: %w", err)
	}

	return verifiedToken, nil
}

func getJWTKey() (string, error) {
	key, ok := os.LookupEnv("JWT_KEY")
	if !ok {
		return "", fmt.Errorf("JWT_KEY is not present in environment")
	}

	return key, nil
}
