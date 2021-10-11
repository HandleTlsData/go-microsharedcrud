package main

import (
	"fmt"
	"net/http"
	"sharedcrud/dbmanager"
	"time"

	"github.com/gorilla/mux"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
	fmt.Println("incoming test request")
	dbmanager.Connect()
}

func main() {
	fmt.Println("Starting router and web server")
	r := mux.NewRouter()
	r.HandleFunc("/ping", testHandler)
	r.Use(mux.CORSMethodMiddleware(r))
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	panic(srv.ListenAndServe())
}
