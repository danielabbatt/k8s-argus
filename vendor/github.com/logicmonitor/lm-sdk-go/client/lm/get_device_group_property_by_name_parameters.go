// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDeviceGroupPropertyByNameParams creates a new GetDeviceGroupPropertyByNameParams object
// with the default values initialized.
func NewGetDeviceGroupPropertyByNameParams() *GetDeviceGroupPropertyByNameParams {
	var ()
	return &GetDeviceGroupPropertyByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceGroupPropertyByNameParamsWithTimeout creates a new GetDeviceGroupPropertyByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceGroupPropertyByNameParamsWithTimeout(timeout time.Duration) *GetDeviceGroupPropertyByNameParams {
	var ()
	return &GetDeviceGroupPropertyByNameParams{

		timeout: timeout,
	}
}

// NewGetDeviceGroupPropertyByNameParamsWithContext creates a new GetDeviceGroupPropertyByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceGroupPropertyByNameParamsWithContext(ctx context.Context) *GetDeviceGroupPropertyByNameParams {
	var ()
	return &GetDeviceGroupPropertyByNameParams{

		Context: ctx,
	}
}

// NewGetDeviceGroupPropertyByNameParamsWithHTTPClient creates a new GetDeviceGroupPropertyByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceGroupPropertyByNameParamsWithHTTPClient(client *http.Client) *GetDeviceGroupPropertyByNameParams {
	var ()
	return &GetDeviceGroupPropertyByNameParams{
		HTTPClient: client,
	}
}

/*GetDeviceGroupPropertyByNameParams contains all the parameters to send to the API endpoint
for the get device group property by name operation typically these are written to a http.Request
*/
type GetDeviceGroupPropertyByNameParams struct {

	/*Fields*/
	Fields *string
	/*Gid
	  group ID

	*/
	Gid int32
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device group property by name params
func (o *GetDeviceGroupPropertyByNameParams) WithTimeout(timeout time.Duration) *GetDeviceGroupPropertyByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device group property by name params
func (o *GetDeviceGroupPropertyByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device group property by name params
func (o *GetDeviceGroupPropertyByNameParams) WithContext(ctx context.Context) *GetDeviceGroupPropertyByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device group property by name params
func (o *GetDeviceGroupPropertyByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device group property by name params
func (o *GetDeviceGroupPropertyByNameParams) WithHTTPClient(client *http.Client) *GetDeviceGroupPropertyByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device group property by name params
func (o *GetDeviceGroupPropertyByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get device group property by name params
func (o *GetDeviceGroupPropertyByNameParams) WithFields(fields *string) *GetDeviceGroupPropertyByNameParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get device group property by name params
func (o *GetDeviceGroupPropertyByNameParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithGid adds the gid to the get device group property by name params
func (o *GetDeviceGroupPropertyByNameParams) WithGid(gid int32) *GetDeviceGroupPropertyByNameParams {
	o.SetGid(gid)
	return o
}

// SetGid adds the gid to the get device group property by name params
func (o *GetDeviceGroupPropertyByNameParams) SetGid(gid int32) {
	o.Gid = gid
}

// WithName adds the name to the get device group property by name params
func (o *GetDeviceGroupPropertyByNameParams) WithName(name string) *GetDeviceGroupPropertyByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get device group property by name params
func (o *GetDeviceGroupPropertyByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceGroupPropertyByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	// path param gid
	if err := r.SetPathParam("gid", swag.FormatInt32(o.Gid)); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
