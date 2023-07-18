// pembuatan custom multiplexer
// -> agar mempermudah manajemen middleware
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// mux (multiplexer) -> route
	// semua routing dilakukan lewat objek mux
	mux := new(CustomMux)

	// register new route
	mux.HandleFunc("/student", ActionStudent)

	// register new middleware
	mux.RegisterMiddleware(MiddlewareAuth)
	mux.RegisterMiddleware(MiddlewareAllowOnlyGet)

	server := new(http.Server)
	server.Addr = ":8080"
	server.Handler = mux

	log.Println("run server :8080")
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, o any) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
