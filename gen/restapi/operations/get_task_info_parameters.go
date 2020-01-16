// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetTaskInfoParams creates a new GetTaskInfoParams object
// no default values defined in spec.
func NewGetTaskInfoParams() GetTaskInfoParams {

	return GetTaskInfoParams{}
}

// GetTaskInfoParams contains all the bound params for the get task info operation
// typically these are obtained from a http.Request
//
// swagger:parameters getTaskInfo
type GetTaskInfoParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The ID of the task to retrieve info for
	  Required: true
	  In: path
	*/
	TaskID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetTaskInfoParams() beforehand.
func (o *GetTaskInfoParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rTaskID, rhkTaskID, _ := route.Params.GetOK("taskID")
	if err := o.bindTaskID(rTaskID, rhkTaskID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindTaskID binds and validates parameter TaskID from path.
func (o *GetTaskInfoParams) bindTaskID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.TaskID = raw

	return nil
}
