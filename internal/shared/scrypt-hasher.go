package shared

import "github.com/elithrar/simple-scrypt"

type ScryptHasher struct{}

func (h ScryptHasher) Hash(p string) (string, error) {
	hash, err := scrypt.GenerateFromPassword([]byte(p), scrypt.DefaultParams)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (h ScryptHasher) Compare(hash, pass string) bool {
	return scrypt.CompareHashAndPassword([]byte(hash), []byte(pass)) == nil
}
