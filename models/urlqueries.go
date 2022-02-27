package models

import (
	"database/sql"
)

type UrlQueries struct {
	Db *sql.DB
	Sq *SessionQueries
}

func (u *UrlQueries) InsertUrl(fullUrl string, shortenedUrl string, token string) error {
	userId, err := u.Sq.getIdFromToken(token)

	if err != nil {
		return err
	}

	_, err = u.Db.Query("INSERT INTO shortener.urls(full_url, shortened_url, generated_by) VALUES(?,?,?)",
	                     fullUrl, shortenedUrl, userId)

	if err!=nil{
		return err
	}

	return nil
}

func (u *UrlQueries) ShortExists(url string) bool {
	var exists int
	_ = u.Db.QueryRow("SELECT EXISTS(SELECT * FROM shortener.urls WHERE shortened_url='?');", url).Scan(&exists)

	return exists != 0
}

func (u *UrlQueries) GetFullUrlFromShortened(shortUrl string) string {
	var fullUrl string
	_ = u.Db.QueryRow("SELECT full_url FROM shortener.urls WHERE short_url='?'",shortUrl).Scan(&fullUrl)

	return fullUrl
}