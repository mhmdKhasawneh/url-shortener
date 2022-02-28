package database

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/mhmdKhasawneh/url-shortener/models"

)

type SessionQueries struct {
	Db *sql.DB
}

func (s *SessionQueries) CreateNewSession(email string) (models.Session, error) {
	userId, err := s.getIdFromEmail(email)

	if err != nil {
		panic(err.Error())
	}

	tok := uuid.New().String()

	_, errs := s.Db.Query(
		"INSERT INTO shortener.sessions(user_id,token) VALUES(?,?);",
		userId, tok)

	if errs != nil {
		return models.Session{}, errs
	}

	return models.Session{UserId: userId, Token: tok}, nil
}

func (s *SessionQueries) getIdFromEmail(email string) (int, error) {
	var retrievedId int
	err := s.Db.QueryRow("SELECT id FROM shortener.users WHERE email=?;", email).Scan(&retrievedId)

	if err != nil {
		return -1, err
	}

	return retrievedId, nil
}

func (s *SessionQueries) getIdFromToken(token string) (int, error) {
	var retrievedId int
	err := s.Db.QueryRow("SELECT user_id id FROM shortener.sessions WHERE token=?;", token).Scan(&retrievedId)

	if err != nil {
		return -1, err
	}

	return retrievedId, nil
}

func (s *SessionQueries) DeleteSession(token string) error {
	_, err := s.Db.Query("DELETE FROM shortener.sessions WHERE token=?;", token)

	if err != nil {
		return err
	}

	return nil

}
