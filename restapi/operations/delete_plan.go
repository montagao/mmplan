// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeletePlanHandlerFunc turns a function with the right signature into a delete plan handler
type DeletePlanHandlerFunc func(DeletePlanParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeletePlanHandlerFunc) Handle(params DeletePlanParams) middleware.Responder {
	return fn(params)
}

// DeletePlanHandler interface for that can handle valid delete plan params
type DeletePlanHandler interface {
	Handle(DeletePlanParams) middleware.Responder
}

// NewDeletePlan creates a new http.Handler for the delete plan operation
func NewDeletePlan(ctx *middleware.Context, handler DeletePlanHandler) *DeletePlan {
	return &DeletePlan{Context: ctx, Handler: handler}
}

/*DeletePlan swagger:route DELETE /v1/plan/{id} deletePlan

DeletePlan delete plan API

*/
type DeletePlan struct {
	Context *middleware.Context
	Handler DeletePlanHandler
}

func (o *DeletePlan) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeletePlanParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
