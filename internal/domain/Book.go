package domain

type Book struct {
	id              int64
	isbn            string
	title           string
	genre           string
	publicationYear int
	author          string
	copies          []BookCopy
}

func NewBook(id int64, isbn, title, genre string, publicationYear int, author string, copies []BookCopy) *Book {
	return &Book{
		id:              id,
		isbn:            isbn,
		title:           title,
		genre:           genre,
		publicationYear: publicationYear,
		author:          author,
		copies:          copies,
	}
}

func (b Book) ISBN() string {
	return b.isbn
}

func (b Book) IsAvailable() bool {
	for _, copy := range b.copies {
		if copy.GetAvailable() {
			return true
		}
	}
	return false
}

func (b Book) GetCopyToBorrow() BookCopy {
	for _, copy := range b.copies {
		if copy.GetAvailable() {
			return copy
		}
	}
	return BookCopy{}
}
