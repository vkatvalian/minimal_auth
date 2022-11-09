package database

import (
    "os"
    "log"
    "context"
    "github.com/joho/godotenv"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Repository struct {
    Conn *sql.DB
}

func Connect(ctx context.Context) *Repository {
    err := godotenv.Load()
    if err != nil {
      log.Fatal("Error loading .env file")
    }
  
    DSN := os.Getenv("DSN")

    db, err := sql.Open("mysql", DSN)
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }

    repo := &Repository{db}
    repo.CreateTables(ctx)
    return repo
}
