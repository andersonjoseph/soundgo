package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/andersonjoseph/soundgo/internal/core/user/service"
	"github.com/andersonjoseph/soundgo/internal/shared"

	"github.com/andersonjoseph/soundgo/internal/core/user/use-case/register"

	registerRepository "github.com/andersonjoseph/soundgo/internal/core/user/use-case/register/repository"
)

func (s *server) handleRegisterUser(idEncoder shared.IDEncoder) http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type response struct {
		ID       string `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	}

	service := register.New(
		registerRepository.NewPGRepository(s.db),
		service.BcryptHasher{},
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
			"creating user",
			slog.Group(
				"body",
				slog.String("username", body.Username),
				slog.String("email", body.Email),
			),
		)

		u, err := service.RegisterUser(r.Context(), register.Dto{
			Email:    body.Email,
			Username: body.Username,
			Password: body.Password,
		})

		if err != nil {
			s.handleError(r.Context(), err, w)
			return
		}

		logger.LogAttrs(
			r.Context(),
			slog.LevelInfo,
			"user created",
			slog.Int("http_code", http.StatusCreated),
			slog.Group(
				"user",
				slog.Int("id", u.ID),
			),
		)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		encodedId, err := idEncoder.Encode(u.ID)

		if err != nil {
			s.handleError(r.Context(), fmt.Errorf("error encoding user id: %w", err), w)
			return
		}

		err = sendResponse(w, response{
			ID:       encodedId,
			Email:    u.Email.String(),
			Username: u.Username.String(),
		})

		if err != nil {
			s.handleError(r.Context(), err, w)
			return
		}

		return
	}
}
