package persistence

import (
	"github.com/uptrace/bun"
	"libraryInventory/infrastructure"
)

type BookCopyRepository interface {
	ListByBookID(bookID int64) ([]copyBookORM, error)
}

type copyBookRepository struct {
	persistence infrastructure.Persistence
}

func NewCopyBookRepository(db infrastructure.Persistence) BookCopyRepository {
	return &copyBookRepository{persistence: db}
}

type copyBookORM struct {
	bun.BaseModel `bun:"table:book_copies"`
	ID            int64 `bun:"id,pk,autoincrement"`
	BookID        int64 `bun:"book_id,notnull"`
	Available     bool  `bun:"available,notnull"`
}

func (c copyBookRepository) ListByBookID(bookID int64) ([]copyBookORM, error) {
	var copyBooks []copyBookORM
	err := c.persistence.Select().Model(&copyBooks).Where("book_id = ?", bookID).Scan(c.persistence.Context())
	return copyBooks, err
}
