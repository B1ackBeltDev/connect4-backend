package routing

import (
	"net/http"

	"github.com/B1ackBeltDev/connect4-backend/internals/utils"
)

func RegisterMatchmakingRoutes(router *http.ServeMux) {
	mmRouter := http.NewServeMux()
	router.Handle("/matchmaking/", http.StripPrefix("/matchmaking", mmRouter))

	mmRouter.HandleFunc("POST /join_queue", utils.HelloWorldHandler)
	mmRouter.HandleFunc("POST /create_match", utils.HelloWorldHandler)
	mmRouter.HandleFunc("GET /match/{match_token}", utils.HelloWorldHandler)
}
