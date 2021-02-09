// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/montagao/monplan/models"
)

// UpdatePlanOKCode is the HTTP code returned for type UpdatePlanOK
const UpdatePlanOKCode int = 200

/*UpdatePlanOK OK

swagger:response updatePlanOK
*/
type UpdatePlanOK struct {

	/*
	  In: Body
	*/
	Payload *models.Plan `json:"body,omitempty"`
}

// NewUpdatePlanOK creates UpdatePlanOK with default headers values
func NewUpdatePlanOK() *UpdatePlanOK {

	return &UpdatePlanOK{}
}

// WithPayload adds the payload to the update plan o k response
func (o *UpdatePlanOK) WithPayload(payload *models.Plan) *UpdatePlanOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update plan o k response
func (o *UpdatePlanOK) SetPayload(payload *models.Plan) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePlanOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdatePlanDefault error

swagger:response updatePlanDefault
*/
type UpdatePlanDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdatePlanDefault creates UpdatePlanDefault with default headers values
func NewUpdatePlanDefault(code int) *UpdatePlanDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdatePlanDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update plan default response
func (o *UpdatePlanDefault) WithStatusCode(code int) *UpdatePlanDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update plan default response
func (o *UpdatePlanDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update plan default response
func (o *UpdatePlanDefault) WithPayload(payload *models.Error) *UpdatePlanDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update plan default response
func (o *UpdatePlanDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePlanDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
