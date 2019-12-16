package main

import (
	"fmt"
	"net/http"
)

type response struct {
	Content string `json:"content"`
}

func BasicAuth(handler http.HandlerFunc, realm string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		user, pass, ok := r.BasicAuth()

		if ok && validateUser(user, pass) {
			handler(w, r)
		} else {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			return
		}
	}
}

func writeBadRequest(w http.ResponseWriter, s string) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Print(w, response{s})
}
