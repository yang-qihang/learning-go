package ex3

import (
	"context"
	"fmt"
	"net/http"
)

type levelKey int

const (
	_ levelKey = iota
	key
)

type Level string

const Debug Level = "debug"

const Info Level = "info"

func ContextWithLevel(ctx context.Context, level Level) context.Context {
	return context.WithValue(ctx, key, level)
}

func LevelFromContext(ctx context.Context) (Level, bool) {
	l, ok := ctx.Value(key).(Level)
	return l, ok
}

func levelMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		// if level := req.URL.Query().Get("log_level"); level != "" {
		// 	switch {
		// 	case level == "debug":
		// 		req = req.WithContext(ContextWithLevel(ctx, Debug))
		// 	case level == "info":
		// 		req = req.WithContext(ContextWithLevel(ctx, Info))
		// 	}
		// }

		// We actually don't need the extra if statement around the switch statement
		// Since if we cannot find a parameter called log_level, we will match on an empty string.
		switch req.URL.Query().Get("log_level") {
		case "debug":
			req = req.WithContext(ContextWithLevel(ctx, Debug))
		case "info":
			// notice here a shallow copy is returned
			// it doesn't modify the original request
			req = req.WithContext(ContextWithLevel(ctx, Info))
		}
		h.ServeHTTP(rw, req)
	})
}

func Log(ctx context.Context, level Level, message string) {
	//	this is redundant
	//	var inLevel Level

	// inLevel means the logging level associated with the incoming request.
	inLevel, ok := LevelFromContext(ctx)

	// If the log level is not assigned not is not a valid value,
	// nothing should be printed.
	if !ok {
		return
	}

	if level == Debug && inLevel == Debug {
		fmt.Println(message)
	}
	if level == Info && (inLevel == Debug || inLevel == Info) {
		fmt.Println(message)
	}
}
