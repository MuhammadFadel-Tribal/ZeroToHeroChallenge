package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func RenderJSON(ctx context.Context, w http.ResponseWriter, httpStatusCode int, payload interface{}) {
	// Headers
	w.Header().Set(middleware.RequestIDHeader, middleware.GetReqID(ctx))
	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(httpStatusCode)
	_, _ = w.Write(js)
}
