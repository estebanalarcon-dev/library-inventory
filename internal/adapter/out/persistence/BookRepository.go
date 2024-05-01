package persistence

import (
	"github.com/uptrace/bun"
	"libraryInventory/infrastructure"
)

type BookRepository interface {
	FindByISBN(isbn string) (bookORM, error)
}

type bookRepository struct {
	persistence infrastructure.Persistence
}

func NewBookRepository(db infrastructure.Persistence) BookRepository {
	return &bookRepository{persistence: db}
}

type bookORM struct {
	bun.BaseModel   `bun:"table:books"`
	ID              int64  `bun:"id,pk,autoincrement"`
	Isbn            string `bun:"isbn,notnull"`
	Title           string `bun:"title,notnull"`
	PublicationYear int    `bun:"publication_year,notnull"`
	Genre           string `bun:"genre,notnull"`
	Author          string `bun:"author,notnull"`
}

func (b bookRepository) FindByISBN(isbn string) (bookORM, error) {
	var book bookORM
	err := b.persistence.Select().Model(&book).Where("isbn = ?", isbn).Scan(b.persistence.Context())
	return book, err
}
