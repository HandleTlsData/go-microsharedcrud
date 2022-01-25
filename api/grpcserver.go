package sharedserver

import (
	"context"
	"fmt"
	"log"
	"net"
	alpharpc "sharedcrud/apirpc/alpha"
	betarpc "sharedcrud/apirpc/beta"
	gdbmanager "sharedcrud/gormdb"
	"sharedcrud/models"

	"google.golang.org/grpc"
)

func StartAlphaGRPC() {
	s := grpc.NewServer()
	srv := &AlphaGRPCServer{}
	alpharpc.RegisterAlphaCRUDRPCServer(s, srv)
	conn, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := s.Serve(conn); err != nil {
		log.Fatal(err.Error())
	}

}

func StartBetaGRPC() {
	s := grpc.NewServer()
	srv := &BetaGRPCServer{}
	betarpc.RegisterBetaCRUDRPCServer(s, srv)
	conn, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := s.Serve(conn); err != nil {
		log.Fatal(err.Error())
	}

}

var AlphaConn *grpc.ClientConn
var BetaConn *grpc.ClientConn

type AlphaGRPCServer struct {
	alpharpc.UnimplementedAlphaCRUDRPCServer
}

func (s *AlphaGRPCServer) GetAlphaInformationByID(ctx context.Context, req *alpharpc.AlphaGetByIDRequest) (*alpharpc.AlphaGetReply, error) {
	entity, err := gdbmanager.GetEntityByID(req.GetID())
	if err != nil {
		log.Println(err)
	} else {
		log.Println(entity)
	}
	return &alpharpc.AlphaGetReply{ID: entity.ID, Name: entity.Name, Description: entity.Description, Status: entity.Status}, err
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

	if req.GetID() > 0 {
		entity, err := gdbmanager.GetEntityByID(req.GetID())

		if err == nil {
			err = gdbmanager.UpdateEntity(models.Entity{Name: req.GetName(), Status: req.GetStatus(), Description: req.GetDescription()}, entity.ID)
			if err != nil {
				log.Println(err.Error())
				return &alpharpc.AlphaUpdateReply{Status: "500"}, err
			}
			return &alpharpc.AlphaUpdateReply{Status: "200"}, nil
		}
		return &alpharpc.AlphaUpdateReply{Status: "500"}, fmt.Errorf("Incorrect entity given. No user was found with specified ID")
	}

	entity, err := gdbmanager.GetEntityByName(req.GetName())

	if err != nil {
		if err.Error() == gdbmanager.StrNoRecords {
			err = gdbmanager.StoreEntity(models.Entity{Name: req.GetName(), Status: req.GetStatus(), Description: req.GetDescription()})
			if err != nil {
				log.Println(err.Error())
				return &alpharpc.AlphaUpdateReply{Status: "500"}, err
			}
			return &alpharpc.AlphaUpdateReply{Status: "200"}, nil
		}
		return &alpharpc.AlphaUpdateReply{Status: "500"}, fmt.Errorf("%s: %s", err.Error(), entity.Name)
	}
	return &alpharpc.AlphaUpdateReply{Status: "500"}, fmt.Errorf("Entity with such name already exists. Specify entity ID for update")
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
	entity, err := gdbmanager.GetEntityByName(req.GetName())

	if err == nil {
		err = gdbmanager.UpdateEntity(models.Entity{Name: req.GetName(), Status: entity.Status, Description: req.GetDescription()}, entity.ID)
		if err != nil {
			log.Println(err.Error())
			return &betarpc.BetaUpdateReply{Status: "500"}, err
		}
		return &betarpc.BetaUpdateReply{Status: "200"}, nil
	} else {
		if err.Error() == gdbmanager.StrNoRecords {
			err = gdbmanager.StoreEntity(models.Entity{ID: 0, Name: req.GetName(), Status: req.GetStatus(), Description: req.GetDescription()})
			if err != nil {
				log.Println(err.Error())
				return &betarpc.BetaUpdateReply{Status: "500"}, err
			}
			return &betarpc.BetaUpdateReply{Status: "200"}, nil
		}
	}
	return &betarpc.BetaUpdateReply{Status: "500"}, err
}

func (s *AlphaGRPCServer) DeleteAlphaInformation(ctx context.Context, req *alpharpc.AlphaGetRequest) (*alpharpc.AlphaUpdateReply, error) {
	err := gdbmanager.DeleteEntityByName(req.GetEntityName())
	if err != nil {
		return &alpharpc.AlphaUpdateReply{Status: "500"}, err
	}
	return &alpharpc.AlphaUpdateReply{Status: "200"}, nil
}
