package routes

import (
	"net/http"

)

func handleRoomRoutes(mux *chi.Mux) {
	roomHandlers := handlers.NewRoomHandler()

	mux.Route("/rooms", func(wsRouter chi.Router) {
		wsRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
			roomHandlers.GetRooms(w, r)
		})
	})
}
