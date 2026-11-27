package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/learning-go-book-2e/ch14/sample_code/context_guid/tracker"
)

// Note that the interface for Logic in the two different packages are DIFFERENT!
type Logic interface {
	QueryHandler(ctx context.Context, query string) (string, error)
}
type Controller struct {
	Logic Logic
}

// the second controller
func (c Controller) Second(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	query := req.URL.Query().Get("query")

	// the second business logic that corresponds to the second controller
	result, err := c.Logic.QueryHandler(ctx, query)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	rw.Write([]byte(result))
}

type Logger interface {
	Log(context.Context, string)
}

type LogicImpl struct {
	Logger Logger
	Remote string
}

func (l LogicImpl) QueryHandler(ctx context.Context, query string) (string, error) {
	l.Logger.Log(ctx, "starting QueryHandler with query: "+query)
	return fmt.Sprintf("got query: '%s' from first", query), nil
}

func main() {
	r := chi.NewRouter()
	r.Use(tracker.Middleware)
	controller := Controller{
		Logic: LogicImpl{
			Logger: tracker.Logger{},
		},
	}

	// we pass a function here
	// and it's converted to be another wrapped type that's a function in essence
	// since it's the underlying type, the conversion happens smoothly
	r.Get("/second", controller.Second)
	http.ListenAndServe(":4000", r)
}
