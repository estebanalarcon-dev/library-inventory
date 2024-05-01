package out

import (
	"libraryInventory/internal/domain"
)

type LoadBookPort interface {
	LoadBook(isbn string) (*domain.Book, error)
}
