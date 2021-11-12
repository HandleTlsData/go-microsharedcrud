package sharedserver

import (
	"context"
	"log"
	alpharpc "sharedcrud/apirpc/alpha"
	betarpc "sharedcrud/apirpc/beta"
	"sharedcrud/dbmanager"

	"google.golang.org/grpc"
)

var AlphaConn *grpc.ClientConn
var BetaConn *grpc.ClientConn

type AlphaGRPCServer struct {
	alpharpc.UnimplementedAlphaCRUDRPCServer
}

func (s *AlphaGRPCServer) GetAlphaInformation(ctx context.Context, req *alpharpc.AlphaGetRequest) (*alpharpc.AlphaGetReply, error) {
	// sharedserver.AlphaGRPCServer = alpharpc.NewAlphaCRUDRPCClient()
	entity, err := dbmanager.GetEntityByName(&dbmanager.AlphaDBConfig, req.GetEntityName())
	if err != nil {
		log.Println(err)
	} else {
		log.Println(entity)
	}
	return &alpharpc.AlphaGetReply{ID: int64(entity.ID), Name: entity.Name, Description: entity.Description, Status: entity.Status}, err
}

func (s *AlphaGRPCServer) UpdateAlphaInformation(ctx context.Context, req *alpharpc.AlphaUpdateRequest) (*alpharpc.AlphaUpdateReply, error) {
	return nil, nil

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
	return nil, nil

}
