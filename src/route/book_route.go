package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/izzul-ali/book-management/src/handler"
)

// BookRoute defines all the HTTP routes related to book resources
func BookRoute(c *chi.Mux, h *handler.BookHandler) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", h.GetBooksHandler)
		r.Get("/{id}", h.GetBookHandler)
		r.Post("/", h.CreateBookHandler)
		r.Put("/{id}", h.UpdateBookHandler)
		r.Delete("/{id}", h.DeleteBookHandler)
	}
}
