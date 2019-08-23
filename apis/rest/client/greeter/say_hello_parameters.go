// Code generated by go-swagger; DO NOT EDIT.

package greeter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSayHelloParams creates a new SayHelloParams object
// with the default values initialized.
func NewSayHelloParams() *SayHelloParams {
	var ()
	return &SayHelloParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSayHelloParamsWithTimeout creates a new SayHelloParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSayHelloParamsWithTimeout(timeout time.Duration) *SayHelloParams {
	var ()
	return &SayHelloParams{

		timeout: timeout,
	}
}

// NewSayHelloParamsWithContext creates a new SayHelloParams object
// with the default values initialized, and the ability to set a context for a request
func NewSayHelloParamsWithContext(ctx context.Context) *SayHelloParams {
	var ()
	return &SayHelloParams{

		Context: ctx,
	}
}

// NewSayHelloParamsWithHTTPClient creates a new SayHelloParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSayHelloParamsWithHTTPClient(client *http.Client) *SayHelloParams {
	var ()
	return &SayHelloParams{
		HTTPClient: client,
	}
}

/*SayHelloParams contains all the parameters to send to the API endpoint
for the say hello operation typically these are written to a http.Request
*/
type SayHelloParams struct {

	/*Name*/
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the say hello params
func (o *SayHelloParams) WithTimeout(timeout time.Duration) *SayHelloParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the say hello params
func (o *SayHelloParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the say hello params
func (o *SayHelloParams) WithContext(ctx context.Context) *SayHelloParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the say hello params
func (o *SayHelloParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the say hello params
func (o *SayHelloParams) WithHTTPClient(client *http.Client) *SayHelloParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the say hello params
func (o *SayHelloParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the say hello params
func (o *SayHelloParams) WithName(name *string) *SayHelloParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the say hello params
func (o *SayHelloParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *SayHelloParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
