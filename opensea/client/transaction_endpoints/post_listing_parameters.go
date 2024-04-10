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

// NewPostListingParams creates a new PostListingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostListingParams() *PostListingParams {
	return &PostListingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostListingParamsWithTimeout creates a new PostListingParams object
// with the ability to set a timeout on a request.
func NewPostListingParamsWithTimeout(timeout time.Duration) *PostListingParams {
	return &PostListingParams{
		timeout: timeout,
	}
}

// NewPostListingParamsWithContext creates a new PostListingParams object
// with the ability to set a context for a request.
func NewPostListingParamsWithContext(ctx context.Context) *PostListingParams {
	return &PostListingParams{
		Context: ctx,
	}
}

// NewPostListingParamsWithHTTPClient creates a new PostListingParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostListingParamsWithHTTPClient(client *http.Client) *PostListingParams {
	return &PostListingParams{
		HTTPClient: client,
	}
}

/*
PostListingParams contains all the parameters to send to the API endpoint

	for the post listing operation.

	Typically these are written to a http.Request.
*/
type PostListingParams struct {

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

// WithDefaults hydrates default values in the post listing params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostListingParams) WithDefaults() *PostListingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post listing params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostListingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post listing params
func (o *PostListingParams) WithTimeout(timeout time.Duration) *PostListingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post listing params
func (o *PostListingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post listing params
func (o *PostListingParams) WithContext(ctx context.Context) *PostListingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post listing params
func (o *PostListingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post listing params
func (o *PostListingParams) WithHTTPClient(client *http.Client) *PostListingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post listing params
func (o *PostListingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post listing params
func (o *PostListingParams) WithBody(body *models.OrderInputWithProtocol) *PostListingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post listing params
func (o *PostListingParams) SetBody(body *models.OrderInputWithProtocol) {
	o.Body = body
}

// WithChain adds the chain to the post listing params
func (o *PostListingParams) WithChain(chain string) *PostListingParams {
	o.SetChain(chain)
	return o
}

// SetChain adds the chain to the post listing params
func (o *PostListingParams) SetChain(chain string) {
	o.Chain = chain
}

// WithProtocol adds the protocol to the post listing params
func (o *PostListingParams) WithProtocol(protocol string) *PostListingParams {
	o.SetProtocol(protocol)
	return o
}

// SetProtocol adds the protocol to the post listing params
func (o *PostListingParams) SetProtocol(protocol string) {
	o.Protocol = protocol
}

// WriteToRequest writes these params to a swagger request
func (o *PostListingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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