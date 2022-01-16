/*
 * Microshared-CRUD
 *
 * CRUD operation over Entities.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"sharedcrud/dbmanager"
	sw "sharedcrud/go"

	sharedserver "sharedcrud/api"
	alpharpc "sharedcrud/apirpc/alpha"
	betarpc "sharedcrud/apirpc/beta"
	restAPI "sharedcrud/cmd/sharedcrud-server"

	"google.golang.org/grpc"
)

func startAlphaGRPC() {
	s := grpc.NewServer()
	srv := &sharedserver.AlphaGRPCServer{}
	alpharpc.RegisterAlphaCRUDRPCServer(s, srv)
	conn, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := s.Serve(conn); err != nil {
		log.Fatal(err.Error())
	}

}

func startBetaGRPC() {
	s := grpc.NewServer()
	srv := &sharedserver.BetaGRPCServer{}
	betarpc.RegisterBetaCRUDRPCServer(s, srv)
	conn, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := s.Serve(conn); err != nil {
		log.Fatal(err.Error())
	}

}

//using two grpc servers linked in the same binary with same message names (GetRequest, GetReply, etc...)
//can cause issues during application runtime.
func startService(serviceName string) {
	dbmanager.CurrentAppConfig = serviceName
	go restAPI.Entry()
	switch serviceName {
	case "alpha":
		go startAlphaGRPC()
		router := sw.NewRouter(serviceName)
		log.Fatal(http.ListenAndServe(":8080", router))
	case "beta":
		go startBetaGRPC()
		router := sw.NewRouter(serviceName)
		log.Fatal(http.ListenAndServe(":8081", router))
	default:
		log.Printf("unknown microservice role specified")
		return
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Println("Specify microservice role as command-line argument")
		return
	}

	switch os.Args[1] {
	case "alpha":
		dbmanager.AlphaDBConfig = dbmanager.DBConfig{"95.214.55.115", "alpha", "A7bdLipLxw6CeEAf", "alpha", nil}
	case "beta":
		dbmanager.BetaDBConfig = dbmanager.DBConfig{"95.214.55.115", "beta", "LFmaH2X8tLiDsCDS", "beta", nil}
	default:
		log.Printf("unknown microservice role specified")
		return
	}

	startService(os.Args[1])
}
