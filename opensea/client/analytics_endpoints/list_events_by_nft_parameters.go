// Code generated by go-swagger; DO NOT EDIT.

package analytics_endpoints

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

// NewListEventsByNftParams creates a new ListEventsByNftParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListEventsByNftParams() *ListEventsByNftParams {
	return &ListEventsByNftParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListEventsByNftParamsWithTimeout creates a new ListEventsByNftParams object
// with the ability to set a timeout on a request.
func NewListEventsByNftParamsWithTimeout(timeout time.Duration) *ListEventsByNftParams {
	return &ListEventsByNftParams{
		timeout: timeout,
	}
}

// NewListEventsByNftParamsWithContext creates a new ListEventsByNftParams object
// with the ability to set a context for a request.
func NewListEventsByNftParamsWithContext(ctx context.Context) *ListEventsByNftParams {
	return &ListEventsByNftParams{
		Context: ctx,
	}
}

// NewListEventsByNftParamsWithHTTPClient creates a new ListEventsByNftParams object
// with the ability to set a custom HTTPClient for a request.
func NewListEventsByNftParamsWithHTTPClient(client *http.Client) *ListEventsByNftParams {
	return &ListEventsByNftParams{
		HTTPClient: client,
	}
}

/*
ListEventsByNftParams contains all the parameters to send to the API endpoint

	for the list events by nft operation.

	Typically these are written to a http.Request.
*/
type ListEventsByNftParams struct {

	/* Address.

	   The unique public blockchain identifier for the contract or wallet.
	*/
	Address string

	/* After.

	   Filter to only include events that occurred at or after the given timestamp. The Unix epoch timstamp must be in seconds
	*/
	After *float64

	/* Before.

	   Filter to only include events that occurred before the given timestamp. The Unix epoch timstamp must be in seconds.
	*/
	Before *float64

	/* Chain.

	   The blockchain on which to filter the results.
	*/
	Chain string

	/* EventType.

	   The type of event to filter by. If not provided, only sales will be returned.
	*/
	EventType []string

	/* Identifier.

	   The NFT token id.
	*/
	Identifier string

	/* Next.

	   The cursor for the next page of results. This is returned from a previous request.
	*/
	Next *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list events by nft params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListEventsByNftParams) WithDefaults() *ListEventsByNftParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list events by nft params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListEventsByNftParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list events by nft params
func (o *ListEventsByNftParams) WithTimeout(timeout time.Duration) *ListEventsByNftParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list events by nft params
func (o *ListEventsByNftParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list events by nft params
func (o *ListEventsByNftParams) WithContext(ctx context.Context) *ListEventsByNftParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list events by nft params
func (o *ListEventsByNftParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list events by nft params
func (o *ListEventsByNftParams) WithHTTPClient(client *http.Client) *ListEventsByNftParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list events by nft params
func (o *ListEventsByNftParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the list events by nft params
func (o *ListEventsByNftParams) WithAddress(address string) *ListEventsByNftParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the list events by nft params
func (o *ListEventsByNftParams) SetAddress(address string) {
	o.Address = address
}

// WithAfter adds the after to the list events by nft params
func (o *ListEventsByNftParams) WithAfter(after *float64) *ListEventsByNftParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the list events by nft params
func (o *ListEventsByNftParams) SetAfter(after *float64) {
	o.After = after
}

// WithBefore adds the before to the list events by nft params
func (o *ListEventsByNftParams) WithBefore(before *float64) *ListEventsByNftParams {
	o.SetBefore(before)
	return o
}

// SetBefore adds the before to the list events by nft params
func (o *ListEventsByNftParams) SetBefore(before *float64) {
	o.Before = before
}

// WithChain adds the chain to the list events by nft params
func (o *ListEventsByNftParams) WithChain(chain string) *ListEventsByNftParams {
	o.SetChain(chain)
	return o
}

// SetChain adds the chain to the list events by nft params
func (o *ListEventsByNftParams) SetChain(chain string) {
	o.Chain = chain
}

// WithEventType adds the eventType to the list events by nft params
func (o *ListEventsByNftParams) WithEventType(eventType []string) *ListEventsByNftParams {
	o.SetEventType(eventType)
	return o
}

// SetEventType adds the eventType to the list events by nft params
func (o *ListEventsByNftParams) SetEventType(eventType []string) {
	o.EventType = eventType
}

// WithIdentifier adds the identifier to the list events by nft params
func (o *ListEventsByNftParams) WithIdentifier(identifier string) *ListEventsByNftParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the list events by nft params
func (o *ListEventsByNftParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WithNext adds the next to the list events by nft params
func (o *ListEventsByNftParams) WithNext(next *string) *ListEventsByNftParams {
	o.SetNext(next)
	return o
}

// SetNext adds the next to the list events by nft params
func (o *ListEventsByNftParams) SetNext(next *string) {
	o.Next = next
}

// WriteToRequest writes these params to a swagger request
func (o *ListEventsByNftParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param address
	if err := r.SetPathParam("address", o.Address); err != nil {
		return err
	}

	if o.After != nil {

		// query param after
		var qrAfter float64

		if o.After != nil {
			qrAfter = *o.After
		}
		qAfter := swag.FormatFloat64(qrAfter)
		if qAfter != "" {

			if err := r.SetQueryParam("after", qAfter); err != nil {
				return err
			}
		}
	}

	if o.Before != nil {

		// query param before
		var qrBefore float64

		if o.Before != nil {
			qrBefore = *o.Before
		}
		qBefore := swag.FormatFloat64(qrBefore)
		if qBefore != "" {

			if err := r.SetQueryParam("before", qBefore); err != nil {
				return err
			}
		}
	}

	// path param chain
	if err := r.SetPathParam("chain", o.Chain); err != nil {
		return err
	}

	if o.EventType != nil {

		// binding items for event_type
		joinedEventType := o.bindParamEventType(reg)

		// query array param event_type
		if err := r.SetQueryParam("event_type", joinedEventType...); err != nil {
			return err
		}
	}

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
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

// bindParamListEventsByNft binds the parameter event_type
func (o *ListEventsByNftParams) bindParamEventType(formats strfmt.Registry) []string {
	eventTypeIR := o.EventType

	var eventTypeIC []string
	for _, eventTypeIIR := range eventTypeIR { // explode []string

		eventTypeIIV := eventTypeIIR // string as string
		eventTypeIC = append(eventTypeIC, eventTypeIIV)
	}

	// items.CollectionFormat: "multi"
	eventTypeIS := swag.JoinByFormat(eventTypeIC, "multi")

	return eventTypeIS
}
