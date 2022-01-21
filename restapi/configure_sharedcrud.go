// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"net/url"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"google.golang.org/grpc"

	alpharpc "sharedcrud/apirpc/alpha"
	betarpc "sharedcrud/apirpc/beta"
	"sharedcrud/dbmanager"
	gdbmanager "sharedcrud/gormdb"
	"sharedcrud/models"
	"sharedcrud/restapi/operations"
	"sharedcrud/restapi/operations/entity"
	APIEntity "sharedcrud/restapi/operations/entity"
)

//go:generate swagger generate server --target ..\..\test-server --name Sharedcrud --spec ..\swagger\swagger.yml --principal interface{}

func AlphaEntityDelete(params APIEntity.EntityDeleteParams) middleware.Responder {
	var err error
	entityName, err := url.QueryUnescape(params.EntityName)
	if err != nil {
		entityName = params.EntityName
	}
	entity, err := gdbmanager.GetEntityByName(entityName)
	if err != nil {
		if err.Error() == gdbmanager.StrNoRecords {
			return APIEntity.NewEntityDeleteNotFound()
		} else {
			log.Println(err.Error())
			return middleware.Error(500, err.Error())
		}
	}
	err = gdbmanager.DeleteEntityByName(entity.Name)
	if err != nil {
		log.Println(err.Error())
		return middleware.Error(500, err.Error())
	}
	return APIEntity.NewEntityDeleteOK()
}

func AlphaEntityGet(params APIEntity.EntityGetParams) middleware.Responder {
	var err error
	entityName, err := url.QueryUnescape(params.EntityName)
	if err != nil {
		log.Println(err.Error())
		entityName = params.EntityName
	}
	log.Println(entityName)

	entity, err := gdbmanager.GetEntityByName(entityName)
	if err != nil {
		if err.Error() == gdbmanager.StrNoRecords {
			return APIEntity.NewEntityGetNotFound()
		} else {
			log.Println(err.Error())
			return middleware.Error(500, err.Error())
		}
	}

	conn, err := grpc.Dial(":8001", grpc.WithInsecure())
	if err != nil {
		log.Println(err.Error())
	}

	//asking Beta microservice for entity from its own db
	bRPC := betarpc.NewBetaCRUDRPCClient(conn)
	betaEntity, grpcerr := bRPC.GetBetaInformation(context.Background(), &betarpc.BetaGetRequest{EntityName: entity.Name})

	if grpcerr != nil || betaEntity == nil {
		log.Println(grpcerr.Error())
		return middleware.Error(500, grpcerr.Error())
	}

	// Beta always contains very last Description since it is only responsible for Description modifications
	entity.Description = betaEntity.GetDescription()

	log.Printf("%+v\n", entity)
	var responsePayload []*models.Entity
	responsePayload = append(responsePayload, &entity)
	return APIEntity.NewEntityGetOK().WithPayload(responsePayload)

}

func AlphaEntityStore(params APIEntity.EntityStoreParams) middleware.Responder {
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
		log.Println("update existing entity")

		if err != nil {
			return middleware.Error(500, err.Error())
		}

		betaDBEntity, grpcerr := bRPC.GetBetaInformation(context.Background(), &betarpc.BetaGetRequest{EntityName: oldEntity.Name})

		if grpcerr != nil || betaDBEntity == nil {
			log.Println(grpcerr.Error())
			return middleware.Error(500, grpcerr.Error())
		}

		//we can't modify description here
		newEntity.Description = betaDBEntity.Description

		err = gdbmanager.UpdateEntity(newEntity, oldEntity.ID)
		if err != nil {
			return middleware.Error(500, err.Error())
		}
		betaResponse, grpcerr := bRPC.UpdateBetaInformation(context.Background(), &betarpc.BetaUpdateRequest{ID: 0,
			Name: newEntity.Name, Status: newEntity.Status, Description: newEntity.Description})

		if grpcerr != nil || betaResponse == nil {
			log.Println(grpcerr.Error())
			return middleware.Error(500, grpcerr.Error())
		}

	} else {
		log.Println("create a new entity")
		err = gdbmanager.StoreEntity(*params.Body)
		if err != nil {
			return middleware.Error(500, err.Error())
		}
		betaEntity, grpcerr := bRPC.UpdateBetaInformation(context.Background(), &betarpc.BetaUpdateRequest{ID: 0,
			Name: newEntity.Name, Status: newEntity.Status, Description: newEntity.Description})

		if grpcerr != nil || betaEntity == nil {
			log.Println(grpcerr.Error())
			return middleware.Error(500, grpcerr.Error())
		}
	}

	return APIEntity.NewEntityStoreOK()
}

