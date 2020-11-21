package app

import (
	"Honest-Game-Reviews/src/datasource/mysql/database"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

var (
	listenAddr string
)


func StartApplication() {
	flag.StringVar(&listenAddr, "listen-addr", ":8080", "server listen address")

	database.NewDatabase()

	router := chi.NewRouter()
	gamesURLMapping(router)
	reviewsURLMapping(router)

	server := http.Server{
		Addr:           listenAddr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Running server...")


	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

