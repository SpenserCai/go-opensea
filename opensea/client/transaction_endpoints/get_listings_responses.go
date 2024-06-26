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

// GetListingsReader is a Reader for the GetListings structure.
type GetListingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetListingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetListingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetListingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetListingsOK creates a GetListingsOK with default headers values
func NewGetListingsOK() *GetListingsOK {
	return &GetListingsOK{}
}

/*
GetListingsOK describes a response with status code 200, with default header values.

GetListingsOK get listings o k
*/
type GetListingsOK struct {
	Payload *models.GetListingsResponse
}

// IsSuccess returns true when this get listings o k response has a 2xx status code
func (o *GetListingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get listings o k response has a 3xx status code
func (o *GetListingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listings o k response has a 4xx status code
func (o *GetListingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get listings o k response has a 5xx status code
func (o *GetListingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get listings o k response a status code equal to that given
func (o *GetListingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get listings o k response
func (o *GetListingsOK) Code() int {
	return 200
}

func (o *GetListingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/orders/{chain}/{protocol}/listings][%d] getListingsOK  %+v", 200, o.Payload)
}

func (o *GetListingsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/orders/{chain}/{protocol}/listings][%d] getListingsOK  %+v", 200, o.Payload)
}

func (o *GetListingsOK) GetPayload() *models.GetListingsResponse {
	return o.Payload
}

func (o *GetListingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetListingsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingsInternalServerError creates a GetListingsInternalServerError with default headers values
func NewGetListingsInternalServerError() *GetListingsInternalServerError {
	return &GetListingsInternalServerError{}
}

/*
GetListingsInternalServerError describes a response with status code 500, with default header values.

Internal server error. Please open a support ticket so OpenSea can investigate.
*/
type GetListingsInternalServerError struct {
}

// IsSuccess returns true when this get listings internal server error response has a 2xx status code
func (o *GetListingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get listings internal server error response has a 3xx status code
func (o *GetListingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listings internal server error response has a 4xx status code
func (o *GetListingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get listings internal server error response has a 5xx status code
func (o *GetListingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get listings internal server error response a status code equal to that given
func (o *GetListingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get listings internal server error response
func (o *GetListingsInternalServerError) Code() int {
	return 500
}

func (o *GetListingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/orders/{chain}/{protocol}/listings][%d] getListingsInternalServerError ", 500)
}

func (o *GetListingsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/orders/{chain}/{protocol}/listings][%d] getListingsInternalServerError ", 500)
}

func (o *GetListingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
