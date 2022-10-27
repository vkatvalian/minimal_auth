package database

import (
    "os"
    "log"
    "context"
    "github.com/jackc/pgx/v4"

    "gitlab.com/kamee/picverse/config"
)

type Repository struct {
Conn *pgx.Conn
}

func Connection(ctx context.Context) *Repository {
    conn, err := pgx.Connect(ctx, config.DBConfig().DSN)
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }

    repo := &Repository{conn}
    return repo
}
