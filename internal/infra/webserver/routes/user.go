package routes



func handleUserRoutes(mux *chi.Mux) {
	userHandlers := handlers.NewUserHandler()

	mux.Route("/users", func(wsRouter chi.Router) {
		wsRouter.Post("/", userHandlers.CreateUser)
	})
}
