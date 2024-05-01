package out

import (
	"libraryInventory/internal/domain"
)

type LoadUserPort interface {
	LoadUser(email string) (*domain.User, error)
}
