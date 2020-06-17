package handlers

import (
	"github.com/go-chi/chi"
	"github.com/olehmell/simple-task-manager/handlers/column"
	"github.com/olehmell/simple-task-manager/handlers/comment"
	"github.com/olehmell/simple-task-manager/handlers/desk"
	"github.com/olehmell/simple-task-manager/handlers/task"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/desk", func(r chi.Router) {
		r.Get("/all", desk.GetAll)
		r.Get("/{desk_id}", desk.Get)
		r.Delete("/{desk_id}", desk.Delete)

		r.Post("/create", desk.Create)
		r.Put("/update", desk.Update)
	})

	r.Route("/column", func(r chi.Router) {
		r.Get("/all/{desk_id}", column.GetAll)
		r.Get("/{column_id}", column.Get)
		r.Delete("/{column_id}", column.Delete)

		r.Post("/create", column.Create)
		r.Put("/update", column.Update)
	})

	r.Route("/task", func(r chi.Router) {
		r.Get("/all/{column_id}", task.GetAll)
		r.Get("/{task_id}", task.Get)
		r.Delete("/{task_id}", task.Delete)

		r.Post("/create", task.Create)
		r.Put("/update", task.Update)
	})

	r.Route("/comment", func(r chi.Router) {
		r.Get("/all/{task_id}", comment.GetAll)
		r.Get("/{comment_id}", comment.Get)
		r.Delete("/{comment_id}", comment.Delete)

		r.Post("/create", comment.Create)
		r.Put("/update", comment.Update)
	})

	return r
}
