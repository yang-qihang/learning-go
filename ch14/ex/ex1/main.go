package ex1

import (
	"context"
	"net/http"
	"time"
)

func TimeoutMiddleware(milliseconds int) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			req = req.WithContext(ctx)

			ctx, cancel := context.WithTimeout(
				ctx,
				time.Duration(milliseconds)*time.Millisecond,
			)
			defer cancel()

			req = req.WithContext(ctx)

			// The middleware establishes the deadline,
			// but trusts the wrapped handler to respect it!!

			h.ServeHTTP(rw, req)
		})
	}
}
