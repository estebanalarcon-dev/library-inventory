package domain

type LoanPolicy struct{}

func NewLoanPolicy() *LoanPolicy {
	return &LoanPolicy{}
}

func (p LoanPolicy) IsAllowed(numberOfDays int) bool {
	if numberOfDays > 1 && numberOfDays <= 21 {
		return true
	}
	return false
}
