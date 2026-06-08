package routing

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	rootRouter := http.NewServeMux()

	apiRouter := http.NewServeMux()
	rootRouter.Handle("/api/v1/", http.StripPrefix("/api/v1", apiRouter))

	apiRouter.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	registerAuthRoutes(apiRouter)

	return rootRouter
}
