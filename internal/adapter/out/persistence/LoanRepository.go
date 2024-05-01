package persistence

import (
	"github.com/uptrace/bun"
	"libraryInventory/infrastructure"
	"time"
)

type LoanRepository interface {
	ListByUserID(userID int64) ([]loanORM, error)
	Create(loan *loanORM) error
}

type loanRepository struct {
	persistence infrastructure.Persistence
}

func NewLoanRepository(db infrastructure.Persistence) LoanRepository {
	return &loanRepository{persistence: db}
}

type loanORM struct {
	bun.BaseModel      `bun:"table:loans"`
	id                 int64     `bun:"id,pk,autoincrement"`
	borrowedCopyID     int64     `bun:"borrowed_copy_id,notnull"`
	borrowingUserID    int64     `bun:"borrowing_user_id,notnull"`
	loanDate           time.Time `bun:"loan_date,notnull"`
	expectedReturnDate time.Time `bun:"expected_return_date,notnull"`
}

func (l loanRepository) ListByUserID(userID int64) ([]loanORM, error) {
	var loans []loanORM
	err := l.persistence.Select().Model(&loans).Where("borrowing_user_id = ?", userID).Scan(l.persistence.Context())
	return loans, err
}

func (l loanRepository) Create(loan *loanORM) error {
	_, err := l.persistence.Insert().Model(loan).Exec(l.persistence.Context())
	return err

}
