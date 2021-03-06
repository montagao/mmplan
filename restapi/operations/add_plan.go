// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddPlanHandlerFunc turns a function with the right signature into a add plan handler
type AddPlanHandlerFunc func(AddPlanParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddPlanHandlerFunc) Handle(params AddPlanParams) middleware.Responder {
	return fn(params)
}

// AddPlanHandler interface for that can handle valid add plan params
type AddPlanHandler interface {
	Handle(AddPlanParams) middleware.Responder
}

// NewAddPlan creates a new http.Handler for the add plan operation
func NewAddPlan(ctx *middleware.Context, handler AddPlanHandler) *AddPlan {
	return &AddPlan{Context: ctx, Handler: handler}
}

/* AddPlan swagger:route POST /v1/plan/ addPlan

AddPlan add plan API

*/
type AddPlan struct {
	Context *middleware.Context
	Handler AddPlanHandler
}

func (o *AddPlan) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddPlanParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
