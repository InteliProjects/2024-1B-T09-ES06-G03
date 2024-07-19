package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/cmd/api"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/config"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/db"
)

func main() {
	cfg := config.Envs

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := db.NewPostgresStorage(dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	initStorage(db)

	server := api.NewApiServer(fmt.Sprintf(":%s", cfg.Port), db)
  
	if err := server.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func initStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Database connection pool established")
}
