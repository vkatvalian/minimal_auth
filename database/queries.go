package database

import (
    "fmt"
    "log"
    "time"
    "context"
    "errors"
)

func (db *Repository) CreateTables(ctx context.Context) error {
    query := `CREATE TABLE IF NOT EXISTS users(
	    id int primary key auto_increment,
	    username text not null unique,
	    email text not null unique,
	    password text not null
    )`

    ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    res, err := db.Conn.ExecContext(ctx, query)
    if err != nil {
        log.Println("Error %s when creating product table", err)
	return err
    }

    rows, err := res.RowsAffected()
    if err != nil {
        log.Printf("Error %s when getting rows affected", err)
        return err
    }

    log.Printf("Rows affected when creating table: %d", rows)
    return nil
}

func (db *Repository) InsertUsers(ctx context.Context, username, email, password string) error {
    query := `INSERT INTO users(username, email, password) VALUES (?, ?, ?);`
    ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
    defer cancelfunc()

    stmt, err := db.Conn.PrepareContext(ctx, query)
    if err != nil {
	return err
    }

    defer stmt.Close()
    res, err := stmt.ExecContext(ctx, username, email, password)
    if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
	custom_err := errors.New("User exists")
	return custom_err
    }

    rows, err := res.RowsAffected()
    if err != nil {
	return err
    }
    log.Printf("Rows affected when creating table: %d", rows)
    return err
}

func (db *Repository) FetchUsers(ctx context.Context, _name string) (string, string, string, error) {
    var username, email, password string

    query := `SELECT username, email, password from users where username = ?`
    ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
    defer cancelfunc()

    err := db.Conn.QueryRow(query, _name).Scan(&username, &email, &password)
    if err != nil {}

    return username, email, password, nil
}
