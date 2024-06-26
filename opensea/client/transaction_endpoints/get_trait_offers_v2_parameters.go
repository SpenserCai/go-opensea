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

// NewGetTraitOffersV2Params creates a new GetTraitOffersV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTraitOffersV2Params() *GetTraitOffersV2Params {
	return &GetTraitOffersV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTraitOffersV2ParamsWithTimeout creates a new GetTraitOffersV2Params object
// with the ability to set a timeout on a request.
func NewGetTraitOffersV2ParamsWithTimeout(timeout time.Duration) *GetTraitOffersV2Params {
	return &GetTraitOffersV2Params{
		timeout: timeout,
	}
}

// NewGetTraitOffersV2ParamsWithContext creates a new GetTraitOffersV2Params object
// with the ability to set a context for a request.
func NewGetTraitOffersV2ParamsWithContext(ctx context.Context) *GetTraitOffersV2Params {
	return &GetTraitOffersV2Params{
		Context: ctx,
	}
}

// NewGetTraitOffersV2ParamsWithHTTPClient creates a new GetTraitOffersV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetTraitOffersV2ParamsWithHTTPClient(client *http.Client) *GetTraitOffersV2Params {
	return &GetTraitOffersV2Params{
		HTTPClient: client,
	}
}

/*
GetTraitOffersV2Params contains all the parameters to send to the API endpoint

	for the get trait offers v2 operation.

	Typically these are written to a http.Request.
*/
type GetTraitOffersV2Params struct {

	/* CollectionSlug.

	   Unique string to identify a collection on OpenSea. This can be found by visiting the collection on the OpenSea website and noting the last path parameter.
	*/
	CollectionSlug string

	/* FloatValue.

	   The value of the trait (e.g. `0.5`). This is only used for decimal-based numeric traits to ensure it is parsed correctly.

	   Format: float
	*/
	FloatValue *float32

	/* IntValue.

	   The value of the trait (e.g. `10`). This is only used for numeric traits to ensure it is parsed correctly.
	*/
	IntValue *int64

	/* Type.

	   The name of the trait (e.g. 'Background')
	*/
	Type *string

	/* Value.

	   The value of the trait (e.g. 'Red')
	*/
	Value *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get trait offers v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTraitOffersV2Params) WithDefaults() *GetTraitOffersV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get trait offers v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTraitOffersV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get trait offers v2 params
func (o *GetTraitOffersV2Params) WithTimeout(timeout time.Duration) *GetTraitOffersV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get trait offers v2 params
func (o *GetTraitOffersV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get trait offers v2 params
func (o *GetTraitOffersV2Params) WithContext(ctx context.Context) *GetTraitOffersV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get trait offers v2 params
func (o *GetTraitOffersV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get trait offers v2 params
func (o *GetTraitOffersV2Params) WithHTTPClient(client *http.Client) *GetTraitOffersV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get trait offers v2 params
func (o *GetTraitOffersV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCollectionSlug adds the collectionSlug to the get trait offers v2 params
func (o *GetTraitOffersV2Params) WithCollectionSlug(collectionSlug string) *GetTraitOffersV2Params {
	o.SetCollectionSlug(collectionSlug)
	return o
}

// SetCollectionSlug adds the collectionSlug to the get trait offers v2 params
func (o *GetTraitOffersV2Params) SetCollectionSlug(collectionSlug string) {
	o.CollectionSlug = collectionSlug
}

// WithFloatValue adds the floatValue to the get trait offers v2 params
func (o *GetTraitOffersV2Params) WithFloatValue(floatValue *float32) *GetTraitOffersV2Params {
	o.SetFloatValue(floatValue)
	return o
}

// SetFloatValue adds the floatValue to the get trait offers v2 params
func (o *GetTraitOffersV2Params) SetFloatValue(floatValue *float32) {
	o.FloatValue = floatValue
}

// WithIntValue adds the intValue to the get trait offers v2 params
func (o *GetTraitOffersV2Params) WithIntValue(intValue *int64) *GetTraitOffersV2Params {
	o.SetIntValue(intValue)
	return o
}

// SetIntValue adds the intValue to the get trait offers v2 params
func (o *GetTraitOffersV2Params) SetIntValue(intValue *int64) {
	o.IntValue = intValue
}

// WithType adds the typeVar to the get trait offers v2 params
func (o *GetTraitOffersV2Params) WithType(typeVar *string) *GetTraitOffersV2Params {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get trait offers v2 params
func (o *GetTraitOffersV2Params) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithValue adds the value to the get trait offers v2 params
func (o *GetTraitOffersV2Params) WithValue(value *string) *GetTraitOffersV2Params {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the get trait offers v2 params
func (o *GetTraitOffersV2Params) SetValue(value *string) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *GetTraitOffersV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param collection_slug
	if err := r.SetPathParam("collection_slug", o.CollectionSlug); err != nil {
		return err
	}

	if o.FloatValue != nil {

		// query param float_value
		var qrFloatValue float32

		if o.FloatValue != nil {
			qrFloatValue = *o.FloatValue
		}
		qFloatValue := swag.FormatFloat32(qrFloatValue)
		if qFloatValue != "" {

			if err := r.SetQueryParam("float_value", qFloatValue); err != nil {
				return err
			}
		}
	}

	if o.IntValue != nil {

		// query param int_value
		var qrIntValue int64

		if o.IntValue != nil {
			qrIntValue = *o.IntValue
		}
		qIntValue := swag.FormatInt64(qrIntValue)
		if qIntValue != "" {

			if err := r.SetQueryParam("int_value", qIntValue); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if o.Value != nil {

		// query param value
		var qrValue string

		if o.Value != nil {
			qrValue = *o.Value
		}
		qValue := qrValue
		if qValue != "" {

			if err := r.SetQueryParam("value", qValue); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
