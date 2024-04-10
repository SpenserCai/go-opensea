// Code generated by go-swagger; DO NOT EDIT.

package analytics_endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/SpenserCai/go-opensea/opensea/models"
)

// ListEventsByCollectionReader is a Reader for the ListEventsByCollection structure.
type ListEventsByCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEventsByCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEventsByCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListEventsByCollectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListEventsByCollectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListEventsByCollectionOK creates a ListEventsByCollectionOK with default headers values
func NewListEventsByCollectionOK() *ListEventsByCollectionOK {
	return &ListEventsByCollectionOK{}
}

/*
ListEventsByCollectionOK describes a response with status code 200, with default header values.

ListEventsByCollectionOK list events by collection o k
*/
type ListEventsByCollectionOK struct {
	Payload *models.ListEventsResponse
}

// IsSuccess returns true when this list events by collection o k response has a 2xx status code
func (o *ListEventsByCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list events by collection o k response has a 3xx status code
func (o *ListEventsByCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list events by collection o k response has a 4xx status code
func (o *ListEventsByCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list events by collection o k response has a 5xx status code
func (o *ListEventsByCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list events by collection o k response a status code equal to that given
func (o *ListEventsByCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list events by collection o k response
func (o *ListEventsByCollectionOK) Code() int {
	return 200
}

func (o *ListEventsByCollectionOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/events/collection/{collection_slug}][%d] listEventsByCollectionOK  %+v", 200, o.Payload)
}

func (o *ListEventsByCollectionOK) String() string {
	return fmt.Sprintf("[GET /api/v2/events/collection/{collection_slug}][%d] listEventsByCollectionOK  %+v", 200, o.Payload)
}

func (o *ListEventsByCollectionOK) GetPayload() *models.ListEventsResponse {
	return o.Payload
}

func (o *ListEventsByCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListEventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEventsByCollectionBadRequest creates a ListEventsByCollectionBadRequest with default headers values
func NewListEventsByCollectionBadRequest() *ListEventsByCollectionBadRequest {
	return &ListEventsByCollectionBadRequest{}
}

/*
ListEventsByCollectionBadRequest describes a response with status code 400, with default header values.

For error reasons, review the response data.
*/
type ListEventsByCollectionBadRequest struct {
}

// IsSuccess returns true when this list events by collection bad request response has a 2xx status code
func (o *ListEventsByCollectionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list events by collection bad request response has a 3xx status code
func (o *ListEventsByCollectionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list events by collection bad request response has a 4xx status code
func (o *ListEventsByCollectionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list events by collection bad request response has a 5xx status code
func (o *ListEventsByCollectionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list events by collection bad request response a status code equal to that given
func (o *ListEventsByCollectionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list events by collection bad request response
func (o *ListEventsByCollectionBadRequest) Code() int {
	return 400
}

func (o *ListEventsByCollectionBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/events/collection/{collection_slug}][%d] listEventsByCollectionBadRequest ", 400)
}

func (o *ListEventsByCollectionBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/events/collection/{collection_slug}][%d] listEventsByCollectionBadRequest ", 400)
}

func (o *ListEventsByCollectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEventsByCollectionInternalServerError creates a ListEventsByCollectionInternalServerError with default headers values
func NewListEventsByCollectionInternalServerError() *ListEventsByCollectionInternalServerError {
	return &ListEventsByCollectionInternalServerError{}
}

/*
ListEventsByCollectionInternalServerError describes a response with status code 500, with default header values.

Internal server error. Please open a support ticket so OpenSea can investigate.
*/
type ListEventsByCollectionInternalServerError struct {
}

// IsSuccess returns true when this list events by collection internal server error response has a 2xx status code
func (o *ListEventsByCollectionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list events by collection internal server error response has a 3xx status code
func (o *ListEventsByCollectionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list events by collection internal server error response has a 4xx status code
func (o *ListEventsByCollectionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list events by collection internal server error response has a 5xx status code
func (o *ListEventsByCollectionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list events by collection internal server error response a status code equal to that given
func (o *ListEventsByCollectionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list events by collection internal server error response
func (o *ListEventsByCollectionInternalServerError) Code() int {
	return 500
}

func (o *ListEventsByCollectionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/events/collection/{collection_slug}][%d] listEventsByCollectionInternalServerError ", 500)
}

func (o *ListEventsByCollectionInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/events/collection/{collection_slug}][%d] listEventsByCollectionInternalServerError ", 500)
}

func (o *ListEventsByCollectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
