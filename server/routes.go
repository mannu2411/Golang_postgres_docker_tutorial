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
		api.Get("/allUsers", handler.AllUsers)
		api.Get("/getTasks", handler.GetTasks)
		api.Post("/addUser", handler.AddRow)
		api.Post("/signin", handler.SignInUser)
		api.Post("/signout", handler.SignOut)
		api.Post("/addTask", handler.AddTask)
		api.Put("/completeTask", handler.CompleteTask)
		api.Put("/updateUser", handler.UpdateRow)
		api.Delete("/deleteTask", handler.DeleteTask)
		api.Delete("/deleteUser", handler.DeleteRow)
	})

	return &Server{router}
}

func (svc *Server) Run(port string) error {
	return http.ListenAndServe(port, svc)
}
