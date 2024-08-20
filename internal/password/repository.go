package password

type Repository interface {
	GetCode(email string) (string, error)
	Save(email, password string) error
}
