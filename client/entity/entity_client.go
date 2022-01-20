// Code generated by go-swagger; DO NOT EDIT.

package entity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new entity API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for entity API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	EntityDelete(params *EntityDeleteParams, opts ...ClientOption) (*EntityDeleteOK, error)

	EntityGet(params *EntityGetParams, opts ...ClientOption) (*EntityGetOK, error)

	EntityStore(params *EntityStoreParams, opts ...ClientOption) (*EntityStoreOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  EntityDelete deletes a single entity by given name

  Delete full entity by it's Name field
*/
func (a *Client) EntityDelete(params *EntityDeleteParams, opts ...ClientOption) (*EntityDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEntityDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "entityDelete",
		Method:             "DELETE",
		PathPattern:        "/entity/{entityName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EntityDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EntityDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for entityDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EntityGet returns a single entity by given name

  Finds full entity by it's Name field
*/
func (a *Client) EntityGet(params *EntityGetParams, opts ...ClientOption) (*EntityGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEntityGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "entityGet",
		Method:             "GET",
		PathPattern:        "/entity/{entityName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EntityGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EntityGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for entityGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EntityStore adds a new entity if such name is not exists update otherwise
*/
func (a *Client) EntityStore(params *EntityStoreParams, opts ...ClientOption) (*EntityStoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEntityStoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "entityStore",
		Method:             "POST",
		PathPattern:        "/entity",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EntityStoreReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EntityStoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for entityStore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
