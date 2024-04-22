package repositories

import (
	"database/sql"
	"fmt"
)

type DB struct {
	Conn *sql.DB
}

func NewDB(username, password, dbname, port string) (*DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", username, password, dbname, port)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return &DB{Conn: db}, nil
}
