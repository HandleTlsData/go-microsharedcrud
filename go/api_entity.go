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
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	alpharpc "sharedcrud/apirpc/alpha"
	betarpc "sharedcrud/apirpc/beta"
	"sharedcrud/dbmanager"
	"strconv"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func EntityHandlerAlpha(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, err := strconv.Atoi(vars["entID"])
	if err != nil {
		// handle error
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	entity, err := dbmanager.GetEntityByID(&dbmanager.AlphaDBConfig, int64(i))

	if err != nil {
		log.Println(err.Error())
	}

	conn, err := grpc.Dial(":8001", grpc.WithInsecure())
	if err != nil {
		log.Println(err.Error())
	}

	//asking from Beta microservice to give entity from Beta's own db
	bRPC := betarpc.NewBetaCRUDRPCClient(conn)
	betaEntity, grpcerr := bRPC.GetBetaInformation(context.Background(), &betarpc.BetaGetRequest{EntityID: int64(entity.ID)})

	if grpcerr != nil || betaEntity == nil {
		log.Println(grpcerr.Error())
		fmt.Fprintf(w, "%s", grpcerr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Beta always contains very last Description version since Beta is only responsible for Description modifications
	entity.Description = betaEntity.GetDescription()

	fmt.Fprintf(w, "%+v\n", entity)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func EntityHandlerBeta(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, err := strconv.Atoi(vars["entID"])
	if err != nil {
		// handle error
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Println(err.Error())
	}

	//asking Alpha first since Alpha is owner of Name field, which is our key to search records
	aRPC := alpharpc.NewAlphaCRUDRPCClient(conn)
	alphaEntity, grpcerr := aRPC.GetAlphaInformation(context.Background(), &alpharpc.AlphaGetRequest{EntityID: int64(i)})

	if grpcerr != nil || alphaEntity == nil {
		log.Println(grpcerr.Error())
		fmt.Fprintf(w, "%s", grpcerr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	entity, err := dbmanager.GetEntityByID(&dbmanager.BetaDBConfig, alphaEntity.GetID())
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	// Beta always contains very last Description version since Beta is only responsible for Description modifications
	entity.Name = alphaEntity.GetName()

	fmt.Fprintf(w, "%+v\n", entity)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func EntityStoreAlpha(w http.ResponseWriter, r *http.Request) {
	ent := dbmanager.DBEntity{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ent); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println(ent)

	err := dbmanager.StoreEntityAlpha(&dbmanager.AlphaDBConfig, ent)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	conn, err := grpc.Dial(":8001", grpc.WithInsecure())
	if err != nil {
		log.Println(err.Error())
	}

	bRPC := betarpc.NewBetaCRUDRPCClient(conn)
	betaReply, grpcerr := bRPC.UpdateBetaInformation(context.Background(), &betarpc.BetaUpdateRequest{ID: strconv.Itoa(ent.ID),
		Name: ent.Name, Description: ent.Description, Status: ent.Status})

	if grpcerr != nil || betaReply == nil {
		log.Println(grpcerr.Error())
		fmt.Fprintf(w, "%s", grpcerr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func EntityStoreBeta(w http.ResponseWriter, r *http.Request) {
	ent := dbmanager.DBEntity{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ent); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := dbmanager.StoreEntityBeta(&dbmanager.BetaDBConfig, ent)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Println(err.Error())
	}

	aRPC := alpharpc.NewAlphaCRUDRPCClient(conn)
	alphaReply, grpcerr := aRPC.UpdateAlphaInformation(context.Background(), &alpharpc.AlphaUpdateRequest{ID: string(ent.ID),
		Name: ent.Name, Description: ent.Description, Status: ent.Status})

	if grpcerr != nil || alphaReply == nil {
		log.Println(grpcerr.Error())
		fmt.Fprintf(w, "%s", grpcerr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func EntityDeleteAlpha(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, err := strconv.Atoi(vars["entID"])
	if err != nil {
		// handle error
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = dbmanager.DeleteEntity(&dbmanager.AlphaDBConfig, i)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func EntityDeleteBeta(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, err := strconv.Atoi(vars["entID"])
	if err != nil {
		// handle error
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = dbmanager.DeleteEntity(&dbmanager.BetaDBConfig, i)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
