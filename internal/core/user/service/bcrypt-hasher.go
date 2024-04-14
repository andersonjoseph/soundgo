package service

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct{}

func (h BcryptHasher) Hash(p string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)

	return string(hashedPwd), err
}

func (h BcryptHasher) Compare(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
