package app

import (
	"Honest-Game-Reviews/src/datasource/mysql/database"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
)

var (
	listenAddr string
)


func StartApplication() {
	flag.StringVar(&listenAddr, "listen-addr", ":8080", "server listen address")

	database.DatabaseClient.NewDatabase()

	router := chi.NewRouter()
	// routes
	gamesURLMapping(router)
	usersURLMapping(router)
	usersAuthenticationsURLMapping(router)


	server := http.Server{
		Addr:           listenAddr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	quit := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		fmt.Println("server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15)
		defer cancel()


		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("could not gracefully shutdown server %v\n",  err)
		}

		close(done)
	}()

	fmt.Printf("server is ready to handle requests at %s\n", listenAddr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Could not listen on %s: %v\n", listenAddr, err)
	}
	<-done
	fmt.Println("server stopped")
}

