// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTaskInfoHandlerFunc turns a function with the right signature into a get task info handler
type GetTaskInfoHandlerFunc func(GetTaskInfoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTaskInfoHandlerFunc) Handle(params GetTaskInfoParams) middleware.Responder {
	return fn(params)
}

// GetTaskInfoHandler interface for that can handle valid get task info params
type GetTaskInfoHandler interface {
	Handle(GetTaskInfoParams) middleware.Responder
}

// NewGetTaskInfo creates a new http.Handler for the get task info operation
func NewGetTaskInfo(ctx *middleware.Context, handler GetTaskInfoHandler) *GetTaskInfo {
	return &GetTaskInfo{Context: ctx, Handler: handler}
}

/*GetTaskInfo swagger:route GET /astrolabe/tasks/{taskID} getTaskInfo

Gets info about a running or recently completed task

*/
type GetTaskInfo struct {
	Context *middleware.Context
	Handler GetTaskInfoHandler
}

func (o *GetTaskInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTaskInfoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
