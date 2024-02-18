package database

import (
	"context"
	"database/sql"
	"fmt"
	"goth/internal/config"
	"goth/internal/database/schema"
	"goth/internal/slogger"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	Seed()
	GetUsers() []schema.User
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

var DB = newDb()

var logger = slogger.Get()

func newDb() Service {
	// Opening a driver typically will not attempt to connect to the database.
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname))
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		logger.Fatal(err.Error())
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	s := &service{db: db}
	return s
}

func (s *service) Seed() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// DEV: drop table
	if config.App.APP_ENV == "development" {
		_, err := s.db.ExecContext(ctx, schema.DropUserTableSchema)
		if err != nil {
			logger.Fatal(fmt.Sprintf("cannot drop table: %v", err))
			return
		}
	}

	_, err := s.db.ExecContext(ctx, schema.CreateUserTableSchema)
	if err != nil {
		logger.Fatal(fmt.Sprintf("cannot create users schema: %v", err))
		return
	}

	// DEV: seed table
	if config.App.APP_ENV == "development" {
		_, err = s.db.ExecContext(ctx, schema.CreateUserSchema,
			"usr_d895753efd95444da56e2d39bfe1d13a",
			"testuser",
			"test@example.com",
			"password123",
		)
		if err != nil {
			logger.Fatal(fmt.Sprintf("cannot create test user: %v", err))
			return
		}

		logger.Info("database seeded for dev")
	}
}

func (s *service) GetUsers() []schema.User {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, schema.GetUsersSchema)
	if err != nil {
		logger.Fatal(fmt.Sprintf("cannot get users: %v", err))
	}

	users := []schema.User{}

	defer rows.Close()
	for rows.Next() {
		user := schema.User{}
			
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			logger.Error(fmt.Sprintf("cannot scan users: %v", err))
		}
		users = append(users, user)
	}

	return users
}
