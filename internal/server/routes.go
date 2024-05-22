package server

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	"time"

	"github.com/sqids/sqids-go"
)

func (s server) middlewarePrepareHandler(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := rand.Int()

		ctx := context.WithValue(r.Context(), "req_id", reqID)
		ctx = context.WithValue(ctx, "req_path", r.Method+" "+r.URL.Path)

		logger := s.logger.With(
			slog.Group(
				"req",
				"id", ctx.Value("req_id"),
				"path", ctx.Value("req_path"),
			),
		)

		logger.LogAttrs(
			ctx,
			slog.LevelInfo,
			"incoming request",
		)

		now := time.Now()

		h(w, r.WithContext(ctx))

		logger.LogAttrs(
			ctx,
			slog.LevelInfo,
			"response sent",
			slog.Int("duration", int(time.Since(now).Milliseconds())),
		)
	})
}

type sqidsIDEncoder struct {
	sqids sqids.Sqids
}

func newSqidsEncoder() (sqidsIDEncoder, error) {
	s, err := sqids.New()

	if err != nil {
		return sqidsIDEncoder{}, err
	}

	return sqidsIDEncoder{
		sqids: *s,
	}, err
}

func (e sqidsIDEncoder) Encode(id int) (string, error) {
	return e.sqids.Encode([]uint64{uint64(id)})
}

func (e sqidsIDEncoder) Decode(id string) int {
	return int(e.sqids.Decode(id)[0])
}

func (s *server) registerRoutes() {
	idEncoder, err := newSqidsEncoder()

	if err != nil {
		panic(fmt.Errorf("error creating id encoder: %w", err))
	}

	s.handler.HandleFunc("GET /api/v1/health", s.handleHealthCheck())

	s.handler.HandleFunc("POST /api/v1/sessions", s.middlewarePrepareHandler(s.handleCreateSession()))

	s.handler.HandleFunc("POST /api/v1/password-reset", s.middlewarePrepareHandler(s.handleRequestResetPassword()))
	s.handler.HandleFunc("PUT /api/v1/password-reset", s.middlewarePrepareHandler(s.handleResetPassword()))

	s.handler.HandleFunc("POST /api/v1/users",
		s.middlewarePrepareHandler(
			s.handleRegisterUser(
				idEncoder,
			),
		),
	)
}
