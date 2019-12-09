package src

import (
	"fmt"
	"html"
	"net/http"
)

func main() {

	DBConn()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/news", BasicAuth(HandlerNews, "admin", "123456", "Please enter your username and password for this site"))
	//http.HandleFunc("/user", BasicAuth(HandlerPutUser, "admin", "123456", "Please enter your username and password for this site"))
	//http.HandleFunc("/login", BasicAuth(HandlerLogin, "admin", "123456", "Please enter your username and password for this site"))

	fmt.Print(http.ListenAndServe(":8080", nil))

}
