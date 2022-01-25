package betaAPI

import (
	"context"
	"fmt"
	"log"
	"net/url"
	alpharpc "sharedcrud/apirpc/alpha"
	gdbmanager "sharedcrud/gormdb"
	"sharedcrud/models"
	APIEntity "sharedcrud/restapi/operations/entity"

	"google.golang.org/grpc"
)

func EntityStore(params APIEntity.EntityStoreParams) error {
	var err error
	newEntity := *params.Body

	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Println(err.Error())
	}
	aRPC := alpharpc.NewAlphaCRUDRPCClient(conn)

	//means user want to update existing entity. Primary ID stored in alpha db
	if newEntity.ID > 0 {
		alphaEntity, grpcerr := aRPC.GetAlphaInformationByID(context.Background(), &alpharpc.AlphaGetByIDRequest{ID: newEntity.ID})

		if grpcerr != nil || alphaEntity == nil {
			return fmt.Errorf(grpc.ErrorDesc(grpcerr))
			// log.Println(grpcerr.Error())
			// return middleware.Error(500, grpcerr.Error())
		}

		//beta can only modify Description and Status
		newEntity.Name = alphaEntity.Name

		alphaReply, grpcerr := aRPC.UpdateAlphaInformation(context.Background(), &alpharpc.AlphaUpdateRequest{ID: newEntity.ID,
			Name: newEntity.Name, Description: newEntity.Description, Status: newEntity.Status})

		if grpcerr != nil || alphaReply == nil {
			return fmt.Errorf(grpc.ErrorDesc(grpcerr))
			// log.Println(grpcerr.Error())
			// return middleware.Error(500, grpcerr.Error())
		}

		betaEntity, err := gdbmanager.GetEntityByName(newEntity.Name)
		betaEntity.Description = newEntity.Description

		err = gdbmanager.UpdateEntity(betaEntity, betaEntity.ID)
		if err != nil {
			return err
			// return middleware.Error(500, err.Error())
		}

		return nil
		// return APIEntity.NewEntityStoreOK()

	} else {
		err = gdbmanager.StoreEntity(newEntity)
		if err != nil {
			return err
			// return middleware.Error(500, err.Error())
		}

		alphaReply, grpcerr := aRPC.UpdateAlphaInformation(context.Background(), &alpharpc.AlphaUpdateRequest{ID: newEntity.ID,
			Name: newEntity.Name, Description: newEntity.Description, Status: newEntity.Status})

		if grpcerr != nil || alphaReply == nil {
			return fmt.Errorf(grpc.ErrorDesc(grpcerr))
			// log.Println(grpcerr.Error())
			// return middleware.Error(500, grpcerr.Error())
		}
		return nil
		// return APIEntity.NewEntityStoreOK()

	}
}

func EntityDelete(params APIEntity.EntityDeleteParams) error {
	var err error
	entityName, err := url.QueryUnescape(params.EntityName)
	if err != nil {
		entityName = params.EntityName
	}

	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Println(err.Error())
	}

	//asking Alpha first since Alpha is owner of Name field, which is our key to search records
	aRPC := alpharpc.NewAlphaCRUDRPCClient(conn)
	alphaEntity, grpcerr := aRPC.DeleteAlphaInformation(context.Background(), &alpharpc.AlphaGetRequest{EntityName: entityName})

	if grpcerr != nil || alphaEntity == nil {
		return fmt.Errorf(grpc.ErrorDesc(grpcerr))
		// log.Println(grpcerr.Error())
		// return middleware.Error(500, grpcerr.Error())
	}

	entity, err := gdbmanager.GetEntityByName(entityName)
	if err != nil {
		return err
		// if err.Error() == gdbmanager.StrNoRecords {
		// 	return APIEntity.NewEntityDeleteNotFound()
		// } else {
		// 	log.Println(err.Error())
		// 	return middleware.Error(500, err.Error())
		// }
	}
	err = gdbmanager.DeleteEntityByName(entity.Name)
	if err != nil {
		return err
		// log.Println(err.Error())
		// return middleware.Error(500, err.Error())
	}
	return nil
	// return APIEntity.NewEntityDeleteOK()
}

func EntityGet(params APIEntity.EntityGetParams) ([]*models.Entity, error) {
	var err error
	entityName, err := url.QueryUnescape(params.EntityName)
	if err != nil {
		log.Println(err.Error())
		entityName = params.EntityName
	}
	log.Println(entityName)

	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	//asking Alpha first since Alpha is owner of Name field, which is our key to search records
	aRPC := alpharpc.NewAlphaCRUDRPCClient(conn)
	alphaEntity, grpcerr := aRPC.GetAlphaInformation(context.Background(), &alpharpc.AlphaGetRequest{EntityName: entityName})

	if grpcerr != nil || alphaEntity == nil {
		return nil, fmt.Errorf(grpc.ErrorDesc(grpcerr))
		// if grpcerr.Error() == gdbmanager.StrNoRecords {
		// 	return APIEntity.NewEntityGetNotFound()
		// }
		// log.Println(grpcerr.Error())
		// return middleware.Error(500, grpcerr.Error())
	}

	entity, err := gdbmanager.GetEntityByName(entityName)
	if err != nil {
		return nil, err
		// if err.Error() == gdbmanager.StrNoRecords {
		// 	return APIEntity.NewEntityGetNotFound()
		// } else {
		// 	log.Println(err.Error())
		// 	return middleware.Error(500, err.Error())
		// }
	}

	entity.ID = alphaEntity.GetID()
	entity.Status = alphaEntity.GetStatus()

	// log.Printf("%+v\n", entity)
	var responsePayload []*models.Entity
	responsePayload = append(responsePayload, &entity)
	return responsePayload, nil
	// return APIEntity.NewEntityGetOK().WithPayload(responsePayload)
}
