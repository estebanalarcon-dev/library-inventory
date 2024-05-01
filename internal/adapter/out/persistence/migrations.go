package persistence

import (
	"libraryInventory/infrastructure"
)

func CreateUser(db infrastructure.Persistence, id int64, email, name string) error {
	user := userORM{
		ID:    id,
		Email: email,
		Name:  name,
	}
	_, err := db.Insert().Model(&user).Exec(db.Context())
	return err
}

func CreateBook(db infrastructure.Persistence, id int64, isbn, title, genre string, publicationYear int) error {
	book := bookORM{
		ID:              id,
		Isbn:            isbn,
		Title:           title,
		PublicationYear: publicationYear,
		Genre:           genre,
	}
	_, err := db.Insert().Model(&book).Exec(db.Context())
	return err
}

//func Create
