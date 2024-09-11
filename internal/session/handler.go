package session

import (
	"github.com/andersonjoseph/soundgo/internal/shared"
	"github.com/andersonjoseph/soundgo/internal/user"
)

type Entity struct {
	ID     string
	Token  string
	UserID string
}

type Handler struct {
	repository     Repository
	userRepository user.Repository
	hasher         shared.PasswordHasher
	jwtHandler     shared.JWTHandler
}

func NewHandler(
	repository Repository,
	userRepository user.Repository,
	hasher shared.PasswordHasher,
	jwtHandler shared.JWTHandler,
) Handler {
	return Handler{
		repository:     repository,
		userRepository: userRepository,
		hasher:         hasher,
		jwtHandler:     jwtHandler,
	}
}
