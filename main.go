package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/izzul-ali/book-management/src/route"
	"github.com/izzul-ali/book-management/src/service"
	"github.com/izzul-ali/book-management/src/utils/helper"
)

func main() {
	//
	r := chi.NewRouter()

	// Define middleware
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.Logger)

	// Global 405 Method Not Allowed
	r.MethodNotAllowed(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helper.ResponseError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
	}))

	// Global 404 Not Found
	r.NotFound(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helper.ResponseError(w, http.StatusNotFound, "Not Found")
	}))

	// Initial storage
	storage := service.InitBookStorage()

	// Register all route
	route.InitRoutes(r, storage)

	log.Println("Server running on :3000")
	http.ListenAndServe(":3000", r)
}
