/*
 * Microshared-CRUD
 *
 * CRUD operation over Entities.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"sharedcrud/dbmanager"

	"github.com/gorilla/mux"
)

func EntityHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entity, err := dbmanager.GetEntityByName(&dbmanager.CurrentDBConfig, vars["entityName"])
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	} else {
		fmt.Fprintf(w, "%+v\n", entity)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func EntityStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
