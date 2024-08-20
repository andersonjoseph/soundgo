package shared

type EmailSender interface {
	SendPasswordCode(email string) error
}
