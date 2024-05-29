package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/andersonjoseph/soundgo/internal/shared"
)

func decodeBody[T any](body io.Reader) (T, error) {
	var res T
	err := json.NewDecoder(body).Decode(&res)

	if err != nil {
		return res, shared.ErrValidation
	}

	return res, nil
}

func sendResponse(dest io.Writer, body any) error {
	if err := json.NewEncoder(dest).Encode(body); err != nil {
		return shared.ErrValidation
	}

	return nil
}

func sendError(w http.ResponseWriter, err string, code int) {
	res := struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	}{
		Error:   http.StatusText(code),
		Message: err,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)
}
