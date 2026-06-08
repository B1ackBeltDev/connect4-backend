package routing

import (
	"net/http"

	"github.com/B1ackBeltDev/connect4-backend/internals/utils"
)

func registerAuthRoutes(router *http.ServeMux) {
	authRouter := http.NewServeMux()
	router.Handle("/auth/", http.StripPrefix("/auth", authRouter))

	authRouter.HandleFunc("POST /register", utils.HelloWorldHandler)
	authRouter.HandleFunc("GET /login", utils.HelloWorldHandler)
}
