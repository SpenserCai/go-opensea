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

// GetOrderReader is a Reader for the GetOrder structure.
type GetOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrderOK creates a GetOrderOK with default headers values
func NewGetOrderOK() *GetOrderOK {
	return &GetOrderOK{}
}

/*
GetOrderOK describes a response with status code 200, with default header values.

GetOrderOK get order o k
*/
type GetOrderOK struct {
	Payload *models.GetOrderResult
}

// IsSuccess returns true when this get order o k response has a 2xx status code
func (o *GetOrderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get order o k response has a 3xx status code
func (o *GetOrderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order o k response has a 4xx status code
func (o *GetOrderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get order o k response has a 5xx status code
func (o *GetOrderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get order o k response a status code equal to that given
func (o *GetOrderOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get order o k response
func (o *GetOrderOK) Code() int {
	return 200
}

func (o *GetOrderOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/orders/chain/{chain}/protocol/{protocol_address}/{order_hash}][%d] getOrderOK  %+v", 200, o.Payload)
}

func (o *GetOrderOK) String() string {
	return fmt.Sprintf("[GET /api/v2/orders/chain/{chain}/protocol/{protocol_address}/{order_hash}][%d] getOrderOK  %+v", 200, o.Payload)
}

func (o *GetOrderOK) GetPayload() *models.GetOrderResult {
	return o.Payload
}

func (o *GetOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetOrderResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderBadRequest creates a GetOrderBadRequest with default headers values
func NewGetOrderBadRequest() *GetOrderBadRequest {
	return &GetOrderBadRequest{}
}

/*
GetOrderBadRequest describes a response with status code 400, with default header values.

For error reasons, review the response data.
*/
type GetOrderBadRequest struct {
}

// IsSuccess returns true when this get order bad request response has a 2xx status code
func (o *GetOrderBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get order bad request response has a 3xx status code
func (o *GetOrderBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order bad request response has a 4xx status code
func (o *GetOrderBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get order bad request response has a 5xx status code
func (o *GetOrderBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get order bad request response a status code equal to that given
func (o *GetOrderBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get order bad request response
func (o *GetOrderBadRequest) Code() int {
	return 400
}

func (o *GetOrderBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/orders/chain/{chain}/protocol/{protocol_address}/{order_hash}][%d] getOrderBadRequest ", 400)
}

func (o *GetOrderBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/orders/chain/{chain}/protocol/{protocol_address}/{order_hash}][%d] getOrderBadRequest ", 400)
}

func (o *GetOrderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrderInternalServerError creates a GetOrderInternalServerError with default headers values
func NewGetOrderInternalServerError() *GetOrderInternalServerError {
	return &GetOrderInternalServerError{}
}

/*
GetOrderInternalServerError describes a response with status code 500, with default header values.

Internal server error. Please open a support ticket so OpenSea can investigate.
*/
type GetOrderInternalServerError struct {
}

// IsSuccess returns true when this get order internal server error response has a 2xx status code
func (o *GetOrderInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get order internal server error response has a 3xx status code
func (o *GetOrderInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order internal server error response has a 4xx status code
func (o *GetOrderInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get order internal server error response has a 5xx status code
func (o *GetOrderInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get order internal server error response a status code equal to that given
func (o *GetOrderInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get order internal server error response
func (o *GetOrderInternalServerError) Code() int {
	return 500
}

func (o *GetOrderInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/orders/chain/{chain}/protocol/{protocol_address}/{order_hash}][%d] getOrderInternalServerError ", 500)
}

func (o *GetOrderInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/orders/chain/{chain}/protocol/{protocol_address}/{order_hash}][%d] getOrderInternalServerError ", 500)
}

func (o *GetOrderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
