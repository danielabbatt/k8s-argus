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

// NewGetDatasourceByIDParams creates a new GetDatasourceByIDParams object
// with the default values initialized.
func NewGetDatasourceByIDParams() *GetDatasourceByIDParams {
	var (
		formatDefault = string("json")
	)
	return &GetDatasourceByIDParams{
		Format: &formatDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatasourceByIDParamsWithTimeout creates a new GetDatasourceByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDatasourceByIDParamsWithTimeout(timeout time.Duration) *GetDatasourceByIDParams {
	var (
		formatDefault = string("json")
	)
	return &GetDatasourceByIDParams{
		Format: &formatDefault,

		timeout: timeout,
	}
}

// NewGetDatasourceByIDParamsWithContext creates a new GetDatasourceByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDatasourceByIDParamsWithContext(ctx context.Context) *GetDatasourceByIDParams {
	var (
		formatDefault = string("json")
	)
	return &GetDatasourceByIDParams{
		Format: &formatDefault,

		Context: ctx,
	}
}

// NewGetDatasourceByIDParamsWithHTTPClient creates a new GetDatasourceByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDatasourceByIDParamsWithHTTPClient(client *http.Client) *GetDatasourceByIDParams {
	var (
		formatDefault = string("json")
	)
	return &GetDatasourceByIDParams{
		Format:     &formatDefault,
		HTTPClient: client,
	}
}

/*GetDatasourceByIDParams contains all the parameters to send to the API endpoint
for the get datasource by Id operation typically these are written to a http.Request
*/
type GetDatasourceByIDParams struct {

	/*Fields*/
	Fields *string
	/*Format*/
	Format *string
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get datasource by Id params
func (o *GetDatasourceByIDParams) WithTimeout(timeout time.Duration) *GetDatasourceByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get datasource by Id params
func (o *GetDatasourceByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get datasource by Id params
func (o *GetDatasourceByIDParams) WithContext(ctx context.Context) *GetDatasourceByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get datasource by Id params
func (o *GetDatasourceByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get datasource by Id params
func (o *GetDatasourceByIDParams) WithHTTPClient(client *http.Client) *GetDatasourceByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get datasource by Id params
func (o *GetDatasourceByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get datasource by Id params
func (o *GetDatasourceByIDParams) WithFields(fields *string) *GetDatasourceByIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get datasource by Id params
func (o *GetDatasourceByIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFormat adds the format to the get datasource by Id params
func (o *GetDatasourceByIDParams) WithFormat(format *string) *GetDatasourceByIDParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get datasource by Id params
func (o *GetDatasourceByIDParams) SetFormat(format *string) {
	o.Format = format
}

// WithID adds the id to the get datasource by Id params
func (o *GetDatasourceByIDParams) WithID(id int32) *GetDatasourceByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get datasource by Id params
func (o *GetDatasourceByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatasourceByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Format != nil {

		// query param format
		var qrFormat string
		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {
			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
