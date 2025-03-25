// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/yandex-cloud/examples/serverless/alice-shareable-todolist/app/generated/openapi/models"
)

// InviteUserHandlerFunc turns a function with the right signature into a invite user handler
type InviteUserHandlerFunc func(InviteUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn InviteUserHandlerFunc) Handle(params InviteUserParams) middleware.Responder {
	return fn(params)
}

// InviteUserHandler interface for that can handle valid invite user params
type InviteUserHandler interface {
	Handle(InviteUserParams) middleware.Responder
}

// NewInviteUser creates a new http.Handler for the invite user operation
func NewInviteUser(ctx *middleware.Context, handler InviteUserHandler) *InviteUser {
	return &InviteUser{Context: ctx, Handler: handler}
}

/* InviteUser swagger:route POST /api/todo-lists/{list_id}/inviteUser inviteUser

InviteUser invite user API

*/
type InviteUser struct {
	Context *middleware.Context
	Handler InviteUserHandler
}

func (o *InviteUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewInviteUserParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// InviteUserBody invite user body
//
// swagger:model InviteUserBody
type InviteUserBody struct {

	// access mode
	// Required: true
	AccessMode *models.AccessMode `json:"access_mode"`

	// invitee
	// Required: true
	Invitee *string `json:"invitee"`
}

// Validate validates this invite user body
func (o *InviteUserBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccessMode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateInvitee(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InviteUserBody) validateAccessMode(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"access_mode", "body", o.AccessMode); err != nil {
		return err
	}

	if err := validate.Required("body"+"."+"access_mode", "body", o.AccessMode); err != nil {
		return err
	}

	if o.AccessMode != nil {
		if err := o.AccessMode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "access_mode")
			}
			return err
		}
	}

	return nil
}

func (o *InviteUserBody) validateInvitee(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"invitee", "body", o.Invitee); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this invite user body based on the context it is used
func (o *InviteUserBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAccessMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InviteUserBody) contextValidateAccessMode(ctx context.Context, formats strfmt.Registry) error {

	if o.AccessMode != nil {
		if err := o.AccessMode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "access_mode")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *InviteUserBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InviteUserBody) UnmarshalBinary(b []byte) error {
	var res InviteUserBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
