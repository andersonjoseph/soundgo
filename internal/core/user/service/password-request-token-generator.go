package service

type RequestPasswordTokenHandler interface {
	Generate(userID int) (string, error)
	Decode(token string) (int, error)
}
