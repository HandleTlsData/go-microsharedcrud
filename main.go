package main

import (
	"fmt"
	"net/http"
	"sharedcrud/dbmanager"
	"time"

	"github.com/gorilla/mux"
)

var currentDBConfig dbmanager.DBConfig

func entityHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entity, err := dbmanager.GetEntityByName(&currentDBConfig, vars["name"])
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	} else {
		fmt.Fprintf(w, "%+v\n", entity)
	}

}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("incoming test request")
	err := dbmanager.Connect(&currentDBConfig)
	if err != nil {
		fmt.Fprint(w, "database connection failed: ", err)
	} else {
		fmt.Fprint(w, "pong")
	}
	dbmanager.Disconnect(&currentDBConfig)
}

func main() {
	currentDBConfig = dbmanager.DBConfig{"95.214.55.115", "testdb", "2pRSTXAMBh5wLMtF", "testdb", nil}
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
