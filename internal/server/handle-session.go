package server

import (
	"fmt"
	"net/http"

	"github.com/andersonjoseph/soundgo/internal/core/user/service"
	"github.com/andersonjoseph/soundgo/internal/core/user/use-case/login"
	loginRepository "github.com/andersonjoseph/soundgo/internal/core/user/use-case/login/repository"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func (s *server) handleCreateSession() http.HandlerFunc {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type response struct {
		Token string `json:"token"`
	}

	service := login.New(
		loginRepository.NewPGRepository(s.db),
		service.BcryptHasher{},
	)

	return func(w http.ResponseWriter, r *http.Request) {
		body, err := decodeBody[request](r.Body)

		if err != nil {
			s.handleError(r.Context(), err, w)
			return
		}

		session, err := service.LoginUser(r.Context(), login.Dto{
			Username: body.Username,
			Password: body.Password,
		})

		if err != nil {
			s.handleError(r.Context(), err, w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		token := jwt.New()
		if err := token.Set(jwt.SubjectKey, fmt.Sprint(session.ID)); err != nil {
			s.handleError(r.Context(), fmt.Errorf("error creating session token: %w", err), w)
			return
		}

		signedToken, err := jwt.Sign(token, jwt.WithKey(jwa.HS256, s.sessionTokenSecret))

		if err != nil {
			s.handleError(r.Context(), fmt.Errorf("error signin session token: %w", err), w)
			return
		}

		err = sendResponse(w, response{
			Token: string(signedToken),
		})

		if err != nil {
			s.handleError(r.Context(), err, w)
			return
		}
	}
}
