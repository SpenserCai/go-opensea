// Code generated by go-swagger; DO NOT EDIT.

package transaction_endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/SpenserCai/go-opensea/opensea/models"
)

// GetCollectionOffersV2Reader is a Reader for the GetCollectionOffersV2 structure.
type GetCollectionOffersV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCollectionOffersV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCollectionOffersV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetCollectionOffersV2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCollectionOffersV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCollectionOffersV2OK creates a GetCollectionOffersV2OK with default headers values
func NewGetCollectionOffersV2OK() *GetCollectionOffersV2OK {
	return &GetCollectionOffersV2OK{}
}

/*
GetCollectionOffersV2OK describes a response with status code 200, with default header values.

GetCollectionOffersV2OK get collection offers v2 o k
*/
type GetCollectionOffersV2OK struct {
	Payload *models.OfferList
}

// IsSuccess returns true when this get collection offers v2 o k response has a 2xx status code
func (o *GetCollectionOffersV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get collection offers v2 o k response has a 3xx status code
func (o *GetCollectionOffersV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get collection offers v2 o k response has a 4xx status code
func (o *GetCollectionOffersV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get collection offers v2 o k response has a 5xx status code
func (o *GetCollectionOffersV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get collection offers v2 o k response a status code equal to that given
func (o *GetCollectionOffersV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get collection offers v2 o k response
func (o *GetCollectionOffersV2OK) Code() int {
	return 200
}

func (o *GetCollectionOffersV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/offers/collection/{collection_slug}][%d] getCollectionOffersV2OK  %+v", 200, o.Payload)
}

func (o *GetCollectionOffersV2OK) String() string {
	return fmt.Sprintf("[GET /api/v2/offers/collection/{collection_slug}][%d] getCollectionOffersV2OK  %+v", 200, o.Payload)
}

func (o *GetCollectionOffersV2OK) GetPayload() *models.OfferList {
	return o.Payload
}

func (o *GetCollectionOffersV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OfferList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCollectionOffersV2NotFound creates a GetCollectionOffersV2NotFound with default headers values
func NewGetCollectionOffersV2NotFound() *GetCollectionOffersV2NotFound {
	return &GetCollectionOffersV2NotFound{}
}

/*
GetCollectionOffersV2NotFound describes a response with status code 404, with default header values.

Returned when the collection does not exist.
*/
type GetCollectionOffersV2NotFound struct {
}

// IsSuccess returns true when this get collection offers v2 not found response has a 2xx status code
func (o *GetCollectionOffersV2NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get collection offers v2 not found response has a 3xx status code
func (o *GetCollectionOffersV2NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get collection offers v2 not found response has a 4xx status code
func (o *GetCollectionOffersV2NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get collection offers v2 not found response has a 5xx status code
func (o *GetCollectionOffersV2NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get collection offers v2 not found response a status code equal to that given
func (o *GetCollectionOffersV2NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get collection offers v2 not found response
func (o *GetCollectionOffersV2NotFound) Code() int {
	return 404
}

func (o *GetCollectionOffersV2NotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/offers/collection/{collection_slug}][%d] getCollectionOffersV2NotFound ", 404)
}

func (o *GetCollectionOffersV2NotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/offers/collection/{collection_slug}][%d] getCollectionOffersV2NotFound ", 404)
}

func (o *GetCollectionOffersV2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCollectionOffersV2InternalServerError creates a GetCollectionOffersV2InternalServerError with default headers values
func NewGetCollectionOffersV2InternalServerError() *GetCollectionOffersV2InternalServerError {
	return &GetCollectionOffersV2InternalServerError{}
}

/*
GetCollectionOffersV2InternalServerError describes a response with status code 500, with default header values.

Internal server error. Please open a support ticket so OpenSea can investigate.
*/
type GetCollectionOffersV2InternalServerError struct {
}

// IsSuccess returns true when this get collection offers v2 internal server error response has a 2xx status code
func (o *GetCollectionOffersV2InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get collection offers v2 internal server error response has a 3xx status code
func (o *GetCollectionOffersV2InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get collection offers v2 internal server error response has a 4xx status code
func (o *GetCollectionOffersV2InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get collection offers v2 internal server error response has a 5xx status code
func (o *GetCollectionOffersV2InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get collection offers v2 internal server error response a status code equal to that given
func (o *GetCollectionOffersV2InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get collection offers v2 internal server error response
func (o *GetCollectionOffersV2InternalServerError) Code() int {
	return 500
}

func (o *GetCollectionOffersV2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/offers/collection/{collection_slug}][%d] getCollectionOffersV2InternalServerError ", 500)
}

func (o *GetCollectionOffersV2InternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/offers/collection/{collection_slug}][%d] getCollectionOffersV2InternalServerError ", 500)
}

func (o *GetCollectionOffersV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
