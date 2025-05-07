package service

import (
	"fmt"
	"net/url"
	"sync"

	"github.com/izzul-ali/book-management/src/model"
	"github.com/izzul-ali/book-management/src/utils/helper"
)

// BookStorage is a singleton in-memory storage for books.
type BookStorage struct {
	books map[string]model.Book
	mu    sync.RWMutex
}

var (
	bookStorageInstance *BookStorage

	// Singleton instance and sync.Once to ensure the storage is initialized only once
	once sync.Once
)

// InitBookStorage returns the singleton instance of BookStorage.
func InitBookStorage() *BookStorage {

	once.Do(func() {
		bookStorageInstance = &BookStorage{
			books: make(map[string]model.Book),
		}
	})

	return bookStorageInstance
}

// Lock the mutex to ensure exclusive access to the shared resource
// This prevents other goroutines from reading or writing until the lock is released
// s.mu.Lock()

// Ensure the mutex is unlocked after the function completes
// even if an error or panic occurs
// defer s.mu.Unlock()

// Add a new book and save it in memory
func (s *BookStorage) Create(book model.Book) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.books[book.ID] = book
}

// Find all books with search query
func (s *BookStorage) GetAll(query url.Values) []model.Book {
	s.mu.Lock()
	defer s.mu.Unlock()

	var search = query.Get("search")

	result := make([]model.Book, 0, len(s.books))

	// get key value
	for _, book := range s.books {
		if search == "" || helper.ContainsIgnoreCase(book.Title, search) || helper.ContainsIgnoreCase(book.Author, search) {
			result = append(result, book)
		}
	}

	return result
}

// Find a book by id
func (s *BookStorage) GetOne(id string) (model.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	book, exists := s.books[id]

	if !exists {
		return model.Book{}, fmt.Errorf("book with id %s not found", id)
	}

	return book, nil
}

// Update book by id and update its value
func (s *BookStorage) Update(book model.Book) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.books[book.ID]

	if !exists {
		return fmt.Errorf("book with id %s not found", book.ID)
	}

	s.books[book.ID] = book

	return nil
}

// Delete book by id
func (s *BookStorage) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.books[id]

	if !exists {
		return fmt.Errorf("book with id %s not found", id)
	}

	delete(s.books, id)

	return nil
}
