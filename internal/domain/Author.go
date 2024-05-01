package domain

type Author struct {
	id   int64
	name string
}

func NewAuthor(id int64, name string) *Author {
	return &Author{
		id:   id,
		name: name,
	}
}
