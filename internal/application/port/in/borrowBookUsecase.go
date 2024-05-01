package in

import "time"

type BorrowBookUsecase interface {
	BorrowBook(BorrowBookInput) (*BorrowBookOutput, error)
}

type BorrowBookInput struct {
	UserEmail    string
	BookISBN     string
	NumberOfDays int
}

type BorrowBookOutput struct {
	LoanID             int64     `json:"loan_id"`
	BorrowingUserEmail string    `json:"borrowing_user_email"`
	BorrowingBookISBN  string    `json:"borrowing_book_isbn"`
	LoanDate           time.Time `json:"loan_date"`
	ReturnDate         time.Time `json:"return_date"`
}

func NewBorrowBookInput(userEmail string, bookISBN string, numberOfDays int) BorrowBookInput {
	return BorrowBookInput{
		UserEmail:    userEmail,
		BookISBN:     bookISBN,
		NumberOfDays: numberOfDays,
	}
}
