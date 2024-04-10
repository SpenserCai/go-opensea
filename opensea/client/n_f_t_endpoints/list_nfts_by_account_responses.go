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

// ListNftsByAccountReader is a Reader for the ListNftsByAccount structure.
type ListNftsByAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNftsByAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListNftsByAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListNftsByAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListNftsByAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListNftsByAccountOK creates a ListNftsByAccountOK with default headers values
func NewListNftsByAccountOK() *ListNftsByAccountOK {
	return &ListNftsByAccountOK{}
}

/*
ListNftsByAccountOK describes a response with status code 200, with default header values.

ListNftsByAccountOK list nfts by account o k
*/
type ListNftsByAccountOK struct {
	Payload *models.ListNftsResponse
}

// IsSuccess returns true when this list nfts by account o k response has a 2xx status code
func (o *ListNftsByAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list nfts by account o k response has a 3xx status code
func (o *ListNftsByAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list nfts by account o k response has a 4xx status code
func (o *ListNftsByAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list nfts by account o k response has a 5xx status code
func (o *ListNftsByAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list nfts by account o k response a status code equal to that given
func (o *ListNftsByAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list nfts by account o k response
func (o *ListNftsByAccountOK) Code() int {
	return 200
}

func (o *ListNftsByAccountOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/chain/{chain}/account/{address}/nfts][%d] listNftsByAccountOK  %+v", 200, o.Payload)
}

func (o *ListNftsByAccountOK) String() string {
	return fmt.Sprintf("[GET /api/v2/chain/{chain}/account/{address}/nfts][%d] listNftsByAccountOK  %+v", 200, o.Payload)
}

func (o *ListNftsByAccountOK) GetPayload() *models.ListNftsResponse {
	return o.Payload
}

func (o *ListNftsByAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListNftsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNftsByAccountBadRequest creates a ListNftsByAccountBadRequest with default headers values
func NewListNftsByAccountBadRequest() *ListNftsByAccountBadRequest {
	return &ListNftsByAccountBadRequest{}
}

/*
ListNftsByAccountBadRequest describes a response with status code 400, with default header values.

For error reasons, review the response data.
*/
type ListNftsByAccountBadRequest struct {
}

// IsSuccess returns true when this list nfts by account bad request response has a 2xx status code
func (o *ListNftsByAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list nfts by account bad request response has a 3xx status code
func (o *ListNftsByAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list nfts by account bad request response has a 4xx status code
func (o *ListNftsByAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list nfts by account bad request response has a 5xx status code
func (o *ListNftsByAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list nfts by account bad request response a status code equal to that given
func (o *ListNftsByAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list nfts by account bad request response
func (o *ListNftsByAccountBadRequest) Code() int {
	return 400
}

func (o *ListNftsByAccountBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/chain/{chain}/account/{address}/nfts][%d] listNftsByAccountBadRequest ", 400)
}

func (o *ListNftsByAccountBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/chain/{chain}/account/{address}/nfts][%d] listNftsByAccountBadRequest ", 400)
}

func (o *ListNftsByAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListNftsByAccountInternalServerError creates a ListNftsByAccountInternalServerError with default headers values
func NewListNftsByAccountInternalServerError() *ListNftsByAccountInternalServerError {
	return &ListNftsByAccountInternalServerError{}
}

/*
ListNftsByAccountInternalServerError describes a response with status code 500, with default header values.

Internal server error. Please open a support ticket so OpenSea can investigate.
*/
type ListNftsByAccountInternalServerError struct {
}

// IsSuccess returns true when this list nfts by account internal server error response has a 2xx status code
func (o *ListNftsByAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list nfts by account internal server error response has a 3xx status code
func (o *ListNftsByAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list nfts by account internal server error response has a 4xx status code
func (o *ListNftsByAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list nfts by account internal server error response has a 5xx status code
func (o *ListNftsByAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list nfts by account internal server error response a status code equal to that given
func (o *ListNftsByAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list nfts by account internal server error response
func (o *ListNftsByAccountInternalServerError) Code() int {
	return 500
}

func (o *ListNftsByAccountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/chain/{chain}/account/{address}/nfts][%d] listNftsByAccountInternalServerError ", 500)
}

func (o *ListNftsByAccountInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/chain/{chain}/account/{address}/nfts][%d] listNftsByAccountInternalServerError ", 500)
}

func (o *ListNftsByAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
