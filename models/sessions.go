package models

type Session struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Token  string `json:"token"`
}