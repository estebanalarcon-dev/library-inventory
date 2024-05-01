package main

import (
	"libraryInventory/infrastructure"
	"libraryInventory/internal/adapter/in/web"
	persistence2 "libraryInventory/internal/adapter/out/persistence"
	"libraryInventory/internal/app"
	"libraryInventory/internal/application/service"
	"libraryInventory/internal/config"
)

func main() {
	app := buildInManually()
	app.Start()
}

// Assembling via Plain Code, with manually Dependency Injection
func buildInManually() *app.App {
	dbConf := config.NewPostgresConf()
	persistence := infrastructure.NewPersistence(dbConf)
	bookRepository := persistence2.NewBookRepository(persistence)
	copyBookRepository := persistence2.NewCopyBookRepository(persistence)
	userRepository := persistence2.NewUserRepository(persistence)
	loanRepository := persistence2.NewLoanRepository(persistence)

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

	borrowBookController := web.NewBorrowBookController(borrowBookUsecase)

	router := app.NewRouter(borrowBookController)
	return app.NewApp(router)
}
