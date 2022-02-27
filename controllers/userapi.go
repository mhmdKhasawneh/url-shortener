package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mhmdKhasawneh/url-shortener/models"
)

type UserAPI struct {
	UserDb    *models.UserQueries
	SessionDb *models.SessionQueries
}

func (u *UserAPI) SignupUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var user models.User
	_ = json.NewDecoder(req.Body).Decode(&user)
	err := u.UserDb.InsertUser(user)
	if err != nil {
		json.NewEncoder(res).Encode(respondWithMessage("couldn't sign up new user"))
		return
	}

	sess, err := u.SessionDb.CreateNewSession(user.Email)
	if err != nil {
		json.NewEncoder(res).Encode(respondWithMessage("couldn't create new session"))
		return
	}

	json.NewEncoder(res).Encode(map[string]string{
		"token": sess.Token,
	})
}

func (u *UserAPI) LoginUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var user models.User
	_ = json.NewDecoder(req.Body).Decode(&user)

	flag, err := u.UserDb.UserExists(user.Email)

	if err != nil {
		json.NewEncoder(res).Encode(respondWithMessage("server side error"))
		return
	}

	if !flag {
		json.NewEncoder(res).Encode(respondWithMessage("email or password is incorrect"))
		return
	}

	matches, err := u.UserDb.PasswordMatches(user.Password, user.Email)

	if err != nil {
		json.NewEncoder(res).Encode(respondWithMessage("server side error"))
		return
	}

	if !matches {
		json.NewEncoder(res).Encode(respondWithMessage("email or password is incorrect"))
		return
	}

	sess, err := u.SessionDb.CreateNewSession(user.Email)

	if err != nil {
		json.NewEncoder(res).Encode(respondWithMessage("couldn't create new session"))
		return
	}

	json.NewEncoder(res).Encode(map[string]string{
		"token": sess.Token,
	})
}

func (u *UserAPI) LogoutUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var sess models.Session
	json.NewDecoder(req.Body).Decode(&sess)

	err := u.SessionDb.DeleteSession(sess.Token)

	if err != nil {
		json.NewEncoder(res).Encode(respondWithMessage("server side error"))
	}

	json.NewEncoder(res).Encode(respondWithMessage("success"))

}

func respondWithMessage(s string) map[string]string {
	return map[string]string{
		"message": s,
	}
}