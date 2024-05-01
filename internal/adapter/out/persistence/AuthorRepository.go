package persistence

import "github.com/uptrace/bun"

type authorORM struct {
	bun.BaseModel `bun:"table:authors"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Name          string `bun:"name,notnull"`
}
