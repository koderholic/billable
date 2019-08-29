package main

import (
	"billable/api"
	"billable/config"
	"flag"
	"fmt"
	"log"
	"net/http"
	// "os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	_ = Config.Init(configDir)
	logPath, port := Config.LogPath, Config.Port

	app.Router = mux.NewRouter()
	app.LogPath = logPath
	app.RegisterRoutes()

	serviceAddress := ":" +port

	// if port, ok := os.LookupEnv("BILLABLEAPI_PORT"); ok {
	// 	serviceAddress = ":" + port
	// }

	allowOriginFunc := func(origin string) bool {
		return true
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowedHeaders: []string{"Accept", "Accept-Encoding", "Content-Type", "*"},
		AllowOriginFunc: allowOriginFunc,
	})

	log.Println("Server started and listening on port ", serviceAddress)

	log.Fatal(http.ListenAndServe(serviceAddress, c.Handler(app.Router)))
}
