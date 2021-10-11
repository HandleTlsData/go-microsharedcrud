package main

import (
	"fmt"
	"net/http"
	"sharedcrud/dbmanager"
	"time"

	"github.com/gorilla/mux"
)

func entityHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "%+v\n", dbmanager.GetEntityByName(vars["name"]))

}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
	fmt.Println("incoming test request")
	dbmanager.TestConnect()
}

func main() {
	fmt.Println("Starting router and web server")
	r := mux.NewRouter()
	r.HandleFunc("/ping", testHandler)
	r.HandleFunc("/entity/{name}", entityHandler)
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
