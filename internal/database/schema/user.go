package schema

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	// Password  string `json:"-"` // don't expose password in API
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

var UserTableSchemaCreate = `
CREATE TABLE IF NOT EXISTS users (
	id TEXT PRIMARY KEY,
	username TEXT NOT NULL,
	email TEXT NOT NULL,
	password TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

var UserTableSchemaDrop = `
DROP TABLE IF EXISTS users;
`

type UserPropsSignup struct {
	Id       string
	Username string
	Email    string
	Password string
}

var UserSchemaSignup = `
INSERT INTO users (
	id, username, email, password
) VALUES (
	?, ?, ?, ?
);
`

type UserPropsSignin struct {
	UsernameOrEmail string
	Password        string
}

var UserSchemaSignin = `
SELECT id, username, email, created_at, updated_at
FROM users
WHERE (username = ? OR email = ?) AND password = ?;
`

var UserSchemaGetByID = `
SELECT id, username, email, created_at, updated_at
FROM users
WHERE id = ?;
`

type UserPropsGetAll struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

var UserSchemaGetAll = `
SELECT id, username, email, created_at, updated_at
FROM users;
`

var UserSchemaUpdate = `
UPDATE users
SET username = ?, email = ?, password = ?
WHERE id = ?;
`

var UserSchemaDelete = `
DELETE FROM users
WHERE id = ?;
`
