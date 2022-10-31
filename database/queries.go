package database

import (
    "log"
    "time"
    "context"
)

func (db *Repository) CreateUsersTable(ctx context.Context) error {
    query := `CREATE TABLE IF NOT EXISTS users(
	    id int primary key auto_increment,
	    username text not null,
	    email text not null,
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
