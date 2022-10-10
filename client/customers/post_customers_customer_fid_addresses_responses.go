// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PostCustomersCustomerFidAddressesReader is a Reader for the PostCustomersCustomerFidAddresses structure.
type PostCustomersCustomerFidAddressesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomersCustomerFidAddressesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCustomersCustomerFidAddressesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostCustomersCustomerFidAddressesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCustomersCustomerFidAddressesOK creates a PostCustomersCustomerFidAddressesOK with default headers values
func NewPostCustomersCustomerFidAddressesOK() *PostCustomersCustomerFidAddressesOK {
	return &PostCustomersCustomerFidAddressesOK{}
}

/*
PostCustomersCustomerFidAddressesOK describes a response with status code 200, with default header values.

Address Added
*/
type PostCustomersCustomerFidAddressesOK struct {
}

// IsSuccess returns true when this post customers customer fid addresses o k response has a 2xx status code
func (o *PostCustomersCustomerFidAddressesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post customers customer fid addresses o k response has a 3xx status code
func (o *PostCustomersCustomerFidAddressesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post customers customer fid addresses o k response has a 4xx status code
func (o *PostCustomersCustomerFidAddressesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post customers customer fid addresses o k response has a 5xx status code
func (o *PostCustomersCustomerFidAddressesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post customers customer fid addresses o k response a status code equal to that given
func (o *PostCustomersCustomerFidAddressesOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostCustomersCustomerFidAddressesOK) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/addresses][%d] postCustomersCustomerFidAddressesOK ", 200)
}

func (o *PostCustomersCustomerFidAddressesOK) String() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/addresses][%d] postCustomersCustomerFidAddressesOK ", 200)
}

func (o *PostCustomersCustomerFidAddressesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomersCustomerFidAddressesDefault creates a PostCustomersCustomerFidAddressesDefault with default headers values
func NewPostCustomersCustomerFidAddressesDefault(code int) *PostCustomersCustomerFidAddressesDefault {
	return &PostCustomersCustomerFidAddressesDefault{
		_statusCode: code,
	}
}

/*
PostCustomersCustomerFidAddressesDefault describes a response with status code -1, with default header values.

Error
*/
type PostCustomersCustomerFidAddressesDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post customers customer fid addresses default response
func (o *PostCustomersCustomerFidAddressesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this post customers customer fid addresses default response has a 2xx status code
func (o *PostCustomersCustomerFidAddressesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post customers customer fid addresses default response has a 3xx status code
func (o *PostCustomersCustomerFidAddressesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post customers customer fid addresses default response has a 4xx status code
func (o *PostCustomersCustomerFidAddressesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post customers customer fid addresses default response has a 5xx status code
func (o *PostCustomersCustomerFidAddressesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post customers customer fid addresses default response a status code equal to that given
func (o *PostCustomersCustomerFidAddressesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PostCustomersCustomerFidAddressesDefault) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/addresses][%d] PostCustomersCustomerFidAddresses default  %+v", o._statusCode, o.Payload)
}

func (o *PostCustomersCustomerFidAddressesDefault) String() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/addresses][%d] PostCustomersCustomerFidAddresses default  %+v", o._statusCode, o.Payload)
}

func (o *PostCustomersCustomerFidAddressesDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostCustomersCustomerFidAddressesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
