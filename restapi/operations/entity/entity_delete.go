// Code generated by go-swagger; DO NOT EDIT.

package entity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// EntityDeleteHandlerFunc turns a function with the right signature into a entity delete handler
type EntityDeleteHandlerFunc func(EntityDeleteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn EntityDeleteHandlerFunc) Handle(params EntityDeleteParams) middleware.Responder {
	return fn(params)
}

// EntityDeleteHandler interface for that can handle valid entity delete params
type EntityDeleteHandler interface {
	Handle(EntityDeleteParams) middleware.Responder
}

// NewEntityDelete creates a new http.Handler for the entity delete operation
func NewEntityDelete(ctx *middleware.Context, handler EntityDeleteHandler) *EntityDelete {
	return &EntityDelete{Context: ctx, Handler: handler}
}

/* EntityDelete swagger:route DELETE /entity/{entityName} entity entityDelete

Delete a single Entity by given Name

Delete full entity by it's Name field

*/
type EntityDelete struct {
	Context *middleware.Context
	Handler EntityDeleteHandler
}

func (o *EntityDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewEntityDeleteParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
