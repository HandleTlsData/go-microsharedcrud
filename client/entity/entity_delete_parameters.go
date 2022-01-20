// Code generated by go-swagger; DO NOT EDIT.

package entity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewEntityDeleteParams creates a new EntityDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEntityDeleteParams() *EntityDeleteParams {
	return &EntityDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEntityDeleteParamsWithTimeout creates a new EntityDeleteParams object
// with the ability to set a timeout on a request.
func NewEntityDeleteParamsWithTimeout(timeout time.Duration) *EntityDeleteParams {
	return &EntityDeleteParams{
		timeout: timeout,
	}
}

// NewEntityDeleteParamsWithContext creates a new EntityDeleteParams object
// with the ability to set a context for a request.
func NewEntityDeleteParamsWithContext(ctx context.Context) *EntityDeleteParams {
	return &EntityDeleteParams{
		Context: ctx,
	}
}

// NewEntityDeleteParamsWithHTTPClient creates a new EntityDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewEntityDeleteParamsWithHTTPClient(client *http.Client) *EntityDeleteParams {
	return &EntityDeleteParams{
		HTTPClient: client,
	}
}

/* EntityDeleteParams contains all the parameters to send to the API endpoint
   for the entity delete operation.

   Typically these are written to a http.Request.
*/
type EntityDeleteParams struct {

	/* EntityName.

	   Name of entity that need to be returned
	*/
	EntityName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the entity delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EntityDeleteParams) WithDefaults() *EntityDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the entity delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EntityDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the entity delete params
func (o *EntityDeleteParams) WithTimeout(timeout time.Duration) *EntityDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the entity delete params
func (o *EntityDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the entity delete params
func (o *EntityDeleteParams) WithContext(ctx context.Context) *EntityDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the entity delete params
func (o *EntityDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the entity delete params
func (o *EntityDeleteParams) WithHTTPClient(client *http.Client) *EntityDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the entity delete params
func (o *EntityDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntityName adds the entityName to the entity delete params
func (o *EntityDeleteParams) WithEntityName(entityName string) *EntityDeleteParams {
	o.SetEntityName(entityName)
	return o
}

// SetEntityName adds the entityName to the entity delete params
func (o *EntityDeleteParams) SetEntityName(entityName string) {
	o.EntityName = entityName
}

// WriteToRequest writes these params to a swagger request
func (o *EntityDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entityName
	if err := r.SetPathParam("entityName", o.EntityName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
