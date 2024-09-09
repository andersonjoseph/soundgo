package main

import (
	"context"
	"net/http"
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
		fingerprint := getRequestFingerprint(r)

		ctx := r.Context()
		ctx = context.WithValue(ctx, "client-fingerprint", fingerprint)

		h.ServeHTTP(w, r.WithContext(ctx))
	}), nil
}
