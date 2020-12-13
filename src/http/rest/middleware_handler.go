package middleware_handler

import (
	"context"
	"net/http"
)



func AuthenticateEndpoint(ctx context.Context, h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){

	}
}

