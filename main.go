package main

import (
	"billable/api"
	"billable/config"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)


var (
	app = &api.App{}
)


func main() {

	var configDir string

	flag.StringVar(&configDir, "config", ".", "directory location of config file")
	flag.Parse()

	if configDir == "" {
		fmt.Printf("----------------------------------------------------------\n")
		flag.Usage()
		fmt.Printf("----------------------------------------------------------\n")
	}

	Config := config.Data{}
	Config.Init(configDir)
	logPath := Config.LogPath

	app.Router = mux.NewRouter()
	app.LogPath = logPath
	app.RegisterRoutes()

	serviceAddress := ":8100"

	if port, ok := os.LookupEnv("ASPNETCORE_PORT"); ok {
		serviceAddress = ":" + port
	}

	log.Println("Server started and listening on port ", serviceAddress)
	log.Fatal(http.ListenAndServe(serviceAddress, app.Router))
}
