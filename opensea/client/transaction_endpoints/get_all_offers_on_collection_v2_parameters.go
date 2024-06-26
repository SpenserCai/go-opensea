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

// NewGetAllOffersOnCollectionV2Params creates a new GetAllOffersOnCollectionV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllOffersOnCollectionV2Params() *GetAllOffersOnCollectionV2Params {
	return &GetAllOffersOnCollectionV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllOffersOnCollectionV2ParamsWithTimeout creates a new GetAllOffersOnCollectionV2Params object
// with the ability to set a timeout on a request.
func NewGetAllOffersOnCollectionV2ParamsWithTimeout(timeout time.Duration) *GetAllOffersOnCollectionV2Params {
	return &GetAllOffersOnCollectionV2Params{
		timeout: timeout,
	}
}

// NewGetAllOffersOnCollectionV2ParamsWithContext creates a new GetAllOffersOnCollectionV2Params object
// with the ability to set a context for a request.
func NewGetAllOffersOnCollectionV2ParamsWithContext(ctx context.Context) *GetAllOffersOnCollectionV2Params {
	return &GetAllOffersOnCollectionV2Params{
		Context: ctx,
	}
}

// NewGetAllOffersOnCollectionV2ParamsWithHTTPClient creates a new GetAllOffersOnCollectionV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllOffersOnCollectionV2ParamsWithHTTPClient(client *http.Client) *GetAllOffersOnCollectionV2Params {
	return &GetAllOffersOnCollectionV2Params{
		HTTPClient: client,
	}
}

/*
GetAllOffersOnCollectionV2Params contains all the parameters to send to the API endpoint

	for the get all offers on collection v2 operation.

	Typically these are written to a http.Request.
*/
type GetAllOffersOnCollectionV2Params struct {

	/* CollectionSlug.

	   Unique string to identify a collection on OpenSea. This can be found by visiting the collection on the OpenSea website and noting the last path parameter.
	*/
	CollectionSlug string

	/* Limit.

	   The number of offers to return. Must be between 1 and 100. Default: 100
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

// WithDefaults hydrates default values in the get all offers on collection v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllOffersOnCollectionV2Params) WithDefaults() *GetAllOffersOnCollectionV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all offers on collection v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllOffersOnCollectionV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all offers on collection v2 params
func (o *GetAllOffersOnCollectionV2Params) WithTimeout(timeout time.Duration) *GetAllOffersOnCollectionV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all offers on collection v2 params
func (o *GetAllOffersOnCollectionV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all offers on collection v2 params
func (o *GetAllOffersOnCollectionV2Params) WithContext(ctx context.Context) *GetAllOffersOnCollectionV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all offers on collection v2 params
func (o *GetAllOffersOnCollectionV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all offers on collection v2 params
func (o *GetAllOffersOnCollectionV2Params) WithHTTPClient(client *http.Client) *GetAllOffersOnCollectionV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all offers on collection v2 params
func (o *GetAllOffersOnCollectionV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCollectionSlug adds the collectionSlug to the get all offers on collection v2 params
func (o *GetAllOffersOnCollectionV2Params) WithCollectionSlug(collectionSlug string) *GetAllOffersOnCollectionV2Params {
	o.SetCollectionSlug(collectionSlug)
	return o
}

// SetCollectionSlug adds the collectionSlug to the get all offers on collection v2 params
func (o *GetAllOffersOnCollectionV2Params) SetCollectionSlug(collectionSlug string) {
	o.CollectionSlug = collectionSlug
}

// WithLimit adds the limit to the get all offers on collection v2 params
func (o *GetAllOffersOnCollectionV2Params) WithLimit(limit *int64) *GetAllOffersOnCollectionV2Params {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get all offers on collection v2 params
func (o *GetAllOffersOnCollectionV2Params) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNext adds the next to the get all offers on collection v2 params
func (o *GetAllOffersOnCollectionV2Params) WithNext(next *string) *GetAllOffersOnCollectionV2Params {
	o.SetNext(next)
	return o
}

// SetNext adds the next to the get all offers on collection v2 params
func (o *GetAllOffersOnCollectionV2Params) SetNext(next *string) {
	o.Next = next
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllOffersOnCollectionV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
