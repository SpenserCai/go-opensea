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

// NewListNftsByCollectionParams creates a new ListNftsByCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListNftsByCollectionParams() *ListNftsByCollectionParams {
	return &ListNftsByCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListNftsByCollectionParamsWithTimeout creates a new ListNftsByCollectionParams object
// with the ability to set a timeout on a request.
func NewListNftsByCollectionParamsWithTimeout(timeout time.Duration) *ListNftsByCollectionParams {
	return &ListNftsByCollectionParams{
		timeout: timeout,
	}
}

// NewListNftsByCollectionParamsWithContext creates a new ListNftsByCollectionParams object
// with the ability to set a context for a request.
func NewListNftsByCollectionParamsWithContext(ctx context.Context) *ListNftsByCollectionParams {
	return &ListNftsByCollectionParams{
		Context: ctx,
	}
}

// NewListNftsByCollectionParamsWithHTTPClient creates a new ListNftsByCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewListNftsByCollectionParamsWithHTTPClient(client *http.Client) *ListNftsByCollectionParams {
	return &ListNftsByCollectionParams{
		HTTPClient: client,
	}
}

/*
ListNftsByCollectionParams contains all the parameters to send to the API endpoint

	for the list nfts by collection operation.

	Typically these are written to a http.Request.
*/
type ListNftsByCollectionParams struct {

	/* CollectionSlug.

	   Unique string to identify a collection on OpenSea. This can be found by visiting the collection on the OpenSea website and noting the last path parameter.
	*/
	CollectionSlug string

	/* Limit.

	   The number of NFTs to return. Must be between 1 and 50. Default: 50
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

// WithDefaults hydrates default values in the list nfts by collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNftsByCollectionParams) WithDefaults() *ListNftsByCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list nfts by collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNftsByCollectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list nfts by collection params
func (o *ListNftsByCollectionParams) WithTimeout(timeout time.Duration) *ListNftsByCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list nfts by collection params
func (o *ListNftsByCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list nfts by collection params
func (o *ListNftsByCollectionParams) WithContext(ctx context.Context) *ListNftsByCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list nfts by collection params
func (o *ListNftsByCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list nfts by collection params
func (o *ListNftsByCollectionParams) WithHTTPClient(client *http.Client) *ListNftsByCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list nfts by collection params
func (o *ListNftsByCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCollectionSlug adds the collectionSlug to the list nfts by collection params
func (o *ListNftsByCollectionParams) WithCollectionSlug(collectionSlug string) *ListNftsByCollectionParams {
	o.SetCollectionSlug(collectionSlug)
	return o
}

// SetCollectionSlug adds the collectionSlug to the list nfts by collection params
func (o *ListNftsByCollectionParams) SetCollectionSlug(collectionSlug string) {
	o.CollectionSlug = collectionSlug
}

// WithLimit adds the limit to the list nfts by collection params
func (o *ListNftsByCollectionParams) WithLimit(limit *int64) *ListNftsByCollectionParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list nfts by collection params
func (o *ListNftsByCollectionParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNext adds the next to the list nfts by collection params
func (o *ListNftsByCollectionParams) WithNext(next *string) *ListNftsByCollectionParams {
	o.SetNext(next)
	return o
}

// SetNext adds the next to the list nfts by collection params
func (o *ListNftsByCollectionParams) SetNext(next *string) {
	o.Next = next
}

// WriteToRequest writes these params to a swagger request
func (o *ListNftsByCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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