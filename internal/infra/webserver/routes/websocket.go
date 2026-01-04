package routes


func handleWebsocketRoutes(mux *chi.Mux) {
	websocketHandlers := handlers.NewWebsocketHandler()

	mux.Route("/ws", func(wsRouter chi.Router) {
		wsRouter.Get("/{roomId}", func(w http.ResponseWriter, r *http.Request) {
			roomId := chi.URLParam(r, "roomId")

			websocketHandlers.ServeWs(roomId, w, r)
		})
	})
}