func BetaEntityStore(params APIEntity.EntityStoreParams) middleware.Responder {
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
			log.Println(grpcerr.Error())
			return middleware.Error(500, grpcerr.Error())
		}

		//beta can only modify Description and Status
		newEntity.Name = alphaEntity.Name

		alphaReply, grpcerr := aRPC.UpdateAlphaInformation(context.Background(), &alpharpc.AlphaUpdateRequest{ID: newEntity.ID,
			Name: newEntity.Name, Description: newEntity.Description, Status: newEntity.Status})

		if grpcerr != nil || alphaReply == nil {
			log.Println(grpcerr.Error())
			return middleware.Error(500, grpcerr.Error())
		}

		betaEntity, err := gdbmanager.GetEntityByName(newEntity.Name)
		betaEntity.Description = newEntity.Description

		err = gdbmanager.UpdateEntity(betaEntity, betaEntity.ID)
		if err != nil {
			return middleware.Error(500, err.Error())
		}

		return APIEntity.NewEntityStoreOK()

	} else {
		err = gdbmanager.StoreEntity(newEntity)
		if err != nil {
			return middleware.Error(500, err.Error())
		}

		alphaReply, grpcerr := aRPC.UpdateAlphaInformation(context.Background(), &alpharpc.AlphaUpdateRequest{ID: newEntity.ID,
			Name: newEntity.Name, Description: newEntity.Description, Status: newEntity.Status})

		if grpcerr != nil || alphaReply == nil {
			log.Println(grpcerr.Error())
			return middleware.Error(500, grpcerr.Error())
		}

		return APIEntity.NewEntityStoreOK()

	}
}

func BetaEntityDelete(params APIEntity.EntityDeleteParams) middleware.Responder {
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
		log.Println(grpcerr.Error())
		return middleware.Error(500, grpcerr.Error())
	}

	entity, err := gdbmanager.GetEntityByName(entityName)
	if err != nil {
		if err.Error() == gdbmanager.StrNoRecords {
			return APIEntity.NewEntityDeleteNotFound()
		} else {
			log.Println(err.Error())
			return middleware.Error(500, err.Error())
		}
	}
	err = gdbmanager.DeleteEntityByName(entity.Name)
	if err != nil {
		log.Println(err.Error())
		return middleware.Error(500, err.Error())
	}
	return APIEntity.NewEntityDeleteOK()
}

func BetaEntityGet(params APIEntity.EntityGetParams) middleware.Responder {
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
	}

	//asking Alpha first since Alpha is owner of Name field, which is our key to search records
	aRPC := alpharpc.NewAlphaCRUDRPCClient(conn)
	alphaEntity, grpcerr := aRPC.GetAlphaInformation(context.Background(), &alpharpc.AlphaGetRequest{EntityName: entityName})

	if grpcerr != nil || alphaEntity == nil {
		if grpcerr.Error() == gdbmanager.StrNoRecords {
			return APIEntity.NewEntityGetNotFound()
		}
		log.Println(grpcerr.Error())
		return middleware.Error(500, grpcerr.Error())
	}

	entity, err := gdbmanager.GetEntityByName(entityName)
	if err != nil {
		if err.Error() == gdbmanager.StrNoRecords {
			return APIEntity.NewEntityGetNotFound()
		} else {
			log.Println(err.Error())
			return middleware.Error(500, err.Error())
		}
	}

	entity.ID = alphaEntity.GetID()
	entity.Status = alphaEntity.GetStatus()

	log.Printf("%+v\n", entity)
	var responsePayload []*models.Entity
	responsePayload = append(responsePayload, &entity)
	return APIEntity.NewEntityGetOK().WithPayload(responsePayload)
}

func configureFlags(api *operations.SharedcrudAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SharedcrudAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	gdbmanager.InitDB()

	switch dbmanager.CurrentAppConfig {
	case "alpha":
		api.EntityEntityDeleteHandler = entity.EntityDeleteHandlerFunc(AlphaEntityDelete)
		api.EntityEntityGetHandler = entity.EntityGetHandlerFunc(AlphaEntityGet)
		api.EntityEntityStoreHandler = entity.EntityStoreHandlerFunc(AlphaEntityStore)
	case "beta":
		api.EntityEntityDeleteHandler = entity.EntityDeleteHandlerFunc(BetaEntityDelete)
		api.EntityEntityGetHandler = entity.EntityGetHandlerFunc(BetaEntityGet)
		api.EntityEntityStoreHandler = entity.EntityStoreHandlerFunc(BetaEntityStore)
	default:
		log.Fatal("Unknown App Config. API Functions unimplemented")
	}

	if api.EntityEntityDeleteHandler == nil {
		api.EntityEntityDeleteHandler = entity.EntityDeleteHandlerFunc(func(params entity.EntityDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation entity.EntityDelete has not yet been implemented")
		})
	}
	if api.EntityEntityGetHandler == nil {
		api.EntityEntityGetHandler = entity.EntityGetHandlerFunc(func(params entity.EntityGetParams) middleware.Responder {
			return middleware.NotImplemented("operation entity.EntityGet has not yet been implemented")
		})
	}
	if api.EntityEntityStoreHandler == nil {
		api.EntityEntityStoreHandler = entity.EntityStoreHandlerFunc(func(params entity.EntityStoreParams) middleware.Responder {
			return middleware.NotImplemented("operation entity.EntityStore has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
