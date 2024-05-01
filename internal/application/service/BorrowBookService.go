package service

import (
	"errors"
	"fmt"
	"libraryInventory/internal/application/port/in"
	"libraryInventory/internal/application/port/out"
	"libraryInventory/internal/domain"
	"time"
)

type borrowBookService struct {
	loadUserPort   out.LoadUserPort
	loadBookPort   out.LoadBookPort
	createLoanPort out.CreateLoanPort
}

func NewBorrowBookUseCase(
	loadUserPort out.LoadUserPort,
	loadBookPort out.LoadBookPort,
	createLoanPort out.CreateLoanPort) in.BorrowBookUsecase {
	return &borrowBookService{
		loadBookPort:   loadBookPort,
		loadUserPort:   loadUserPort,
		createLoanPort: createLoanPort,
	}
}

func (b borrowBookService) BorrowBook(input in.BorrowBookInput) (*in.BorrowBookOutput, error) {
	fmt.Printf("Usecase.BorrowBook: %+v\n", input)
	user, err := b.loadUserPort.LoadUser(input.UserEmail)
	if err != nil {
		return nil, err
	}

	book, err := b.loadBookPort.LoadBook(input.BookISBN)
	if err != nil {
		return nil, err
	}

	err = b.checkBusinessValidations(user, book, input.NumberOfDays)
	if err != nil {
		return nil, err
	}

	copyToBorrow := book.GetCopyToBorrow()

	loan := domain.NewLoan(0,
		copyToBorrow.GetID(),
		user.ID(),
		time.Now(),
		time.Now().AddDate(0, 0, input.NumberOfDays))

	err = b.createLoanPort.CreateLoan(loan)
	if err != nil {
		return nil, err
	}

	return &in.BorrowBookOutput{
		LoanID:             loan.ID(),
		BorrowingUserEmail: user.Email(),
		BorrowingBookISBN:  book.ISBN(),
		LoanDate:           loan.LoanDate(),
		ReturnDate:         loan.ExpectedReturnDate(),
	}, nil
}

func (b borrowBookService) checkBusinessValidations(user *domain.User, book *domain.Book, numberOfDays int) error {
	if !user.AvailableForLoan() {
		return errors.New("user can't borrow books")
	}

	if !book.IsAvailable() {
		return errors.New("book is not available for loan")
	}

	if domain.LoanPolicy.IsAllowed(domain.LoanPolicy{}, numberOfDays) {
		return errors.New("number of days to borrow is not allowed. It must be 1-21 days")
	}

	return nil
}
