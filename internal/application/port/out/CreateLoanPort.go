package out

import "libraryInventory/internal/domain"

type CreateLoanPort interface {
	CreateLoan(loan *domain.Loan) error
}
