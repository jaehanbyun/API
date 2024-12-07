package model

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type DBHandler interface {
	GetUsers() ([]User, error)
	GetUserByID(id int) (User, error)
	CreateUser(user User) error
	UpdateUser(user User) error
	DeleteUser(id int) error
	Close()
	Init() error
}

type postgresHandler struct {
	db *sql.DB
}

func newPostgresHandler() DBHandler {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"postgres-service", 5432, "postgres", "postgres", "user",
	)

	database, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = database.Ping()
	if err != nil {
		panic(err)
	}

	return &postgresHandler{database}
}

func (h *postgresHandler) Init() error {
	_, err := h.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL
		);
	`)
	return err
}

func (h *postgresHandler) Close() {
	h.db.Close()
}

func (h *postgresHandler) GetUsers() ([]User, error) {
	rows, err := h.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (h *postgresHandler) GetUserByID(id int) (User, error) {
	var user User
	err := h.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}

func (h *postgresHandler) CreateUser(user User) error {
	_, err := h.db.Exec("INSERT INTO users (username, email) VALUES ($1, $2)", user.Username, user.Email)
	return err
}

func (h *postgresHandler) UpdateUser(user User) error {
	_, err := h.db.Exec("UPDATE users SET username = $1, email = $2 WHERE id = $3", user.Username, user.Email, user.ID)
	return err
}

func (h *postgresHandler) DeleteUser(id int) error {
	_, err := h.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func NewDBHandler() DBHandler {
	return newPostgresHandler()
}
