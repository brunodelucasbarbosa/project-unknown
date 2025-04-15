package api

import (
	"net/http"
	"time"

	"github.com/brunodelucasbarbosa/project-unknown/infrastructure/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
)

func Router(handlerUser handlers.IHandlerUser) http.Handler {
	r := chi.NewRouter()
	r.Use(httprate.Limit(
		10,
		time.Minute,
		httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, `{"error": "Rate-limited. Please, slow down."}`, http.StatusTooManyRequests)
		}),
	))

	r.Mount("/api/v1", v1RouterUsers(r, handlerUser))

	return r
}

func v1RouterUsers(r *chi.Mux, handlerUser handlers.IHandlerUser) http.Handler {
	r.Route("/api/v1/users", func(r chi.Router) {
		r.Post("/login", handlerUser.HandleLogin)
		r.Post("/create", handlerUser.HandleCreateUser)
	})

	r.Route("/api/v1/users/{id}/post", func(r chi.Router) {
		//r.Post("/create", handlerUser.HandleCreatePost)
		//r.Get("/get", handlerUser.HandleGetPost)

		r.Route("/{postId}/comment", func(r chi.Router) {
			//r.Post("/create", handlerUser.HandleCreateComment)
			//r.Get("/get", handlerUser.HandleGetComment)
		})
	})
	return r
}
