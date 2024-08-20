package shared

type PasswordHasher interface {
	Hash(p string) (string, error)
	Compare(hashed, plain string) bool
}
