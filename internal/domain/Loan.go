package domain

import "time"

type Loan struct {
	id                 int64
	borrowedCopyID     int64
	borrowingUserID    int64
	loanDate           time.Time
	expectedReturnDate time.Time
}

func NewLoan(loanID, borrowedCopyID, borrowingUserID int64, loanDate, expectedReturnDate time.Time) *Loan {
	return &Loan{
		id:                 loanID,
		borrowedCopyID:     borrowedCopyID,
		borrowingUserID:    borrowingUserID,
		loanDate:           loanDate,
		expectedReturnDate: expectedReturnDate,
	}
}

func (l Loan) ID() int64 {
	return l.id
}

func (l Loan) LoanDate() time.Time {
	return l.loanDate
}

func (l *Loan) BorrowedCopyID() int64 {
	return l.borrowedCopyID
}

func (l *Loan) BorrowingUserID() int64 {
	return l.borrowingUserID
}

func (l *Loan) ExpectedReturnDate() time.Time {
	return l.expectedReturnDate
}

// IsOverdue if expected return date already happened
func (l Loan) IsOverdue() bool {
	return l.expectedReturnDate.Before(time.Now())
}
