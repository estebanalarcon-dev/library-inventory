package persistence

import (
	"fmt"
	"libraryInventory/internal/domain"
)

type UserPersistenceAdapter struct {
	userRepository UserRepository
	loanRepository LoanRepository
}

func NewUserPersistenceAdapter(userRepository UserRepository, loanRepository LoanRepository) UserPersistenceAdapter {
	return UserPersistenceAdapter{
		userRepository: userRepository,
		loanRepository: loanRepository,
	}
}

func (u UserPersistenceAdapter) LoadUser(email string) (*domain.User, error) {
	user, err := u.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	fmt.Println("USER: ", user)

	loans, err := u.loanRepository.ListByUserID(user.ID)
	if err != nil {
		return nil, err
	}

	return mapToUserDomain(user, loans), nil
}

func (u UserPersistenceAdapter) CreateLoan(loan *domain.Loan) error {
	return u.loanRepository.Create(mapToLoanORM(loan))
}
