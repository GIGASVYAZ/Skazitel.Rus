package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
)

const DEF = `CREATE TABLE skazitel.users (
	id SERIAL PRIMARY KEY,
	username VARCHAR(50) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT NOW(),
	is_online BOOLEAN DEFAULT FALSE
);`
const DATABASE_URL = "postgres://postgres:mysecretpassword@localhost:5432/postgres"

type User struct {
	Id        int64
	Username  string
	Password  string
	CreatedAt time.Time
	IsOnline  bool
}

func CreateUser(Username string, Password string) error {
	var err error
	conn, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		return err
	}

	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "insert into skazitel.users (username, password) values($1, $2)", Username, Password)
	if err != nil {
		return err
	}

	return nil
}

func UserAuthenticate(Username string, Password string) (bool, error) {
	var err error
	conn, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		return false, err
	}

	defer conn.Close(context.Background())

	var passwordUser string
	err = conn.QueryRow(
		context.Background(),
		"SELECT password FROM skazitel.users WHERE username = $1", Username).Scan(&passwordUser)

	if Password == passwordUser {
		return true, nil
	}

	return false, err
}

func UpdateUserStatus(Username string, IsOnline bool) error {
	var err error
	conn, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		return err
	}

	defer conn.Close(context.Background())

	tag, err := conn.Exec(
		context.Background(),
		"update skazitel.users set is_online = $2 where username = $1", Username, IsOnline)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("Ошибка")
	}

	return nil
}
