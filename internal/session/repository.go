package session

type SaveInput struct {
	Token string
}

type Repository interface {
	Save(username string, i SaveInput) (Entity, error)
	Delete(ID string) error
}
