package sharedserver

import (
	"context"
	"log"
	alpharpc "sharedcrud/apirpc/alpha"
	betarpc "sharedcrud/apirpc/beta"
	"sharedcrud/dbmanager"
	gdbmanager "sharedcrud/gormdb"
	"sharedcrud/models"
	"strconv"

	"google.golang.org/grpc"
)

var AlphaConn *grpc.ClientConn
var BetaConn *grpc.ClientConn

type AlphaGRPCServer struct {
	alpharpc.UnimplementedAlphaCRUDRPCServer
}

func (s *AlphaGRPCServer) GetAlphaInformation(ctx context.Context, req *alpharpc.AlphaGetRequest) (*alpharpc.AlphaGetReply, error) {
	entity, err := gdbmanager.GetEntityByName(req.GetEntityName())
	if err != nil {
		log.Println(err)
	} else {
		log.Println(entity)
	}
	return &alpharpc.AlphaGetReply{ID: entity.ID, Name: entity.Name, Description: entity.Description, Status: entity.Status}, err
}

func (s *AlphaGRPCServer) UpdateAlphaInformation(ctx context.Context, req *alpharpc.AlphaUpdateRequest) (*alpharpc.AlphaUpdateReply, error) {
	i, err := strconv.Atoi(req.ID)
	if err != nil {
		// handle error
		log.Println(err.Error())
		return &alpharpc.AlphaUpdateReply{Status: "500"}, err
	}

	err = gdbmanager.StoreEntity(models.Entity{ID: int64(i), Name: req.Name, Description: req.Description, Status: req.Status})
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

	entity, err := gdbmanager.GetEntityByName(req.GetEntityName())
	if err != nil {
		log.Println(err)
	} else {
		log.Println(entity)
	}
	return &betarpc.BetaGetReply{ID: entity.ID, Name: entity.Name, Description: entity.Description, Status: entity.Status}, err
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
