// HTTP basic auth
// request header (base64 dari username:password)
// Authorization: Basic cjasdiaosdji23o==

// middleware -> kode yang dipanggil sebelum/sesudah http request diproses

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// mux (multiplexer) -> route
	// semua routing dilakukan lewat objek mux
	mux := http.DefaultServeMux

	// register new route
	mux.HandleFunc("/student", ActionStudent)

	// register new middleware
	var handler http.Handler = mux
	handler = MiddlewareAuth(handler)
	handler = MiddlewareAllowOnlyGet(handler)

	server := new(http.Server)
	server.Addr = ":8080"
	server.Handler = handler

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
