package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/izzul-ali/book-management/src/model"
	"github.com/izzul-ali/book-management/src/service"
	"github.com/izzul-ali/book-management/src/utils/helper"
	"github.com/izzul-ali/book-management/src/validator"
)

// BookHandler handles HTTP requests related to book resources
type BookHandler struct {
	storage *service.BookStorage
}

// NewBookHandler constructs a new BookHandler, injecting the BookStorage dependency
func NewBookHandler(s *service.BookStorage) *BookHandler {
	return &BookHandler{storage: s}
}

func (h *BookHandler) GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	// Get all query from url
	query := r.URL.Query()

	books := h.storage.GetAll(query)

	helper.ResponseJSON(w, http.StatusOK, "Success!", books)
}

func (h *BookHandler) GetBookHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	book, err := h.storage.GetOne(id)

	if err != nil {
		helper.ResponseError(w, http.StatusNotFound, err.Error())
		return
	}

	helper.ResponseJSON(w, http.StatusOK, "Success!", book)
}

func (h *BookHandler) CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	var book model.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Invalid request body!")
		return
	}

	if err := validator.BookValidator(book); err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Generate unique id
	book.ID = uuid.NewString()

	h.storage.Create(book)

	helper.ResponseJSON(w, http.StatusOK, "Book successfully created!", nil)
}

func (h *BookHandler) UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	var book model.Book

	id := chi.URLParam(r, "id")

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Invalid request body!")
		return
	}

	if err := validator.BookValidator(book); err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	book.ID = id

	err := h.storage.Update(book)

	if err != nil {
		helper.ResponseError(w, http.StatusNotFound, err.Error())
		return
	}

	helper.ResponseJSON(w, http.StatusOK, "Book successfully updated!", book)
}

func (h *BookHandler) DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.storage.Delete(id)

	if err != nil {
		helper.ResponseError(w, http.StatusNotFound, err.Error())
		return
	}

	helper.ResponseJSON(w, http.StatusOK, "Book successfully deleted!", nil)
}
