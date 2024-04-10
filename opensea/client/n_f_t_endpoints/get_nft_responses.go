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

// GetNftReader is a Reader for the GetNft structure.
type GetNftReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNftReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNftOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNftBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNftInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNftOK creates a GetNftOK with default headers values
func NewGetNftOK() *GetNftOK {
	return &GetNftOK{}
}

/*
GetNftOK describes a response with status code 200, with default header values.

GetNftOK get nft o k
*/
type GetNftOK struct {
	Payload *models.GetNftResponse
}

// IsSuccess returns true when this get nft o k response has a 2xx status code
func (o *GetNftOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get nft o k response has a 3xx status code
func (o *GetNftOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get nft o k response has a 4xx status code
func (o *GetNftOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get nft o k response has a 5xx status code
func (o *GetNftOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get nft o k response a status code equal to that given
func (o *GetNftOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get nft o k response
func (o *GetNftOK) Code() int {
	return 200
}

func (o *GetNftOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/chain/{chain}/contract/{address}/nfts/{identifier}][%d] getNftOK  %+v", 200, o.Payload)
}

func (o *GetNftOK) String() string {
	return fmt.Sprintf("[GET /api/v2/chain/{chain}/contract/{address}/nfts/{identifier}][%d] getNftOK  %+v", 200, o.Payload)
}

func (o *GetNftOK) GetPayload() *models.GetNftResponse {
	return o.Payload
}

func (o *GetNftOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetNftResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNftBadRequest creates a GetNftBadRequest with default headers values
func NewGetNftBadRequest() *GetNftBadRequest {
	return &GetNftBadRequest{}
}

/*
GetNftBadRequest describes a response with status code 400, with default header values.

For error reasons, review the response data.
*/
type GetNftBadRequest struct {
}

// IsSuccess returns true when this get nft bad request response has a 2xx status code
func (o *GetNftBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get nft bad request response has a 3xx status code
func (o *GetNftBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get nft bad request response has a 4xx status code
func (o *GetNftBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get nft bad request response has a 5xx status code
func (o *GetNftBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get nft bad request response a status code equal to that given
func (o *GetNftBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get nft bad request response
func (o *GetNftBadRequest) Code() int {
	return 400
}

func (o *GetNftBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/chain/{chain}/contract/{address}/nfts/{identifier}][%d] getNftBadRequest ", 400)
}

func (o *GetNftBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/chain/{chain}/contract/{address}/nfts/{identifier}][%d] getNftBadRequest ", 400)
}

func (o *GetNftBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNftInternalServerError creates a GetNftInternalServerError with default headers values
func NewGetNftInternalServerError() *GetNftInternalServerError {
	return &GetNftInternalServerError{}
}

/*
GetNftInternalServerError describes a response with status code 500, with default header values.

Internal server error. Please open a support ticket so OpenSea can investigate.
*/
type GetNftInternalServerError struct {
}

// IsSuccess returns true when this get nft internal server error response has a 2xx status code
func (o *GetNftInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get nft internal server error response has a 3xx status code
func (o *GetNftInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get nft internal server error response has a 4xx status code
func (o *GetNftInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get nft internal server error response has a 5xx status code
func (o *GetNftInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get nft internal server error response a status code equal to that given
func (o *GetNftInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get nft internal server error response
func (o *GetNftInternalServerError) Code() int {
	return 500
}

func (o *GetNftInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/chain/{chain}/contract/{address}/nfts/{identifier}][%d] getNftInternalServerError ", 500)
}

func (o *GetNftInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/chain/{chain}/contract/{address}/nfts/{identifier}][%d] getNftInternalServerError ", 500)
}

func (o *GetNftInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}