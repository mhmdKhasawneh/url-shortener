package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/mhmdKhasawneh/url-shortener/database"
)

type UrlAPI struct {
	UserDb    *database.UserQueries
	SessionDb *database.SessionQueries
	UrlDB     *database.UrlQueries
	Letnums   []rune
}

type urlInfo struct {
	Token    string `json:"token"`
	Full_url string `json:"full_url"`
}

func (u *UrlAPI) Shorten(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var info urlInfo
	_ = json.NewDecoder(req.Body).Decode(&info)
	shortenedUrl := u.generateRandomString()

	for u.UrlDB.ShortExists(shortenedUrl) {
		shortenedUrl = u.generateRandomString()
	}

	u.UrlDB.InsertUrl(info.Full_url, shortenedUrl, info.Token)

	_ = json.NewEncoder(res).Encode(map[string]string{
		"success": "OK",
	})
}

func (u *UrlAPI) RedirectToOriginal(res http.ResponseWriter, req *http.Request) {
	shortUrl := req.URL.Path[1:]
	if u.UrlDB.ShortExists(shortUrl) {
		fullUrl := u.UrlDB.GetFullUrlFromShortened(shortUrl)
		http.Redirect(res, req, fullUrl, http.StatusMovedPermanently)
		_ = json.NewEncoder(res).Encode(map[string]string{
			"message": "success",
		})
		return
	}
	http.Redirect(res, req, "", 404)
}

func (u *UrlAPI) generateRandomString() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 5)
	for i := range b {
		b[i] = u.Letnums[rand.Intn(len(u.Letnums))]
	}
	return string(b)
}
