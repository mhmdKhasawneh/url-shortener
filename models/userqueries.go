package models

import (
	"crypto/md5"
	"database/sql"
)

type UserQueries struct {
	Db *sql.DB
}

func (u *UserQueries) InsertUser(user User) error {
	_, err := u.Db.Query(
		"INSERT INTO shortener.users(name,email,password,user_type) VALUES(?,?,?,?);",
		user.Name, user.Email, md5.Sum([]byte(user.Password)))

	if err != nil {
		return err
	}

	return nil
}

func (u *UserQueries) UserExists(email string) (bool, error) {
	var check string
	err := u.Db.QueryRow("SELECT email FROM shortener.users WHERE email='?';", email).Scan(&check)

	if err != nil {
		return false, err
	}

	if check == "" {
		return false, nil
	}

	return true, nil
}

func (u *UserQueries) PasswordMatches(password string, email string) (bool, error) {
	hashedPassword := md5.Sum([]byte(password))
	var hashedPasswordFromQuery [16]byte

	err := u.Db.QueryRow("SELECT password FROM shortener.users WHERE email='?';", email).Scan(&hashedPasswordFromQuery)

	if err != nil {
		return false, err
	}

	if hashedPassword != hashedPasswordFromQuery {
		return false, nil
	}

	return true, nil
}