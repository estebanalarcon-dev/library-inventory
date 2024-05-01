package domain

type User struct {
	id    int64
	name  string
	email string
	loans []Loan
}

func NewUser(id int64, name, email string, loans []Loan) *User {
	return &User{
		id:    id,
		name:  name,
		email: email,
		loans: loans,
	}
}

func (u User) ID() int64 {
	return u.id
}

func (u User) Email() string {
	return u.email
}

// AvailableForLoan User can borrow if it does not have overdue loans
func (u User) AvailableForLoan() bool {
	for _, loan := range u.loans {
		if loan.IsOverdue() {
			return false
		}
	}
	return true
}

/*
func (u User) BorrowBook(book Book, copyBook CopyBook, loanDate time.Time, numberOfDays int) {
	expectedReturnDate := loanDate.AddDate(0,0, numberOfDays)
	loan := NewLoan(0, copyBook.id, u.id, loanDate, expectedReturnDate)
}
*/
