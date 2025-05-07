package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/izzul-ali/book-management/src/handler"
	"github.com/izzul-ali/book-management/src/service"
)

// InitRoutes initializes all application routes
func InitRoutes(c *chi.Mux, s *service.BookStorage) {
	// Initialize the book handler with dependency injection of BookStorage
	bookHandler := handler.NewBookHandler(s)

	// Define routes under the /books path and attach corresponding handlers
	c.Route("/books", BookRoute(c, bookHandler))
}
