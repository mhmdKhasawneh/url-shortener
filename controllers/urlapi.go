package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/mhmdKhasawneh/url-shortener/models"
)

type UrlAPI struct {
	UserDb    *models.UserQueries
	SessionDb *models.SessionQueries
	UrlDB     *models.UrlQueries
	Letnums   []rune
}

func (u *UrlAPI) Shorten(res http.ResponseWriter, req *http.Request) {
	token := u.getToken(req)
	fullUrl := u.getUrl(req)
	shortenedUrl := u.generateRandomString()

	for u.UrlDB.ShortExists(shortenedUrl) {
		shortenedUrl = u.generateRandomString()
	}

	u.UrlDB.InsertUrl(fullUrl, shortenedUrl, token)
}

func (u *UrlAPI) RedirectToOriginal(res http.ResponseWriter, req *http.Request) {
	shortUrl := req.URL.Path[1:]
	if u.UrlDB.ShortExists(shortUrl) {
		fullUrl := u.UrlDB.GetFullUrlFromShortened(shortUrl)
		http.Redirect(res, req, fullUrl, 200)
		return
	}
	http.Redirect(res, req, "facebook.com", 404)
}

func (u *UrlAPI) generateRandomString() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 5)
	for i := range b {
		b[i] = u.Letnums[rand.Intn(len(u.Letnums))]
	}
	return string(b)
}

func (u *UrlAPI) getToken(req *http.Request) string {
	var sess models.Session
	_ = json.NewDecoder(req.Body).Decode(&sess)
	return sess.Token
}

func (u *UrlAPI) getUrl(req *http.Request) string {
	var full models.Url
	_ = json.NewDecoder(req.Body).Decode(&full)
	return full.Full_url
}
