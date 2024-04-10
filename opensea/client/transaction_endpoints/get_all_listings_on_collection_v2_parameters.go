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
	"github.com/go-openapi/swag"
)

// NewGetAllListingsOnCollectionV2Params creates a new GetAllListingsOnCollectionV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllListingsOnCollectionV2Params() *GetAllListingsOnCollectionV2Params {
	return &GetAllListingsOnCollectionV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllListingsOnCollectionV2ParamsWithTimeout creates a new GetAllListingsOnCollectionV2Params object
// with the ability to set a timeout on a request.
func NewGetAllListingsOnCollectionV2ParamsWithTimeout(timeout time.Duration) *GetAllListingsOnCollectionV2Params {
	return &GetAllListingsOnCollectionV2Params{
		timeout: timeout,
	}
}

// NewGetAllListingsOnCollectionV2ParamsWithContext creates a new GetAllListingsOnCollectionV2Params object
// with the ability to set a context for a request.
func NewGetAllListingsOnCollectionV2ParamsWithContext(ctx context.Context) *GetAllListingsOnCollectionV2Params {
	return &GetAllListingsOnCollectionV2Params{
		Context: ctx,
	}
}

// NewGetAllListingsOnCollectionV2ParamsWithHTTPClient creates a new GetAllListingsOnCollectionV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllListingsOnCollectionV2ParamsWithHTTPClient(client *http.Client) *GetAllListingsOnCollectionV2Params {
	return &GetAllListingsOnCollectionV2Params{
		HTTPClient: client,
	}
}

/*
GetAllListingsOnCollectionV2Params contains all the parameters to send to the API endpoint

	for the get all listings on collection v2 operation.

	Typically these are written to a http.Request.
*/
type GetAllListingsOnCollectionV2Params struct {

	/* CollectionSlug.

	   Unique string to identify a collection on OpenSea. This can be found by visiting the collection on the OpenSea website and noting the last path parameter.
	*/
	CollectionSlug string

	/* Limit.

	   The number of listings to return. Must be between 1 and 100. Default: 100
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

// WithDefaults hydrates default values in the get all listings on collection v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllListingsOnCollectionV2Params) WithDefaults() *GetAllListingsOnCollectionV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all listings on collection v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllListingsOnCollectionV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all listings on collection v2 params
func (o *GetAllListingsOnCollectionV2Params) WithTimeout(timeout time.Duration) *GetAllListingsOnCollectionV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all listings on collection v2 params
func (o *GetAllListingsOnCollectionV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all listings on collection v2 params
func (o *GetAllListingsOnCollectionV2Params) WithContext(ctx context.Context) *GetAllListingsOnCollectionV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all listings on collection v2 params
func (o *GetAllListingsOnCollectionV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all listings on collection v2 params
func (o *GetAllListingsOnCollectionV2Params) WithHTTPClient(client *http.Client) *GetAllListingsOnCollectionV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all listings on collection v2 params
func (o *GetAllListingsOnCollectionV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCollectionSlug adds the collectionSlug to the get all listings on collection v2 params
func (o *GetAllListingsOnCollectionV2Params) WithCollectionSlug(collectionSlug string) *GetAllListingsOnCollectionV2Params {
	o.SetCollectionSlug(collectionSlug)
	return o
}

// SetCollectionSlug adds the collectionSlug to the get all listings on collection v2 params
func (o *GetAllListingsOnCollectionV2Params) SetCollectionSlug(collectionSlug string) {
	o.CollectionSlug = collectionSlug
}

// WithLimit adds the limit to the get all listings on collection v2 params
func (o *GetAllListingsOnCollectionV2Params) WithLimit(limit *int64) *GetAllListingsOnCollectionV2Params {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get all listings on collection v2 params
func (o *GetAllListingsOnCollectionV2Params) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNext adds the next to the get all listings on collection v2 params
func (o *GetAllListingsOnCollectionV2Params) WithNext(next *string) *GetAllListingsOnCollectionV2Params {
	o.SetNext(next)
	return o
}

// SetNext adds the next to the get all listings on collection v2 params
func (o *GetAllListingsOnCollectionV2Params) SetNext(next *string) {
	o.Next = next
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllListingsOnCollectionV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param collection_slug
	if err := r.SetPathParam("collection_slug", o.CollectionSlug); err != nil {
		return err
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
