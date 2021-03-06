// Code generated by go-swagger; DO NOT EDIT.

package entity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// EntityGetHandlerFunc turns a function with the right signature into a entity get handler
type EntityGetHandlerFunc func(EntityGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn EntityGetHandlerFunc) Handle(params EntityGetParams) middleware.Responder {
	return fn(params)
}

// EntityGetHandler interface for that can handle valid entity get params
type EntityGetHandler interface {
	Handle(EntityGetParams) middleware.Responder
}

// NewEntityGet creates a new http.Handler for the entity get operation
func NewEntityGet(ctx *middleware.Context, handler EntityGetHandler) *EntityGet {
	return &EntityGet{Context: ctx, Handler: handler}
}

/* EntityGet swagger:route GET /entity/{entityName} entity entityGet

Returns a single Entity by given Name

Finds full entity by it's Name field

*/
type EntityGet struct {
	Context *middleware.Context
	Handler EntityGetHandler
}

func (o *EntityGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewEntityGetParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
