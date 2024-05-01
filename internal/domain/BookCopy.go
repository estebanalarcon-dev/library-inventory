package domain

type BookCopy struct {
	id        int64
	bookId    int64
	available bool
}

func NewBookCopy(id, bookId int64, available bool) *BookCopy {
	return &BookCopy{
		id:        id,
		bookId:    bookId,
		available: available,
	}
}

func (c BookCopy) GetAvailable() bool {
	return c.available
}

func (c BookCopy) GetID() int64 {
	return c.id
}
