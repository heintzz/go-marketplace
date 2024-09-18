package auth

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func Run(router chi.Router, db *sql.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	router.Route("/v1/auth", func(r chi.Router) {
		r.Post("/register/user", handler.registerUserHandler)		
		r.Post("/login/user", handler.loginUserHandler)		
	})
}

