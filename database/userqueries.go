package database

import (
	"crypto/md5"
	"database/sql"
	"github.com/mhmdKhasawneh/url-shortener/models"

)

type UserQueries struct {
	Db *sql.DB
}

func (u *UserQueries) InsertUser(user models.User) error {
	_, err := u.Db.Query(
		"INSERT INTO shortener.users(name,email,password) VALUES(?,?,?);",
		user.Name, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserQueries) UserExists(email string) (bool, error) {
	var check string
	err := u.Db.QueryRow("SELECT email FROM shortener.users WHERE email=?;", email).Scan(&check)

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

	err := u.Db.QueryRow("SELECT password FROM shortener.users WHERE email=?;", email).Scan(&hashedPasswordFromQuery)

	if err != nil {
		return false, err
	}

	if hashedPassword != hashedPasswordFromQuery {
		return false, nil
	}

	return true, nil
}