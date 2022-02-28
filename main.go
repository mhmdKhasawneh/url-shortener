package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/mhmdKhasawneh/url-shortener/controllers"
	"github.com/mhmdKhasawneh/url-shortener/database"
)

func main() {
	r := mux.NewRouter()

	db, err := sql.Open("mysql", "root:mysqlmanagerkhasawneh247@tcp(localhost:3306)/shortener")
	if err != nil {
		panic(err.Error())
	}

	userDb := database.UserQueries{Db: db}
	sessionDb := database.SessionQueries{Db: db}
	urlDb := database.UrlQueries{Db: db, Sq: &sessionDb}
	lettersAndNums := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	userOps := controllers.UserAPI{UserDb: &userDb, SessionDb: &sessionDb}
	urlOps := controllers.UrlAPI{UserDb: &userDb, SessionDb: &sessionDb, UrlDB: &urlDb, Letnums: lettersAndNums}

	r.HandleFunc("/signup/", userOps.SignupUser).Methods("POST")
	r.HandleFunc("/login/", userOps.LoginUser).Methods("POST")
	r.HandleFunc("/logout/", userOps.LogoutUser).Methods("DELETE")

	r.HandleFunc("/shorten/", urlOps.Shorten).Methods("POST")
	r.HandleFunc("/{[A-Za-z0-9]{5}}", urlOps.RedirectToOriginal).Methods("GET")

	//r.Handle("/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(":8080", r))
}
