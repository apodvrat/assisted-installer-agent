// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RegisterInfraEnvHandlerFunc turns a function with the right signature into a register infra env handler
type RegisterInfraEnvHandlerFunc func(RegisterInfraEnvParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn RegisterInfraEnvHandlerFunc) Handle(params RegisterInfraEnvParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// RegisterInfraEnvHandler interface for that can handle valid register infra env params
type RegisterInfraEnvHandler interface {
	Handle(RegisterInfraEnvParams, interface{}) middleware.Responder
}

// NewRegisterInfraEnv creates a new http.Handler for the register infra env operation
func NewRegisterInfraEnv(ctx *middleware.Context, handler RegisterInfraEnvHandler) *RegisterInfraEnv {
	return &RegisterInfraEnv{Context: ctx, Handler: handler}
}

/* RegisterInfraEnv swagger:route POST /v2/infra-envs installer registerInfraEnv

Creates a new OpenShift Discovery ISO.

*/
type RegisterInfraEnv struct {
	Context *middleware.Context
	Handler RegisterInfraEnvHandler
}

func (o *RegisterInfraEnv) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRegisterInfraEnvParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
