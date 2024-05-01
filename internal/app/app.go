package app

import (
	"github.com/uptrace/bunrouter"
	"libraryInventory/internal/adapter/in/web"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	router http.Handler
}

func NewRouter(borrowBook web.BorrowBookController) http.Handler {
	router := bunrouter.New()
	router.NewGroup("/v1")
	router.POST("/books/borrow/:userEmail/:bookISBN/:numberOfDays", borrowBook.BorrowBookHandler)
	return router
}

func NewApp(router http.Handler) *App {
	return &App{router: router}
}

func (a *App) Start() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
}
