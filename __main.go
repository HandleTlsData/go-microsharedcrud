package main

import (
	"fmt"
	"net/http"
	"os"
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
	if len(os.Args) < 2 {
		fmt.Println("specify microservice role as command-line argument")
		return
	}
	fmt.Println("got argument: " + os.Args[1])
	switch os.Args[1] {
	case "alpha":
		fmt.Println("alpha microservice")
		currentDBConfig = dbmanager.DBConfig{"95.214.55.115", "alpha", "A7bdLipLxw6CeEAf", "alpha", nil}
	case "beta":
		fmt.Println("beta microservice")
		currentDBConfig = dbmanager.DBConfig{"95.214.55.115", "beta", "LFmaH2X8tLiDsCDS", "beta", nil}
	default:
		fmt.Printf("unknown microservice role specified")
	}
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
