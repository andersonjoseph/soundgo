package shared

type SecretHasher interface {
	Hash(p string) (string, error)
	Compare(hashedPassword, password string) bool
}
