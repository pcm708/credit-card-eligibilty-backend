package main

import (
	"context"
	"fmt"
	"github.com/honestbank/tech-assignment-backend-engineer/db"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/honestbank/tech-assignment-backend-engineer/routes"
	env "github.com/joho/godotenv"
)

const (
	envFile = ".env"
)

var loadEnv = env.Load

func run() (s *http.Server) {
	err := loadEnv(envFile)
	if err != nil {
		log.Fatal(err)
	}

	port, exist := os.LookupEnv("PORT")
	if !exist {
		log.Fatal("no port specified")
	}

	port = fmt.Sprintf(":%s", port)

	mux := routes.SetupRoutes()

	s = &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        mux,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error listening on port: %s\n", err)
		}
	}()

	return
}

func main() {
	s := run()
	db.ConnectToDB()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shut down")
	}
	log.Println("server exiting")
}
