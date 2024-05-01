package web

import (
	"fmt"
	"github.com/uptrace/bunrouter"
	"libraryInventory/internal/application/port/in"
	"net/http"
	"strconv"
)

type BorrowBookController interface {
	BorrowBookHandler(w http.ResponseWriter, req bunrouter.Request) error
}

type borrowBookController struct {
	borrowBookUseCase in.BorrowBookUsecase
}

func NewBorrowBookController(borrowBookUseCase in.BorrowBookUsecase) BorrowBookController {
	return &borrowBookController{borrowBookUseCase: borrowBookUseCase}
}

func (s borrowBookController) BorrowBookHandler(w http.ResponseWriter, req bunrouter.Request) error {
	//think about how validate parameters
	userEmail := req.Params().ByName("userEmail")
	bookISBN := req.Params().ByName("bookISBN")
	auxDays := req.Params().ByName("numberOfDays")
	numberOfDays, _ := strconv.Atoi(auxDays)

	borrowBookInput := in.NewBorrowBookInput(userEmail, bookISBN, numberOfDays)
	fmt.Println("INPUT:", borrowBookInput)
	output, err := s.borrowBookUseCase.BorrowBook(borrowBookInput)
	fmt.Println("OUTPUT:", output)
	fmt.Println("ERR:", err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, err)
	}

	w.WriteHeader(http.StatusOK)
	return bunrouter.JSON(w, output)
}
