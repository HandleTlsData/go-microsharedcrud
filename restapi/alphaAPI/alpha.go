package alphaAPI

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"sharedcrud/models"

	betarpc "sharedcrud/apirpc/beta"
	gdbmanager "sharedcrud/gormdb"
	APIEntity "sharedcrud/restapi/operations/entity"

	"google.golang.org/grpc"
)

func EntityDelete(params APIEntity.EntityDeleteParams) error {
	entityName, err := url.QueryUnescape(params.EntityName)
	if err != nil {
		entityName = params.EntityName
	}
	entity, err := gdbmanager.GetEntityByName(entityName)
	if err != nil {
		return err
	}
	err = gdbmanager.DeleteEntityByName(entity.Name)
	if err != nil {
		return err
	}
	return nil
}

func EntityGet(params APIEntity.EntityGetParams) ([]*models.Entity, error) {
	entityName, err := url.QueryUnescape(params.EntityName)
	if err != nil {
		log.Println(err.Error())
		entityName = params.EntityName
	}
	log.Println(entityName)

	entity, err := gdbmanager.GetEntityByName(entityName)
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(":8001", grpc.WithInsecure())
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	//asking Beta microservice for entity from its own db
	bRPC := betarpc.NewBetaCRUDRPCClient(conn)
	betaEntity, grpcerr := bRPC.GetBetaInformation(context.Background(), &betarpc.BetaGetRequest{EntityName: entity.Name})

	if grpcerr != nil || betaEntity == nil {
		return nil, fmt.Errorf(grpc.ErrorDesc(grpcerr))
	}

	// Beta always contains very last Description since it is only responsible for Description modifications
	entity.Description = betaEntity.GetDescription()

	var responsePayload []*models.Entity
	responsePayload = append(responsePayload, &entity)
	return responsePayload, nil
}

func EntityStore(params APIEntity.EntityStoreParams) error {
	var err error
	newEntity := *params.Body

	conn, err := grpc.Dial(":8001", grpc.WithInsecure())
	if err != nil {
		log.Println(err.Error())
	}

	//asking Beta microservice for entity from its own db
	bRPC := betarpc.NewBetaCRUDRPCClient(conn)

	//means user want to update existing entity. Primary ID stored in alpha db
	if newEntity.ID > 0 {
		oldEntity, err := gdbmanager.GetEntityByID(newEntity.ID)
		if err != nil {
			return err
		}

		betaDBEntity, grpcerr := bRPC.GetBetaInformation(context.Background(), &betarpc.BetaGetRequest{EntityName: oldEntity.Name})

		if grpcerr != nil || betaDBEntity == nil {
			return fmt.Errorf(grpc.ErrorDesc(grpcerr))
		}

		//we can't modify description here
		newEntity.Description = betaDBEntity.Description

		err = gdbmanager.UpdateEntity(newEntity, oldEntity.ID)
		if err != nil {
			return err
		}
		betaResponse, grpcerr := bRPC.UpdateBetaInformation(context.Background(), &betarpc.BetaUpdateRequest{ID: 0,
			Name: newEntity.Name, Status: newEntity.Status, Description: newEntity.Description})

		if grpcerr != nil || betaResponse == nil {
			return fmt.Errorf(grpc.ErrorDesc(grpcerr))
		}

	} else {
		err = gdbmanager.StoreEntity(*params.Body)
		if err != nil {
			return err
		}
		betaEntity, grpcerr := bRPC.UpdateBetaInformation(context.Background(), &betarpc.BetaUpdateRequest{ID: 0,
			Name: newEntity.Name, Status: newEntity.Status, Description: newEntity.Description})

		if grpcerr != nil || betaEntity == nil {
			return fmt.Errorf(grpc.ErrorDesc(grpcerr))
		}
	}

	return nil
}
