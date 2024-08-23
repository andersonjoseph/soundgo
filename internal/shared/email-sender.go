package shared

type EmailSender interface {
	SendPasswordCode(email string, code string) error
}
