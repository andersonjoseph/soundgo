package main

import (
	"net/http"

	"github.com/andersonjoseph/soundgo/internal/reqcontext"
)

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
	ua := r.Header.Get("user-agent")
	ip := readUserIP(r)

	return ua + ip
}

func clientFingerprintMiddleware(h http.Handler) (http.Handler, error) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = reqcontext.Handler{}.SetClientFingerprint(ctx, getRequestFingerprint(r))

		h.ServeHTTP(w, r.WithContext(ctx))
	}), nil
}
