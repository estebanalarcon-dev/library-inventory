package persistence

import (
	"github.com/uptrace/bun"
	"libraryInventory/infrastructure"
)

type UserRepository interface {
	FindByEmail(email string) (userORM, error)
}

type userRepository struct {
	persistence infrastructure.Persistence
}

func NewUserRepository(db infrastructure.Persistence) UserRepository {
	return &userRepository{persistence: db}
}

type userORM struct {
	bun.BaseModel `bun:"table:users"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Email         string `bun:"email,notnull"`
	Name          string `bun:"name,notnull"`
}

func (u userRepository) FindByEmail(email string) (userORM, error) {
	var userORM userORM
	u.persistence.Select().Model(&userORM).Where("email = ?", email).Scan(u.persistence.Context())
	return userORM, nil
}
