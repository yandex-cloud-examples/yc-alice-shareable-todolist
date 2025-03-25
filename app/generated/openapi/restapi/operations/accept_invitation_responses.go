// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/yandex-cloud/examples/serverless/alice-shareable-todolist/app/generated/openapi/models"
)

// AcceptInvitationNoContentCode is the HTTP code returned for type AcceptInvitationNoContent
const AcceptInvitationNoContentCode int = 204

/*AcceptInvitationNoContent OK

swagger:response acceptInvitationNoContent
*/
type AcceptInvitationNoContent struct {
}

// NewAcceptInvitationNoContent creates AcceptInvitationNoContent with default headers values
func NewAcceptInvitationNoContent() *AcceptInvitationNoContent {

	return &AcceptInvitationNoContent{}
}

// WriteResponse to the client
func (o *AcceptInvitationNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*AcceptInvitationDefault error

swagger:response acceptInvitationDefault
*/
type AcceptInvitationDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAcceptInvitationDefault creates AcceptInvitationDefault with default headers values
func NewAcceptInvitationDefault(code int) *AcceptInvitationDefault {
	if code <= 0 {
		code = 500
	}

	return &AcceptInvitationDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the accept invitation default response
func (o *AcceptInvitationDefault) WithStatusCode(code int) *AcceptInvitationDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the accept invitation default response
func (o *AcceptInvitationDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the accept invitation default response
func (o *AcceptInvitationDefault) WithPayload(payload *models.Error) *AcceptInvitationDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the accept invitation default response
func (o *AcceptInvitationDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AcceptInvitationDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
