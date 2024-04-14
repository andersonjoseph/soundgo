package server

import "net/http"

func (s *server) handleHealthCheck() http.HandlerFunc {
	type response struct {
		Status int `json:"status"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := sendResponse(w, response{
			Status: 200,
		})

		if err != nil {
			s.handleError(r.Context(), err, w)
			return
		}

		return
	}

}
