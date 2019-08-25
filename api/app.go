package api

import (
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	LogPath string
}