package server

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/andersonjoseph/soundgo/internal/core/user/service"
	"github.com/andersonjoseph/soundgo/internal/core/user/use-case/request-password-reset"
	requestPasswordResetRepository "github.com/andersonjoseph/soundgo/internal/core/user/use-case/request-password-reset/repository"
	"github.com/andersonjoseph/soundgo/internal/core/user/use-case/reset-password"
	resetPasswordRepository "github.com/andersonjoseph/soundgo/internal/core/user/use-case/reset-password/repository"
	"github.com/andersonjoseph/soundgo/internal/shared"
)

func (s *server) handleRequestResetPassword() http.HandlerFunc {
	type request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	service := requestresetpassword.New(
		requestPasswordResetRepository.NewPGRepository(s.db),
		service.NewJWTPasswordTokenGenerator(s.sessionTokenSecret),
		service.NewFakePasswordTokenSender(*s.logger),
	)

	return func(w http.ResponseWriter, r *http.Request) {
		body, err := decodeBody[request](r.Body)

		if err != nil {
			s.handleError(r.Context(), err, w)
			return
		}

		logger := s.logger.With(
			slog.Group(
				"req",
				"id", r.Context().Value("req_id"),
				"path", r.Context().Value("req_path"),
			),
		)

		logger.LogAttrs(
			r.Context(),
			slog.LevelInfo,
			"rquesting password reset",
			slog.Group(
				"body",
				slog.String("username", body.Username),
				slog.String("email", body.Email),
			),
		)

		err = service.SendRequestPasswordToken(r.Context(), requestresetpassword.Dto{
			Username: body.Username,
			Email:    body.Email,
		})

		if err != nil && !errors.Is(err, shared.ErrNotFound) {
			s.handleError(r.Context(), err, w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}

func (s *server) handleResetPassword() http.HandlerFunc {
	type request struct {
		Token    string `json:"token"`
		Password string `json:"password"`
	}

	service := resetpassword.New(
		resetPasswordRepository.NewPGRepository(s.db),
		service.BcryptHasher{},
		service.NewJWTPasswordTokenGenerator(s.sessionTokenSecret),
	)

	return func(w http.ResponseWriter, r *http.Request) {
		body, err := decodeBody[request](r.Body)

		if err != nil {
			s.handleError(r.Context(), err, w)
			return
		}

		err = service.ResetPassword(r.Context(), resetpassword.Dto{
			Token:    body.Token,
			Password: body.Password,
		})

		if err != nil {
			s.handleError(r.Context(), err, w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}
}
