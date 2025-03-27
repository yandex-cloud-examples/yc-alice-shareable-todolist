// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/yandex-cloud/examples/serverless/alice-shareable-todolist/app/generated/openapi/models"
)

// DeleteListNoContentCode is the HTTP code returned for type DeleteListNoContent
const DeleteListNoContentCode int = 204

/*DeleteListNoContent OK

swagger:response deleteListNoContent
*/
type DeleteListNoContent struct {
}

// NewDeleteListNoContent creates DeleteListNoContent with default headers values
func NewDeleteListNoContent() *DeleteListNoContent {

	return &DeleteListNoContent{}
}

// WriteResponse to the client
func (o *DeleteListNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*DeleteListDefault error

swagger:response deleteListDefault
*/
type DeleteListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteListDefault creates DeleteListDefault with default headers values
func NewDeleteListDefault(code int) *DeleteListDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete list default response
func (o *DeleteListDefault) WithStatusCode(code int) *DeleteListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete list default response
func (o *DeleteListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete list default response
func (o *DeleteListDefault) WithPayload(payload *models.Error) *DeleteListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete list default response
func (o *DeleteListDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
