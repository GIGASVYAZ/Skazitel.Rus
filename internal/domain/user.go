package domain

import "time"

type User struct {
	Id        int64
	Username  string
	Password  string
	CreatedAt time.Time
	IsOnline  bool
}

const UserTableSQL = `CREATE TABLE IF NOT EXISTS skazitel.users (
	id SERIAL PRIMARY KEY,
	username VARCHAR(50) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT NOW(),
	is_online BOOLEAN DEFAULT FALSE
);`
