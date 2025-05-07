package validator

import (
	"errors"
	"strings"

	"github.com/izzul-ali/book-management/src/model"
)

// Validation value
func BookValidator(b model.Book) error {
	if strings.TrimSpace(b.Title) == "" {
		return errors.New("title is required")
	}

	if strings.TrimSpace(b.Author) == "" {
		return errors.New("author is required")
	}

	if strings.TrimSpace(b.Content) == "" {
		return errors.New("content is required")
	}

	return nil
}
