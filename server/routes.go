package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tejashwikalptaru/tutorial/handler"
)

type Server struct {
	chi.Router
}

func SetupRoutes() *Server {
	router := chi.NewRouter()
	router.Route("/api", func(api chi.Router) {
		api.Get("/welcome", handler.Greet)
	})
	router.Route("/apiAll", func(api chi.Router) {
		api.Get("/allUsers", handler.AllUsers)
	})
	router.Route("/apiAdd", func(api chi.Router) {
		api.Post("/addUser", handler.AddRow)
	})
	router.Route("/apiUpdate", func(api chi.Router) {
		api.Put("/updateUser", handler.UpdateRow)
	})
	router.Route("/apiDelete", func(api chi.Router) {
		api.Delete("/deleteUser", handler.DeleteRow)
	})

	return &Server{router}
}

func (svc *Server) Run(port string) error {
	return http.ListenAndServe(port, svc)
}
