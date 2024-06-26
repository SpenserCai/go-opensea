// Code generated by go-swagger; DO NOT EDIT.

package transaction_endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/SpenserCai/go-opensea/opensea/models"
)

// NewPostOfferParams creates a new PostOfferParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostOfferParams() *PostOfferParams {
	return &PostOfferParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostOfferParamsWithTimeout creates a new PostOfferParams object
// with the ability to set a timeout on a request.
func NewPostOfferParamsWithTimeout(timeout time.Duration) *PostOfferParams {
	return &PostOfferParams{
		timeout: timeout,
	}
}

// NewPostOfferParamsWithContext creates a new PostOfferParams object
// with the ability to set a context for a request.
func NewPostOfferParamsWithContext(ctx context.Context) *PostOfferParams {
	return &PostOfferParams{
		Context: ctx,
	}
}

// NewPostOfferParamsWithHTTPClient creates a new PostOfferParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostOfferParamsWithHTTPClient(client *http.Client) *PostOfferParams {
	return &PostOfferParams{
		HTTPClient: client,
	}
}

/*
PostOfferParams contains all the parameters to send to the API endpoint

	for the post offer operation.

	Typically these are written to a http.Request.
*/
type PostOfferParams struct {

	// Body.
	Body *models.OrderInputWithProtocol

	/* Chain.

	   The blockchain on which to filter the results.
	*/
	Chain string

	/* Protocol.

	   The token settlement protocol to use in the request.
	*/
	Protocol string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post offer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOfferParams) WithDefaults() *PostOfferParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post offer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOfferParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post offer params
func (o *PostOfferParams) WithTimeout(timeout time.Duration) *PostOfferParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post offer params
func (o *PostOfferParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post offer params
func (o *PostOfferParams) WithContext(ctx context.Context) *PostOfferParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post offer params
func (o *PostOfferParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post offer params
func (o *PostOfferParams) WithHTTPClient(client *http.Client) *PostOfferParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post offer params
func (o *PostOfferParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post offer params
func (o *PostOfferParams) WithBody(body *models.OrderInputWithProtocol) *PostOfferParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post offer params
func (o *PostOfferParams) SetBody(body *models.OrderInputWithProtocol) {
	o.Body = body
}

// WithChain adds the chain to the post offer params
func (o *PostOfferParams) WithChain(chain string) *PostOfferParams {
	o.SetChain(chain)
	return o
}

// SetChain adds the chain to the post offer params
func (o *PostOfferParams) SetChain(chain string) {
	o.Chain = chain
}

// WithProtocol adds the protocol to the post offer params
func (o *PostOfferParams) WithProtocol(protocol string) *PostOfferParams {
	o.SetProtocol(protocol)
	return o
}

// SetProtocol adds the protocol to the post offer params
func (o *PostOfferParams) SetProtocol(protocol string) {
	o.Protocol = protocol
}

// WriteToRequest writes these params to a swagger request
func (o *PostOfferParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param chain
	if err := r.SetPathParam("chain", o.Chain); err != nil {
		return err
	}

	// path param protocol
	if err := r.SetPathParam("protocol", o.Protocol); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
