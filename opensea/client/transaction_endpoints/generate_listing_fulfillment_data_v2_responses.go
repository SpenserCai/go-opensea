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

// GenerateListingFulfillmentDataV2Reader is a Reader for the GenerateListingFulfillmentDataV2 structure.
type GenerateListingFulfillmentDataV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateListingFulfillmentDataV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateListingFulfillmentDataV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGenerateListingFulfillmentDataV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGenerateListingFulfillmentDataV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGenerateListingFulfillmentDataV2OK creates a GenerateListingFulfillmentDataV2OK with default headers values
func NewGenerateListingFulfillmentDataV2OK() *GenerateListingFulfillmentDataV2OK {
	return &GenerateListingFulfillmentDataV2OK{}
}

/*
GenerateListingFulfillmentDataV2OK describes a response with status code 200, with default header values.

GenerateListingFulfillmentDataV2OK generate listing fulfillment data v2 o k
*/
type GenerateListingFulfillmentDataV2OK struct {
	Payload *models.FulfillmentOutput
}

// IsSuccess returns true when this generate listing fulfillment data v2 o k response has a 2xx status code
func (o *GenerateListingFulfillmentDataV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this generate listing fulfillment data v2 o k response has a 3xx status code
func (o *GenerateListingFulfillmentDataV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate listing fulfillment data v2 o k response has a 4xx status code
func (o *GenerateListingFulfillmentDataV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate listing fulfillment data v2 o k response has a 5xx status code
func (o *GenerateListingFulfillmentDataV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this generate listing fulfillment data v2 o k response a status code equal to that given
func (o *GenerateListingFulfillmentDataV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the generate listing fulfillment data v2 o k response
func (o *GenerateListingFulfillmentDataV2OK) Code() int {
	return 200
}

func (o *GenerateListingFulfillmentDataV2OK) Error() string {
	return fmt.Sprintf("[POST /api/v2/listings/fulfillment_data][%d] generateListingFulfillmentDataV2OK  %+v", 200, o.Payload)
}

func (o *GenerateListingFulfillmentDataV2OK) String() string {
	return fmt.Sprintf("[POST /api/v2/listings/fulfillment_data][%d] generateListingFulfillmentDataV2OK  %+v", 200, o.Payload)
}

func (o *GenerateListingFulfillmentDataV2OK) GetPayload() *models.FulfillmentOutput {
	return o.Payload
}

func (o *GenerateListingFulfillmentDataV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FulfillmentOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateListingFulfillmentDataV2BadRequest creates a GenerateListingFulfillmentDataV2BadRequest with default headers values
func NewGenerateListingFulfillmentDataV2BadRequest() *GenerateListingFulfillmentDataV2BadRequest {
	return &GenerateListingFulfillmentDataV2BadRequest{}
}

/*
	GenerateListingFulfillmentDataV2BadRequest describes a response with status code 400, with default header values.

	The request is invalid

The order_hash does not exist
The chain is not an EVM Chain
The protocol_address is not a supported Seaport contract
For other error reasons, see the response data.
*/
type GenerateListingFulfillmentDataV2BadRequest struct {
}

// IsSuccess returns true when this generate listing fulfillment data v2 bad request response has a 2xx status code
func (o *GenerateListingFulfillmentDataV2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate listing fulfillment data v2 bad request response has a 3xx status code
func (o *GenerateListingFulfillmentDataV2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate listing fulfillment data v2 bad request response has a 4xx status code
func (o *GenerateListingFulfillmentDataV2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate listing fulfillment data v2 bad request response has a 5xx status code
func (o *GenerateListingFulfillmentDataV2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this generate listing fulfillment data v2 bad request response a status code equal to that given
func (o *GenerateListingFulfillmentDataV2BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the generate listing fulfillment data v2 bad request response
func (o *GenerateListingFulfillmentDataV2BadRequest) Code() int {
	return 400
}

func (o *GenerateListingFulfillmentDataV2BadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/listings/fulfillment_data][%d] generateListingFulfillmentDataV2BadRequest ", 400)
}

func (o *GenerateListingFulfillmentDataV2BadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/listings/fulfillment_data][%d] generateListingFulfillmentDataV2BadRequest ", 400)
}

func (o *GenerateListingFulfillmentDataV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGenerateListingFulfillmentDataV2InternalServerError creates a GenerateListingFulfillmentDataV2InternalServerError with default headers values
func NewGenerateListingFulfillmentDataV2InternalServerError() *GenerateListingFulfillmentDataV2InternalServerError {
	return &GenerateListingFulfillmentDataV2InternalServerError{}
}

/*
GenerateListingFulfillmentDataV2InternalServerError describes a response with status code 500, with default header values.

Internal server error. Please open a support ticket so OpenSea can investigate.
*/
type GenerateListingFulfillmentDataV2InternalServerError struct {
}

// IsSuccess returns true when this generate listing fulfillment data v2 internal server error response has a 2xx status code
func (o *GenerateListingFulfillmentDataV2InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate listing fulfillment data v2 internal server error response has a 3xx status code
func (o *GenerateListingFulfillmentDataV2InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate listing fulfillment data v2 internal server error response has a 4xx status code
func (o *GenerateListingFulfillmentDataV2InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate listing fulfillment data v2 internal server error response has a 5xx status code
func (o *GenerateListingFulfillmentDataV2InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this generate listing fulfillment data v2 internal server error response a status code equal to that given
func (o *GenerateListingFulfillmentDataV2InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the generate listing fulfillment data v2 internal server error response
func (o *GenerateListingFulfillmentDataV2InternalServerError) Code() int {
	return 500
}

func (o *GenerateListingFulfillmentDataV2InternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/listings/fulfillment_data][%d] generateListingFulfillmentDataV2InternalServerError ", 500)
}

func (o *GenerateListingFulfillmentDataV2InternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/listings/fulfillment_data][%d] generateListingFulfillmentDataV2InternalServerError ", 500)
}

func (o *GenerateListingFulfillmentDataV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
