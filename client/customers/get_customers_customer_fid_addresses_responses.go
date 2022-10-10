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

// GetCustomersCustomerFidAddressesReader is a Reader for the GetCustomersCustomerFidAddresses structure.
type GetCustomersCustomerFidAddressesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersCustomerFidAddressesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomersCustomerFidAddressesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCustomersCustomerFidAddressesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersCustomerFidAddressesOK creates a GetCustomersCustomerFidAddressesOK with default headers values
func NewGetCustomersCustomerFidAddressesOK() *GetCustomersCustomerFidAddressesOK {
	return &GetCustomersCustomerFidAddressesOK{}
}

/*
GetCustomersCustomerFidAddressesOK describes a response with status code 200, with default header values.

List of addresses
*/
type GetCustomersCustomerFidAddressesOK struct {
	Payload *models.Addresses
}

// IsSuccess returns true when this get customers customer fid addresses o k response has a 2xx status code
func (o *GetCustomersCustomerFidAddressesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get customers customer fid addresses o k response has a 3xx status code
func (o *GetCustomersCustomerFidAddressesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customers customer fid addresses o k response has a 4xx status code
func (o *GetCustomersCustomerFidAddressesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get customers customer fid addresses o k response has a 5xx status code
func (o *GetCustomersCustomerFidAddressesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get customers customer fid addresses o k response a status code equal to that given
func (o *GetCustomersCustomerFidAddressesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCustomersCustomerFidAddressesOK) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/addresses][%d] getCustomersCustomerFidAddressesOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidAddressesOK) String() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/addresses][%d] getCustomersCustomerFidAddressesOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidAddressesOK) GetPayload() *models.Addresses {
	return o.Payload
}

func (o *GetCustomersCustomerFidAddressesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Addresses)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersCustomerFidAddressesDefault creates a GetCustomersCustomerFidAddressesDefault with default headers values
func NewGetCustomersCustomerFidAddressesDefault(code int) *GetCustomersCustomerFidAddressesDefault {
	return &GetCustomersCustomerFidAddressesDefault{
		_statusCode: code,
	}
}

/*
GetCustomersCustomerFidAddressesDefault describes a response with status code -1, with default header values.

Error
*/
type GetCustomersCustomerFidAddressesDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get customers customer fid addresses default response
func (o *GetCustomersCustomerFidAddressesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get customers customer fid addresses default response has a 2xx status code
func (o *GetCustomersCustomerFidAddressesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get customers customer fid addresses default response has a 3xx status code
func (o *GetCustomersCustomerFidAddressesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get customers customer fid addresses default response has a 4xx status code
func (o *GetCustomersCustomerFidAddressesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get customers customer fid addresses default response has a 5xx status code
func (o *GetCustomersCustomerFidAddressesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get customers customer fid addresses default response a status code equal to that given
func (o *GetCustomersCustomerFidAddressesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetCustomersCustomerFidAddressesDefault) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/addresses][%d] GetCustomersCustomerFidAddresses default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidAddressesDefault) String() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/addresses][%d] GetCustomersCustomerFidAddresses default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidAddressesDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetCustomersCustomerFidAddressesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
