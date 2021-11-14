package sharedserver

import (
	"context"
	"log"
	alpharpc "sharedcrud/apirpc/alpha"
	betarpc "sharedcrud/apirpc/beta"
	"sharedcrud/dbmanager"
	"strconv"

	"google.golang.org/grpc"
)

var AlphaConn *grpc.ClientConn
var BetaConn *grpc.ClientConn

type AlphaGRPCServer struct {
	alpharpc.UnimplementedAlphaCRUDRPCServer
}

func (s *AlphaGRPCServer) GetAlphaInformation(ctx context.Context, req *alpharpc.AlphaGetRequest) (*alpharpc.AlphaGetReply, error) {
	entity, err := dbmanager.GetEntityByID(&dbmanager.AlphaDBConfig, req.GetEntityID())
	if err != nil {
		log.Println(err)
	} else {
		log.Println(entity)
	}
	return &alpharpc.AlphaGetReply{ID: int64(entity.ID), Name: entity.Name, Description: entity.Description, Status: entity.Status}, err
}

func (s *AlphaGRPCServer) UpdateAlphaInformation(ctx context.Context, req *alpharpc.AlphaUpdateRequest) (*alpharpc.AlphaUpdateReply, error) {
	i, err := strconv.Atoi(req.ID)
	if err != nil {
		// handle error
		log.Println(err.Error())
		return &alpharpc.AlphaUpdateReply{Status: "500"}, err
	}

	err = dbmanager.StoreEntityAlpha(&dbmanager.AlphaDBConfig, dbmanager.DBEntity{ID: i, Name: req.Name, Description: req.Description, Status: req.Status})
	if err != nil {
		// handle error
		log.Println(err.Error())
		return &alpharpc.AlphaUpdateReply{Status: "500"}, err
	}

	return &alpharpc.AlphaUpdateReply{Status: "200"}, nil

}

type BetaGRPCServer struct {
	betarpc.UnimplementedBetaCRUDRPCServer
}

func (s *BetaGRPCServer) GetBetaInformation(ctx context.Context, req *betarpc.BetaGetRequest) (*betarpc.BetaGetReply, error) {

	entity, err := dbmanager.GetEntityByID(&dbmanager.BetaDBConfig, req.GetEntityID())
	if err != nil {
		log.Println(err)
	} else {
		log.Println(entity)
	}
	return &betarpc.BetaGetReply{ID: int64(entity.ID), Name: entity.Name, Description: entity.Description, Status: entity.Status}, err
}

func (s *BetaGRPCServer) UpdateBetaInformation(ctx context.Context, req *betarpc.BetaUpdateRequest) (*betarpc.BetaUpdateReply, error) {
	i, err := strconv.Atoi(req.ID)
	if err != nil {
		// handle error
		log.Println(err.Error())
		return &betarpc.BetaUpdateReply{Status: "500"}, err
	}

	err = dbmanager.StoreEntityBeta(&dbmanager.BetaDBConfig, dbmanager.DBEntity{ID: i, Name: req.Name,
		Description: req.Description, Status: req.Status})
	if err != nil {
		// handle error
		log.Println(err.Error())
		return &betarpc.BetaUpdateReply{Status: "500"}, err
	}

	return &betarpc.BetaUpdateReply{Status: "200"}, nil
}
