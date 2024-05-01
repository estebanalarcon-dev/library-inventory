package persistence

import (
	"libraryInventory/internal/domain"
)

func mapToUserDomain(userORM userORM, loansORM []loanORM) *domain.User {
	var loans []domain.Loan
	for _, loanORM := range loansORM {
		loans = append(loans,
			*domain.NewLoan(loanORM.id,
				loanORM.borrowedCopyID,
				loanORM.borrowingUserID,
				loanORM.loanDate,
				loanORM.expectedReturnDate),
		)
	}
	return domain.NewUser(userORM.ID, userORM.Name, userORM.Email, loans)
}

func mapToBookDomain(bookORM bookORM, copiesORM []copyBookORM) *domain.Book {

	return domain.NewBook(bookORM.ID,
		bookORM.Isbn,
		bookORM.Title,
		bookORM.Genre,
		bookORM.PublicationYear,
		bookORM.Author,
		mapToBookCopiesDomain(copiesORM))
}

func mapToLoanORM(loan *domain.Loan) *loanORM {
	return &loanORM{
		id:                 loan.ID(),
		borrowedCopyID:     loan.BorrowedCopyID(),
		borrowingUserID:    loan.BorrowingUserID(),
		loanDate:           loan.LoanDate(),
		expectedReturnDate: loan.ExpectedReturnDate(),
	}
}

func mapToAuthorsDomain(authorsOrm []authorORM) []domain.Author {
	var authors []domain.Author
	for _, orm := range authorsOrm {
		authors = append(authors, *domain.NewAuthor(orm.ID, orm.Name))
	}
	return authors
}

func mapToBookCopiesDomain(copiesORM []copyBookORM) []domain.BookCopy {
	var copies []domain.BookCopy
	for _, copyORM := range copiesORM {
		copies = append(copies,
			*domain.NewBookCopy(copyORM.ID, copyORM.BookID, copyORM.Available))
	}
	return copies
}
