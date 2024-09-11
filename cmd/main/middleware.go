package main

import (
	"net/http"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/reqcontext"
	"github.com/andersonjoseph/soundgo/internal/shared"
)

func clientFingerprintMiddleware(s *api.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route, ok := s.FindRoute(r.Method, r.URL.Path)
		if !ok || route.Name() != "GetAudioFile" {
			s.ServeHTTP(w, r)
			return
		}

		ctx := r.Context()
		ctx = reqcontext.ClientFingerprint.Set(ctx, getRequestFingerprint(r))

		s.ServeHTTP(w, r.WithContext(ctx))
	})
}

func requestIDMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id, err := shared.GenerateUUID()
		if err != nil {
			panic(err)
		}

		ctx = reqcontext.RequestID.Set(ctx, id)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func readUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func getRequestFingerprint(r *http.Request) string {
	return r.Header.Get("user-agent") + readUserIP(r)
}
