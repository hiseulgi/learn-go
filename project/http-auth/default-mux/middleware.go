package main

import "net/http"

const USERNAME = "sugab"
const PASSWORD = "secret"

// basic auth: 3 info
// - username
// - password
// - valid tidaknya basic auth pada request
func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("something went wrong"))
			return
		}

		// pengecekan username dan password dengan data pada server
		isValid := (username == USERNAME) && (password == PASSWORD)
		if !isValid {
			w.Write([]byte("wrong username/password"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func MiddlewareAllowOnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Write([]byte("Only GET is allowed"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
