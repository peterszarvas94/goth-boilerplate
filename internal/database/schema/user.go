package schema

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"` // don't return password in JSON
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

var CreateUserTableSchema = `
CREATE TABLE IF NOT EXISTS users (
	id VARCHAR(36) PRIMARY KEY,
	username VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
`

var DropUserTableSchema = `
DROP TABLE IF EXISTS users;
`

var CreateUserSchema = `
INSERT INTO users (
	id, username, email, password
) VALUES (
	?, ?, ?, ?
);
`

var GetUserSchema = `
SELECT id, username, email, password, created_at, updated_at
FROM users
WHERE id = ?;
`

var GetUsersSchema = `
SELECT id, username, email, password, created_at, updated_at
FROM users;
`

var UpdateUserSchema = `
UPDATE users
SET username = ?, email = ?, password = ?
WHERE id = ?;
`

var DeleteUserSchema = `
DELETE FROM users
WHERE id = ?;
`
