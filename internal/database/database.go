package database

import (
	"context"
	"database/sql"
	"fmt"
	"goth/internal/config"
	"goth/internal/slogger"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	Health() map[string]string
	Seed()
	GetUsers() map[int]string
}

type service struct {
	db *sql.DB
}

var (
	dbname   = config.App.DB_DATABASE
	password = config.App.DB_PASSWORD
	username = config.App.DB_USERNAME
	port     = config.App.DB_PORT
	host     = config.App.DB_HOST
)

func New() Service {
	// Opening a driver typically will not attempt to connect to the database.
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname))
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		slogger.Log.Fatal(err.Error())
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	s := &service{db: db}
	return s
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		// log.Fatalf(fmt.Sprintf("db down: %v", err))
		slogger.Log.Fatal(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

func (s *service) Seed() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// create table
	_, err := s.db.ExecContext(
		ctx, "CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, email VARCHAR(255), password VARCHAR(255))",
	)
	if err != nil {
		// log.Fatalf(fmt.Sprintf("db down: %v", err))
		slogger.Log.Fatal(fmt.Sprintf("db down: %v", err))
	}

	// delete all records
	_, err = s.db.ExecContext(ctx, "DELETE FROM users")
	if err != nil {
		// log.Fatalf(fmt.Sprintf("db down: %v", err))
		slogger.Log.Fatal(err.Error())
	}

	// seed the database
	_, err = s.db.ExecContext(
		ctx, "INSERT INTO users (email, password) VALUES (?, ?)",
		"seededuser@example.com", "password123",
	)
	if err != nil {
		slogger.Log.Fatal(fmt.Sprintf("db down: %v", err))
	}
}

func (s *service) GetUsers() map[int]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		slogger.Log.Fatal(fmt.Sprintf("db down: %v", err))
	}

	users := make(map[int]string)

	defer rows.Close()
	for rows.Next() {
		var id int
		var email, password string
		err := rows.Scan(&id, &email, &password)
		if err != nil {
			slogger.Log.Fatal(fmt.Sprintf("db down: %v", err))
		}
		users[id] = email
	}

	return users
}

// func (s *service) NewUser() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()
//
// 	res, err := s.db.ExecContext(
// 		ctx, "INSERT INTO users (email, password) VALUES (?, ?)",
// 		"testemail@example.com", "password",
// 	)
// 	if err != nil {
// 		log.Fatalf(fmt.Sprintf("db down: %v", err))
// 	}
//
// 	id, err := res.LastInsertId()
// 	if err != nil {
// 		log.Fatalf(fmt.Sprintf("db down: %v", err))
// 	}
//
// 	fmt.Printf("Last insert id: %d\n", id)
// }
