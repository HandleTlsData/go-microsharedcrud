// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	APIEntity "sharedcrud/restapi/operations/entity"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	gdbmanager "sharedcrud/gormdb"
	"sharedcrud/restapi/alphaAPI"
	"sharedcrud/restapi/betaAPI"
	"sharedcrud/restapi/operations"
	"sharedcrud/restapi/operations/entity"
)

func EntityDeleteAlpha(params APIEntity.EntityDeleteParams) middleware.Responder {
	err := alphaAPI.EntityDelete(params)
	if err != nil {
		if err.Error() == gdbmanager.StrNoRecords {
			return APIEntity.NewEntityDeleteNotFound()
		} else {
			return middleware.Error(500, err.Error())
		}
	}
	return APIEntity.NewEntityDeleteOK()
}

func EntityGetAlpha(params APIEntity.EntityGetParams) middleware.Responder {
	entity, err := alphaAPI.EntityGet(params)
	if err != nil {
		if err.Error() == gdbmanager.StrNoRecords {
			return APIEntity.NewEntityGetNotFound()
		} else {
			return middleware.Error(500, err.Error())
		}
	}
	if entity != nil {
		return APIEntity.NewEntityGetOK().WithPayload(entity)
	}
	return middleware.Error(500, "")
}

func EntityStoreAlpha(params APIEntity.EntityStoreParams) middleware.Responder {
	err := alphaAPI.EntityStore(params)
	if err != nil {
		return middleware.Error(500, err.Error())
	}
	return APIEntity.NewEntityStoreOK()
}

func EntityStoreBeta(params APIEntity.EntityStoreParams) middleware.Responder {
	err := betaAPI.EntityStore(params)
	if err != nil {
		return middleware.Error(500, err.Error())
	}
	return APIEntity.NewEntityStoreOK()
}

func EntityDeleteBeta(params APIEntity.EntityDeleteParams) middleware.Responder {
	err := alphaAPI.EntityDelete(params)
	if err != nil {
		if err.Error() == gdbmanager.StrNoRecords {
			return APIEntity.NewEntityDeleteNotFound()
		} else {
			return middleware.Error(500, err.Error())
		}
	}
	return APIEntity.NewEntityDeleteOK()
}

func EntityGetBeta(params APIEntity.EntityGetParams) middleware.Responder {
	entity, err := alphaAPI.EntityGet(params)
	if err != nil {
		if err.Error() == gdbmanager.StrNoRecords {
			return APIEntity.NewEntityGetNotFound()
		} else {
			return middleware.Error(500, err.Error())
		}
	}
	if entity != nil {
		return APIEntity.NewEntityGetOK().WithPayload(entity)
	}
	return middleware.Error(500, "")
}

//go:generate swagger generate server --target ..\..\test-server --name Sharedcrud --spec ..\swagger\swagger.yml --principal interface{}

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

	switch gdbmanager.CurrentAppConfig {
	case "alpha":
		api.EntityEntityDeleteHandler = entity.EntityDeleteHandlerFunc(EntityDeleteAlpha)
		api.EntityEntityGetHandler = entity.EntityGetHandlerFunc(EntityGetAlpha)
		api.EntityEntityStoreHandler = entity.EntityStoreHandlerFunc(EntityStoreAlpha)
	case "beta":
		api.EntityEntityDeleteHandler = entity.EntityDeleteHandlerFunc(EntityDeleteBeta)
		api.EntityEntityGetHandler = entity.EntityGetHandlerFunc(EntityGetBeta)
		api.EntityEntityStoreHandler = entity.EntityStoreHandlerFunc(EntityStoreBeta)
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
