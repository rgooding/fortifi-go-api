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

// PutCustomersCustomerFidUpgradePaymentAccountReader is a Reader for the PutCustomersCustomerFidUpgradePaymentAccount structure.
type PutCustomersCustomerFidUpgradePaymentAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomersCustomerFidUpgradePaymentAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCustomersCustomerFidUpgradePaymentAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutCustomersCustomerFidUpgradePaymentAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomersCustomerFidUpgradePaymentAccountOK creates a PutCustomersCustomerFidUpgradePaymentAccountOK with default headers values
func NewPutCustomersCustomerFidUpgradePaymentAccountOK() *PutCustomersCustomerFidUpgradePaymentAccountOK {
	return &PutCustomersCustomerFidUpgradePaymentAccountOK{}
}

/* PutCustomersCustomerFidUpgradePaymentAccountOK describes a response with status code 200, with default header values.

Customer queued for upgrade or already upgraded
*/
type PutCustomersCustomerFidUpgradePaymentAccountOK struct {
}

func (o *PutCustomersCustomerFidUpgradePaymentAccountOK) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/upgradePaymentAccount][%d] putCustomersCustomerFidUpgradePaymentAccountOK ", 200)
}

func (o *PutCustomersCustomerFidUpgradePaymentAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomersCustomerFidUpgradePaymentAccountDefault creates a PutCustomersCustomerFidUpgradePaymentAccountDefault with default headers values
func NewPutCustomersCustomerFidUpgradePaymentAccountDefault(code int) *PutCustomersCustomerFidUpgradePaymentAccountDefault {
	return &PutCustomersCustomerFidUpgradePaymentAccountDefault{
		_statusCode: code,
	}
}

/* PutCustomersCustomerFidUpgradePaymentAccountDefault describes a response with status code -1, with default header values.

Error
*/
type PutCustomersCustomerFidUpgradePaymentAccountDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put customers customer fid upgrade payment account default response
func (o *PutCustomersCustomerFidUpgradePaymentAccountDefault) Code() int {
	return o._statusCode
}

func (o *PutCustomersCustomerFidUpgradePaymentAccountDefault) Error() string {
	return fmt.Sprintf("[PUT /customers/{customerFid}/upgradePaymentAccount][%d] PutCustomersCustomerFidUpgradePaymentAccount default  %+v", o._statusCode, o.Payload)
}
func (o *PutCustomersCustomerFidUpgradePaymentAccountDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutCustomersCustomerFidUpgradePaymentAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
