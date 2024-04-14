package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type jwtPasswordTokenGenerator struct {
	secretKey []byte
}

func NewJWTPasswordTokenGenerator(s []byte) jwtPasswordTokenGenerator {
	return jwtPasswordTokenGenerator{
		secretKey: s,
	}
}

func (g jwtPasswordTokenGenerator) Generate(userID int) (string, error) {
	token := jwt.New()

	if err := token.Set(jwt.IssuedAtKey, time.Now()); err != nil {
		return "", fmt.Errorf("error setting issued at claim: %w", err)
	}

	if err := token.Set(jwt.ExpirationKey, time.Now().Add(time.Hour)); err != nil {
		return "", fmt.Errorf("error setting expiration claim: %w", err)
	}

	if err := token.Set(jwt.SubjectKey, fmt.Sprint(userID)); err != nil {
		return "", fmt.Errorf("error setting subject claim: %w", err)
	}

	signedToken, err := jwt.Sign(
		token,
		jwt.WithKey(jwa.HS256, g.secretKey),
	)

	if err != nil {
		return "", fmt.Errorf("error signin token: %w", err)
	}

	return string(signedToken), nil
}

func (g jwtPasswordTokenGenerator) Decode(t string) (int, error) {
	token, err := jwt.Parse(
		[]byte(t),
		jwt.WithKey(jwa.HS256, g.secretKey),
		jwt.WithValidate(true),
		jwt.WithVerify(true),
	)

	if err != nil {
		return 0, err
	}

	userID, err := strconv.Atoi(token.Subject())

	if err != nil {
		return 0, err
	}

	return userID, nil
}
