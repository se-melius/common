package xhttp

import (
	"context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type ContextWrapper func(context.Context) context.Context

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRouter(contextWrapper ContextWrapper, routes ...Route) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()
				r = r.WithContext(contextWrapper(ctx))
				handlers.LoggingHandler(os.Stdout, handler).ServeHTTP(w, r)
			}))
	}
	return router
}