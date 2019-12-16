package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandlerUser(w http.ResponseWriter, r *http.Request) {

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
