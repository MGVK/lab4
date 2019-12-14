package main

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"net/http"
)

type loginData struct {
	Token string `json:"token"`
}

type response struct {
	Content string `json:"content"`
}

type request struct {
	Token string `json:"token"`
	Data  string `json:"data"`
}

//https://stackoverflow.com/questions/21936332/idiomatic-way-of-requiring-http-basic-auth-in-go
func BasicAuth(handler http.HandlerFunc, username, password, realm string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			return
		}

		handler(w, r)
	}
}

func HandlerNews(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		{
			news := GetAllNewsFromDB()
			j, err := json.Marshal(news)
			if err != nil {
				fmt.Print(err)
			}
			fmt.Print(w, "news...:")
			_, _ = fmt.Fprintf(w, string(j))
			return
		}
	case "POST":
		{
			var news news
			err := json.NewDecoder(r.Body).Decode(&news)
			if err != nil {
				fmt.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if validateNews(news) {
				updateNews(news)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Print(w, response{"News is invalid"})
				return
			}

		}
	case "PUT":
		{
			var news news
			err := json.NewDecoder(r.Body).Decode(&news)
			if err != nil {
				fmt.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if validateNews(news) {
				addNewsToDB(news)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Print(w, response{"News is invalid"})
				return
			}
		}

	}

}

func writeBad(w http.ResponseWriter, s string) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Print(w, response{s})
}
