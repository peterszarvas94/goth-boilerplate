package database

import (
	"context"
	"database/sql"
	"fmt"
	"goth/internal/config"
	"goth/internal/database/schema"
	"goth/internal/slogger"
	"time"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	Seed() error
	UserGetAll() ([]*schema.UserPropsGetAll, error)
	UserSignup(schema.UserPropsSignup) error
	UserSignin(schema.UserPropsSignin) (*schema.User, error)
}

type service struct {
	db *sql.DB
}

var DB = newDb()

var logger = slogger.Get()

func newDb() Service {
	// Opening a driver typically will not attempt to connect to the database.
	connectionStr := fmt.Sprintf("%s?authToken=%s", config.App.DB_URL, config.App.DB_TOKEN)

	db, err := sql.Open("libsql", connectionStr)
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

func (s *service) Seed() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// DEV: drop table
	if config.App.APP_ENV == "development" {
		_, err := s.db.ExecContext(ctx, schema.UserTableSchemaDrop)
		if err != nil {
			logger.Fatal(fmt.Sprintf("cannot drop table: %v", err))
			return err
		}
	}

	_, err := s.db.ExecContext(ctx, schema.UserTableSchemaCreate)
	if err != nil {
		logger.Fatal(fmt.Sprintf("cannot create users schema: %v", err))
		return err
	}

	// DEV: seed table
	if config.App.APP_ENV == "development" {
		_, err = s.db.ExecContext(ctx, schema.UserSchemaSignup,
			"usr_d895753efd95444da56e2d39bfe1d13a",
			"a",
			"a@a",
			"a",
		)
		if err != nil {
			logger.Fatal(fmt.Sprintf("cannot create test user: %v", err))
			return err
		}

		logger.Info("database seeded for dev")
	}

	return nil
}

func (s *service) UserGetAll() ([]*schema.UserPropsGetAll, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, schema.UserSchemaGetAll)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot get users: %v", err))
		return nil, err
	}

	users := []*schema.UserPropsGetAll{}

	defer rows.Close()
	for rows.Next() {
		user := &schema.UserPropsGetAll{}

		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			logger.Error(fmt.Sprintf("cannot scan users: %v", err))
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *service) UserSignup(props schema.UserPropsSignup) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// TODO: hash password
	// TODO: check if user already exists

	_, err := s.db.ExecContext(ctx, schema.UserSchemaSignup,
		props.Id,
		props.Username,
		props.Email,
		props.Password,
	)

	if err != nil {
		logger.Error(fmt.Sprintf("cannot create user: %v", err))
		return err
	}

	logger.Info(fmt.Sprintf("user created: %s", props.Id))

	return nil
}

func (s *service) UserSignin(props schema.UserPropsSignin) (*schema.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// TODO: hash password

	row := s.db.QueryRowContext(ctx, schema.UserSchemaSignin,
		props.UsernameOrEmail,
		props.UsernameOrEmail,
		props.Password,
	)

	user := &schema.User{}

	err := row.Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot scan user: %v", err))
		return nil, err
	}

	return user, nil
}
