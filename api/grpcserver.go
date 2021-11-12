package sharedserver

import (
	"context"
	"sharedcrud/apirpc"
)

type GRPCServer struct {
	apirpc.UnimplementedCRUDIntercommunicationServer
}

func (s *GRPCServer) GetAlphaInformation(ctx context.Context, req *apirpc.GetRequest) (*apirpc.GetReply, error) {
	// entity, err := dbmanager.GetEntityByName(&dbmanager.AlphaDBConfig, req.EntityName)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(entity)
	// }
	// return &apirpc.GetReply{ID: int64(entity.ID), Name: entity.Name, Description: entity.Description, Status: entity.Status}, err
	return nil, nil
}

func (s *GRPCServer) GetBetaInformation(ctx context.Context, req *apirpc.GetRequest) (*apirpc.GetReply, error) {

	// entity, err := dbmanager.GetEntityByName(&dbmanager.AlphaDBConfig, req.EntityName)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(entity)
	// }
	// return &apirpc.GetReply{ID: int64(entity.ID), Name: entity.Name, Description: entity.Description, Status: entity.Status}, err
	return nil, nil
}

func (s *GRPCServer) UpdateBetaInformation(ctx context.Context, req *apirpc.UpdateRequest) (*apirpc.UpdateReply, error) {
	return nil, nil

}
func (s *GRPCServer) UpdateAlphaInformation(ctx context.Context, req *apirpc.UpdateRequest) (*apirpc.UpdateReply, error) {
	return nil, nil

}
