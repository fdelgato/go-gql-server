package auth

import (
	"context"
	"github.com/fdelgato/go-gql-server/internal/gql/resolvers"
	"net/http"

)

// Middleware is used to handle auth logic
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		auth := r.Header.Get("Authorization")
		if auth != "" {
			// Write your fancy token introspection logic here and if valid user then pass appropriate key in header
			// IMPORTANT: DO NOT HANDLE UNAUTHORIZED USER HERE
			ctx = context.WithValue(ctx, resolvers.UserIDCtxKey, auth)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
