package api

import (
	"log"
	"net/http"
	"sync"

	httpSwagger "github.com/swaggo/http-swagger"
)

var (
	once sync.Once
)


func (app *App) RegisterRoutes() {
	once.Do(func() {

		apiRouter := app.Router.PathPrefix("/api").Subrouter()

		app.Router.PathPrefix("/").Handler(httpSwagger.WrapHandler)
		apiRouter.HandleFunc("/ping", app.Ping).Methods(http.MethodGet)
		apiRouter.HandleFunc("/invoice", app.GenerateInvoice()).Methods(http.MethodPost)

	})

	log.Println("App routes registered successfully!")
}
