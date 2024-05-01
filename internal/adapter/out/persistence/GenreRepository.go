package persistence

import "github.com/uptrace/bun"

type genreORM struct {
	bun.BaseModel `bun:"table:genres"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Name          string `bun:"name,notnull"`
}
