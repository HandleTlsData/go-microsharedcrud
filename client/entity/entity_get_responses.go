// Code generated by go-swagger; DO NOT EDIT.

package entity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"sharedcrud/models"
)

// EntityGetReader is a Reader for the EntityGet structure.
type EntityGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EntityGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEntityGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewEntityGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEntityGetOK creates a EntityGetOK with default headers values
func NewEntityGetOK() *EntityGetOK {
	return &EntityGetOK{}
}

/* EntityGetOK describes a response with status code 200, with default header values.

successful operation
*/
type EntityGetOK struct {
	Payload []*models.Entity
}

func (o *EntityGetOK) Error() string {
	return fmt.Sprintf("[GET /entity/{entityName}][%d] entityGetOK  %+v", 200, o.Payload)
}
func (o *EntityGetOK) GetPayload() []*models.Entity {
	return o.Payload
}

func (o *EntityGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEntityGetNotFound creates a EntityGetNotFound with default headers values
func NewEntityGetNotFound() *EntityGetNotFound {
	return &EntityGetNotFound{}
}

/* EntityGetNotFound describes a response with status code 404, with default header values.

No entity was found
*/
type EntityGetNotFound struct {
}

func (o *EntityGetNotFound) Error() string {
	return fmt.Sprintf("[GET /entity/{entityName}][%d] entityGetNotFound ", 404)
}

func (o *EntityGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
