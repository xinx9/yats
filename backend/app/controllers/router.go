package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	user "yats/app/models"
)

func SetupServer() {
	fmt.Printf("Server starting on :8080\n")
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	userHandler := UserHandler{
		storage: user.UserStore{},
	}

	r.Route("/admin", func(r chi.Router) {

		r.Get("/", userHandler.ListUsers)
		r.Post("/", userHandler.CreateUser)
		r.Get("/{id}", userHandler.GetUsers)
		r.Put("/update/{id}", userHandler.UpdateUser)
		r.Delete("/delete/{id}", userHandler.DeleteUser)

	})

	//r.Mount("/admin", UserRoute())
	http.ListenAndServe(":8080", r)
}

func UserRoute() http.Handler {
	r := chi.NewRouter()

	return r

}
