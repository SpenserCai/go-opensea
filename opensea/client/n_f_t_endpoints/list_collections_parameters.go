// Code generated by go-swagger; DO NOT EDIT.

package n_f_t_endpoints

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
	"github.com/go-openapi/swag"
)

// NewListCollectionsParams creates a new ListCollectionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListCollectionsParams() *ListCollectionsParams {
	return &ListCollectionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListCollectionsParamsWithTimeout creates a new ListCollectionsParams object
// with the ability to set a timeout on a request.
func NewListCollectionsParamsWithTimeout(timeout time.Duration) *ListCollectionsParams {
	return &ListCollectionsParams{
		timeout: timeout,
	}
}

// NewListCollectionsParamsWithContext creates a new ListCollectionsParams object
// with the ability to set a context for a request.
func NewListCollectionsParamsWithContext(ctx context.Context) *ListCollectionsParams {
	return &ListCollectionsParams{
		Context: ctx,
	}
}

// NewListCollectionsParamsWithHTTPClient creates a new ListCollectionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListCollectionsParamsWithHTTPClient(client *http.Client) *ListCollectionsParams {
	return &ListCollectionsParams{
		HTTPClient: client,
	}
}

/*
ListCollectionsParams contains all the parameters to send to the API endpoint

	for the list collections operation.

	Typically these are written to a http.Request.
*/
type ListCollectionsParams struct {

	/* ChainIdentifier.

	   The blockchain on which to filter the results
	*/
	ChainIdentifier *string

	/* IncludeHidden.

	   If true, will return hidden collections. Default: false
	*/
	IncludeHidden *bool

	/* Limit.

	   The number of collections to return. Must be between 1 and 100. Default: 100
	*/
	Limit *int64

	/* Next.

	   The cursor for the next page of results. This is returned from a previous request.
	*/
	Next *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list collections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListCollectionsParams) WithDefaults() *ListCollectionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list collections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListCollectionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list collections params
func (o *ListCollectionsParams) WithTimeout(timeout time.Duration) *ListCollectionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list collections params
func (o *ListCollectionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list collections params
func (o *ListCollectionsParams) WithContext(ctx context.Context) *ListCollectionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list collections params
func (o *ListCollectionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list collections params
func (o *ListCollectionsParams) WithHTTPClient(client *http.Client) *ListCollectionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list collections params
func (o *ListCollectionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChainIdentifier adds the chainIdentifier to the list collections params
func (o *ListCollectionsParams) WithChainIdentifier(chainIdentifier *string) *ListCollectionsParams {
	o.SetChainIdentifier(chainIdentifier)
	return o
}

// SetChainIdentifier adds the chainIdentifier to the list collections params
func (o *ListCollectionsParams) SetChainIdentifier(chainIdentifier *string) {
	o.ChainIdentifier = chainIdentifier
}

// WithIncludeHidden adds the includeHidden to the list collections params
func (o *ListCollectionsParams) WithIncludeHidden(includeHidden *bool) *ListCollectionsParams {
	o.SetIncludeHidden(includeHidden)
	return o
}

// SetIncludeHidden adds the includeHidden to the list collections params
func (o *ListCollectionsParams) SetIncludeHidden(includeHidden *bool) {
	o.IncludeHidden = includeHidden
}

// WithLimit adds the limit to the list collections params
func (o *ListCollectionsParams) WithLimit(limit *int64) *ListCollectionsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list collections params
func (o *ListCollectionsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNext adds the next to the list collections params
func (o *ListCollectionsParams) WithNext(next *string) *ListCollectionsParams {
	o.SetNext(next)
	return o
}

// SetNext adds the next to the list collections params
func (o *ListCollectionsParams) SetNext(next *string) {
	o.Next = next
}

// WriteToRequest writes these params to a swagger request
func (o *ListCollectionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ChainIdentifier != nil {

		// query param chain_identifier
		var qrChainIdentifier string

		if o.ChainIdentifier != nil {
			qrChainIdentifier = *o.ChainIdentifier
		}
		qChainIdentifier := qrChainIdentifier
		if qChainIdentifier != "" {

			if err := r.SetQueryParam("chain_identifier", qChainIdentifier); err != nil {
				return err
			}
		}
	}

	if o.IncludeHidden != nil {

		// query param include_hidden
		var qrIncludeHidden bool

		if o.IncludeHidden != nil {
			qrIncludeHidden = *o.IncludeHidden
		}
		qIncludeHidden := swag.FormatBool(qrIncludeHidden)
		if qIncludeHidden != "" {

			if err := r.SetQueryParam("include_hidden", qIncludeHidden); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Next != nil {

		// query param next
		var qrNext string

		if o.Next != nil {
			qrNext = *o.Next
		}
		qNext := qrNext
		if qNext != "" {

			if err := r.SetQueryParam("next", qNext); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
