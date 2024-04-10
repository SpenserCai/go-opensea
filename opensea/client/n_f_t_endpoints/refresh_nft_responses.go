// Code generated by go-swagger; DO NOT EDIT.

package n_f_t_endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RefreshNftReader is a Reader for the RefreshNft structure.
type RefreshNftReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RefreshNftReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRefreshNftOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRefreshNftBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRefreshNftInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRefreshNftOK creates a RefreshNftOK with default headers values
func NewRefreshNftOK() *RefreshNftOK {
	return &RefreshNftOK{}
}

/*
RefreshNftOK describes a response with status code 200, with default header values.

Metadata has been succesfully queued for refresh.
*/
type RefreshNftOK struct {
}

// IsSuccess returns true when this refresh nft o k response has a 2xx status code
func (o *RefreshNftOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this refresh nft o k response has a 3xx status code
func (o *RefreshNftOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this refresh nft o k response has a 4xx status code
func (o *RefreshNftOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this refresh nft o k response has a 5xx status code
func (o *RefreshNftOK) IsServerError() bool {
	return false
}

// IsCode returns true when this refresh nft o k response a status code equal to that given
func (o *RefreshNftOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the refresh nft o k response
func (o *RefreshNftOK) Code() int {
	return 200
}

func (o *RefreshNftOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/chain/{chain}/contract/{address}/nfts/{identifier}/refresh][%d] refreshNftOK ", 200)
}

func (o *RefreshNftOK) String() string {
	return fmt.Sprintf("[POST /api/v2/chain/{chain}/contract/{address}/nfts/{identifier}/refresh][%d] refreshNftOK ", 200)
}

func (o *RefreshNftOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRefreshNftBadRequest creates a RefreshNftBadRequest with default headers values
func NewRefreshNftBadRequest() *RefreshNftBadRequest {
	return &RefreshNftBadRequest{}
}

/*
RefreshNftBadRequest describes a response with status code 400, with default header values.

For error reasons, review the response data.
*/
type RefreshNftBadRequest struct {
}

// IsSuccess returns true when this refresh nft bad request response has a 2xx status code
func (o *RefreshNftBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this refresh nft bad request response has a 3xx status code
func (o *RefreshNftBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this refresh nft bad request response has a 4xx status code
func (o *RefreshNftBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this refresh nft bad request response has a 5xx status code
func (o *RefreshNftBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this refresh nft bad request response a status code equal to that given
func (o *RefreshNftBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the refresh nft bad request response
func (o *RefreshNftBadRequest) Code() int {
	return 400
}

func (o *RefreshNftBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/chain/{chain}/contract/{address}/nfts/{identifier}/refresh][%d] refreshNftBadRequest ", 400)
}

func (o *RefreshNftBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/chain/{chain}/contract/{address}/nfts/{identifier}/refresh][%d] refreshNftBadRequest ", 400)
}

func (o *RefreshNftBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRefreshNftInternalServerError creates a RefreshNftInternalServerError with default headers values
func NewRefreshNftInternalServerError() *RefreshNftInternalServerError {
	return &RefreshNftInternalServerError{}
}

/*
RefreshNftInternalServerError describes a response with status code 500, with default header values.

Internal server error. Please open a support ticket so OpenSea can investigate.
*/
type RefreshNftInternalServerError struct {
}

// IsSuccess returns true when this refresh nft internal server error response has a 2xx status code
func (o *RefreshNftInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this refresh nft internal server error response has a 3xx status code
func (o *RefreshNftInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this refresh nft internal server error response has a 4xx status code
func (o *RefreshNftInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this refresh nft internal server error response has a 5xx status code
func (o *RefreshNftInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this refresh nft internal server error response a status code equal to that given
func (o *RefreshNftInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the refresh nft internal server error response
func (o *RefreshNftInternalServerError) Code() int {
	return 500
}

func (o *RefreshNftInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/chain/{chain}/contract/{address}/nfts/{identifier}/refresh][%d] refreshNftInternalServerError ", 500)
}

func (o *RefreshNftInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/chain/{chain}/contract/{address}/nfts/{identifier}/refresh][%d] refreshNftInternalServerError ", 500)
}

func (o *RefreshNftInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}