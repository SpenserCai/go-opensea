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

// NewListNftsByAccountParams creates a new ListNftsByAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListNftsByAccountParams() *ListNftsByAccountParams {
	return &ListNftsByAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListNftsByAccountParamsWithTimeout creates a new ListNftsByAccountParams object
// with the ability to set a timeout on a request.
func NewListNftsByAccountParamsWithTimeout(timeout time.Duration) *ListNftsByAccountParams {
	return &ListNftsByAccountParams{
		timeout: timeout,
	}
}

// NewListNftsByAccountParamsWithContext creates a new ListNftsByAccountParams object
// with the ability to set a context for a request.
func NewListNftsByAccountParamsWithContext(ctx context.Context) *ListNftsByAccountParams {
	return &ListNftsByAccountParams{
		Context: ctx,
	}
}

// NewListNftsByAccountParamsWithHTTPClient creates a new ListNftsByAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewListNftsByAccountParamsWithHTTPClient(client *http.Client) *ListNftsByAccountParams {
	return &ListNftsByAccountParams{
		HTTPClient: client,
	}
}

/*
ListNftsByAccountParams contains all the parameters to send to the API endpoint

	for the list nfts by account operation.

	Typically these are written to a http.Request.
*/
type ListNftsByAccountParams struct {

	/* Address.

	   The unique public blockchain identifier for the contract or wallet.
	*/
	Address string

	/* Chain.

	   The blockchain on which to filter the results.
	*/
	Chain string

	/* Collection.

	   Unique string to identify a collection on OpenSea. This can be found by visiting the collection on the OpenSea website and noting the last path parameter.
	*/
	Collection *string

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

// WithDefaults hydrates default values in the list nfts by account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNftsByAccountParams) WithDefaults() *ListNftsByAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list nfts by account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNftsByAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list nfts by account params
func (o *ListNftsByAccountParams) WithTimeout(timeout time.Duration) *ListNftsByAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list nfts by account params
func (o *ListNftsByAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list nfts by account params
func (o *ListNftsByAccountParams) WithContext(ctx context.Context) *ListNftsByAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list nfts by account params
func (o *ListNftsByAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list nfts by account params
func (o *ListNftsByAccountParams) WithHTTPClient(client *http.Client) *ListNftsByAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list nfts by account params
func (o *ListNftsByAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the list nfts by account params
func (o *ListNftsByAccountParams) WithAddress(address string) *ListNftsByAccountParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the list nfts by account params
func (o *ListNftsByAccountParams) SetAddress(address string) {
	o.Address = address
}

// WithChain adds the chain to the list nfts by account params
func (o *ListNftsByAccountParams) WithChain(chain string) *ListNftsByAccountParams {
	o.SetChain(chain)
	return o
}

// SetChain adds the chain to the list nfts by account params
func (o *ListNftsByAccountParams) SetChain(chain string) {
	o.Chain = chain
}

// WithCollection adds the collection to the list nfts by account params
func (o *ListNftsByAccountParams) WithCollection(collection *string) *ListNftsByAccountParams {
	o.SetCollection(collection)
	return o
}

// SetCollection adds the collection to the list nfts by account params
func (o *ListNftsByAccountParams) SetCollection(collection *string) {
	o.Collection = collection
}

// WithLimit adds the limit to the list nfts by account params
func (o *ListNftsByAccountParams) WithLimit(limit *int64) *ListNftsByAccountParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list nfts by account params
func (o *ListNftsByAccountParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNext adds the next to the list nfts by account params
func (o *ListNftsByAccountParams) WithNext(next *string) *ListNftsByAccountParams {
	o.SetNext(next)
	return o
}

// SetNext adds the next to the list nfts by account params
func (o *ListNftsByAccountParams) SetNext(next *string) {
	o.Next = next
}

// WriteToRequest writes these params to a swagger request
func (o *ListNftsByAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param address
	if err := r.SetPathParam("address", o.Address); err != nil {
		return err
	}

	// path param chain
	if err := r.SetPathParam("chain", o.Chain); err != nil {
		return err
	}

	if o.Collection != nil {

		// query param collection
		var qrCollection string

		if o.Collection != nil {
			qrCollection = *o.Collection
		}
		qCollection := qrCollection
		if qCollection != "" {

			if err := r.SetQueryParam("collection", qCollection); err != nil {
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