package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/su-de-sh/nestly/internal/config"
	"github.com/su-de-sh/nestly/internal/database"
	"github.com/su-de-sh/nestly/internal/handlers"
	"github.com/su-de-sh/nestly/internal/repository"
)

func main() {
	appConfig, err := config.Load()
	if err != nil {
		fmt.Println("Error loading config:", err)
	}

	postgresDB, err := database.NewPostgresConnection(appConfig)

	if err != nil {
		fmt.Printf("Error Connecting to Postgres DB %v", err)
	}

	defer postgresDB.Close()

	// Setup for running migrations
	driver, err := postgres.WithInstance(postgresDB.DB, &postgres.Config{})
	if err != nil {
		fmt.Printf("Error getting Postgres driver %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		// "file://absolute/path/to/migrations",
		"file://database/migrations",
		"postgres", driver)
	if err != nil {
		fmt.Printf("Error creating migrate instance %v", err)
	}

	// Using `m.Up()` runs all pending migrations in order.
	// We avoid `m.Steps()` here because it could leave the DB in a partial state.
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Printf("Error running migrations %v", err)
	} else {
		fmt.Println("DB Migration ran successfully")
	}

	babySitterRepository := repository.NewBabySitterRepository(postgresDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hello, world! Welcome to Nestly!")
	})

	babysitterHandler := handlers.NewBabysitterHandler(babySitterRepository)

	//babysitter routes
	r.Get("/babysitters", babysitterHandler.GetAll)
	r.Post("/babysitters", babysitterHandler.Create)
	r.Put("/babysitters/{id}", babysitterHandler.UpdateById)
	r.Delete("/babysitters/{id}", babysitterHandler.DeleteById)

	//health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		responose := map[string]string{"msg": "Hello world! Welcome to Nestly"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responose)
	})

	http.ListenAndServe(":"+appConfig.Server.Port, r)
}
