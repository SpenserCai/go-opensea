// Code generated by go-swagger; DO NOT EDIT.

package n_f_t_endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/SpenserCai/go-opensea/opensea/models"
)

// GetContractReader is a Reader for the GetContract structure.
type GetContractReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContractReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContractOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetContractBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetContractInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetContractOK creates a GetContractOK with default headers values
func NewGetContractOK() *GetContractOK {
	return &GetContractOK{}
}

/*
GetContractOK describes a response with status code 200, with default header values.

GetContractOK get contract o k
*/
type GetContractOK struct {
	Payload *models.ListNftsResponse
}

// IsSuccess returns true when this get contract o k response has a 2xx status code
func (o *GetContractOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get contract o k response has a 3xx status code
func (o *GetContractOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contract o k response has a 4xx status code
func (o *GetContractOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contract o k response has a 5xx status code
func (o *GetContractOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get contract o k response a status code equal to that given
func (o *GetContractOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get contract o k response
func (o *GetContractOK) Code() int {
	return 200
}

func (o *GetContractOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/chain/{chain}/contract/{address}][%d] getContractOK  %+v", 200, o.Payload)
}

func (o *GetContractOK) String() string {
	return fmt.Sprintf("[GET /api/v2/chain/{chain}/contract/{address}][%d] getContractOK  %+v", 200, o.Payload)
}

func (o *GetContractOK) GetPayload() *models.ListNftsResponse {
	return o.Payload
}

func (o *GetContractOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListNftsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContractBadRequest creates a GetContractBadRequest with default headers values
func NewGetContractBadRequest() *GetContractBadRequest {
	return &GetContractBadRequest{}
}

/*
GetContractBadRequest describes a response with status code 400, with default header values.

For error reasons, review the response data.
*/
type GetContractBadRequest struct {
}

// IsSuccess returns true when this get contract bad request response has a 2xx status code
func (o *GetContractBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contract bad request response has a 3xx status code
func (o *GetContractBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contract bad request response has a 4xx status code
func (o *GetContractBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contract bad request response has a 5xx status code
func (o *GetContractBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get contract bad request response a status code equal to that given
func (o *GetContractBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get contract bad request response
func (o *GetContractBadRequest) Code() int {
	return 400
}

func (o *GetContractBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/chain/{chain}/contract/{address}][%d] getContractBadRequest ", 400)
}

func (o *GetContractBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/chain/{chain}/contract/{address}][%d] getContractBadRequest ", 400)
}

func (o *GetContractBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetContractInternalServerError creates a GetContractInternalServerError with default headers values
func NewGetContractInternalServerError() *GetContractInternalServerError {
	return &GetContractInternalServerError{}
}

/*
GetContractInternalServerError describes a response with status code 500, with default header values.

Internal server error. Please open a support ticket so OpenSea can investigate.
*/
type GetContractInternalServerError struct {
}

// IsSuccess returns true when this get contract internal server error response has a 2xx status code
func (o *GetContractInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contract internal server error response has a 3xx status code
func (o *GetContractInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contract internal server error response has a 4xx status code
func (o *GetContractInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contract internal server error response has a 5xx status code
func (o *GetContractInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get contract internal server error response a status code equal to that given
func (o *GetContractInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get contract internal server error response
func (o *GetContractInternalServerError) Code() int {
	return 500
}

func (o *GetContractInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/chain/{chain}/contract/{address}][%d] getContractInternalServerError ", 500)
}

func (o *GetContractInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/chain/{chain}/contract/{address}][%d] getContractInternalServerError ", 500)
}

func (o *GetContractInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
