package web

import (
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	. "github.com/onsi/ginkgo/v2"
	"github.com/uptrace/bunrouter"
	"libraryInventory/infrastructure"
	"libraryInventory/internal/adapter/in/web/migrations"
	persistence2 "libraryInventory/internal/adapter/out/persistence"
	"libraryInventory/internal/application/service"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Borrow Book Integration", func() {

	dbTest := infrastructure.NewTestDatabase()
	bookRepository := persistence2.NewBookRepository(dbTest)
	copyBookRepository := persistence2.NewCopyBookRepository(dbTest)
	userRepository := persistence2.NewUserRepository(dbTest)
	loanRepository := persistence2.NewLoanRepository(dbTest)

	bookPersistenceAdapter := persistence2.NewBookPersistenceAdapter(
		bookRepository,
		copyBookRepository)

	userPersistenceAdapter := persistence2.NewUserPersistenceAdapter(
		userRepository,
		loanRepository)
	borrowBookUsecase := service.NewBorrowBookUseCase(
		userPersistenceAdapter, //LoadUserPort
		bookPersistenceAdapter, //LoadBookPort
		userPersistenceAdapter) //CreateLoanPort

	borrowBookController := NewBorrowBookController(borrowBookUsecase)

	It("should create a borrow", func() {
		migrations.Migrate(dbTest.GetDriver())
		persistence2.CreateUser(dbTest, 1, "test@test.com", "test name")
		persistence2.CreateBook(dbTest, 1, "110011", "test book", "fiction", 2020)

		router := bunrouter.New()
		router.POST("/books/borrow/:userEmail/:bookISBN/:numberOfDays", borrowBookController.BorrowBookHandler)

		rr := httptest.NewRecorder()
		defer rr.Result().Body.Close()
		req, _ := http.NewRequest(http.MethodPost, "/books/borrow/test@test.com/110011/10", nil)
		router.ServeHTTP(rr, req)
	})
})
