package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/andersonjoseph/soundgo/internal/api"
	"github.com/andersonjoseph/soundgo/internal/reqcontext"
	"github.com/andersonjoseph/soundgo/internal/shared"
)

type RouteFinder interface {
	FindRoute(method string, path string) (api.Route, bool)
}

func clientFingerprintMiddleware(h http.Handler, rf RouteFinder) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route, ok := rf.FindRoute(r.Method, r.URL.Path)
		if !ok || route.OperationID() != "getAudioFile" {
			h.ServeHTTP(w, r)
			return
		}

		ctx := r.Context()
		ctx = reqcontext.ClientFingerprint.Set(ctx, getRequestFingerprint(r))

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

type ResponseWriterWithStatus struct {
	http.ResponseWriter
	Status int
}

func (w *ResponseWriterWithStatus) WriteHeader(code int) {
	w.Status = code
	w.ResponseWriter.WriteHeader(code)
}

func LogRequestMiddlware(h http.Handler, rf RouteFinder, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route, ok := rf.FindRoute(r.Method, r.URL.Path)
		if !ok {
			h.ServeHTTP(w, r)
			return
		}

		id, err := shared.GenerateUUID()
		if err != nil {
			panic(err)
		}
		ctx := reqcontext.RequestID.Set(r.Context(), id)

		logger.Info(
			"request received",
			"ID", id,
			"operation", route.OperationID(),
			"path", route.PathPattern(),
			"method", r.Method,
		)

		start := time.Now()
		writerWithStatus := &ResponseWriterWithStatus{ResponseWriter: w}

		h.ServeHTTP(writerWithStatus, r.WithContext(ctx))

		logger.Info(
			"response delivered",
			"ID", id,
			"operation", route.OperationID(),
			"path", route.PathPattern(),
			"method", r.Method,
			"duration", time.Since(start).Milliseconds(),
			"status_code", writerWithStatus.Status,
		)
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
