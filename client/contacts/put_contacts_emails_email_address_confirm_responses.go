// Code generated by go-swagger; DO NOT EDIT.

package contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PutContactsEmailsEmailAddressConfirmReader is a Reader for the PutContactsEmailsEmailAddressConfirm structure.
type PutContactsEmailsEmailAddressConfirmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutContactsEmailsEmailAddressConfirmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutContactsEmailsEmailAddressConfirmOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutContactsEmailsEmailAddressConfirmDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutContactsEmailsEmailAddressConfirmOK creates a PutContactsEmailsEmailAddressConfirmOK with default headers values
func NewPutContactsEmailsEmailAddressConfirmOK() *PutContactsEmailsEmailAddressConfirmOK {
	return &PutContactsEmailsEmailAddressConfirmOK{}
}

/*
PutContactsEmailsEmailAddressConfirmOK describes a response with status code 200, with default header values.

Email Address Confirmed
*/
type PutContactsEmailsEmailAddressConfirmOK struct {
}

// IsSuccess returns true when this put contacts emails email address confirm o k response has a 2xx status code
func (o *PutContactsEmailsEmailAddressConfirmOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put contacts emails email address confirm o k response has a 3xx status code
func (o *PutContactsEmailsEmailAddressConfirmOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put contacts emails email address confirm o k response has a 4xx status code
func (o *PutContactsEmailsEmailAddressConfirmOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put contacts emails email address confirm o k response has a 5xx status code
func (o *PutContactsEmailsEmailAddressConfirmOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put contacts emails email address confirm o k response a status code equal to that given
func (o *PutContactsEmailsEmailAddressConfirmOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put contacts emails email address confirm o k response
func (o *PutContactsEmailsEmailAddressConfirmOK) Code() int {
	return 200
}

func (o *PutContactsEmailsEmailAddressConfirmOK) Error() string {
	return fmt.Sprintf("[PUT /contacts/emails/{emailAddress}/confirm][%d] putContactsEmailsEmailAddressConfirmOK ", 200)
}

func (o *PutContactsEmailsEmailAddressConfirmOK) String() string {
	return fmt.Sprintf("[PUT /contacts/emails/{emailAddress}/confirm][%d] putContactsEmailsEmailAddressConfirmOK ", 200)
}

func (o *PutContactsEmailsEmailAddressConfirmOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutContactsEmailsEmailAddressConfirmDefault creates a PutContactsEmailsEmailAddressConfirmDefault with default headers values
func NewPutContactsEmailsEmailAddressConfirmDefault(code int) *PutContactsEmailsEmailAddressConfirmDefault {
	return &PutContactsEmailsEmailAddressConfirmDefault{
		_statusCode: code,
	}
}

/*
PutContactsEmailsEmailAddressConfirmDefault describes a response with status code -1, with default header values.

Error
*/
type PutContactsEmailsEmailAddressConfirmDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this put contacts emails email address confirm default response has a 2xx status code
func (o *PutContactsEmailsEmailAddressConfirmDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put contacts emails email address confirm default response has a 3xx status code
func (o *PutContactsEmailsEmailAddressConfirmDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put contacts emails email address confirm default response has a 4xx status code
func (o *PutContactsEmailsEmailAddressConfirmDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put contacts emails email address confirm default response has a 5xx status code
func (o *PutContactsEmailsEmailAddressConfirmDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put contacts emails email address confirm default response a status code equal to that given
func (o *PutContactsEmailsEmailAddressConfirmDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put contacts emails email address confirm default response
func (o *PutContactsEmailsEmailAddressConfirmDefault) Code() int {
	return o._statusCode
}

func (o *PutContactsEmailsEmailAddressConfirmDefault) Error() string {
	return fmt.Sprintf("[PUT /contacts/emails/{emailAddress}/confirm][%d] PutContactsEmailsEmailAddressConfirm default  %+v", o._statusCode, o.Payload)
}

func (o *PutContactsEmailsEmailAddressConfirmDefault) String() string {
	return fmt.Sprintf("[PUT /contacts/emails/{emailAddress}/confirm][%d] PutContactsEmailsEmailAddressConfirm default  %+v", o._statusCode, o.Payload)
}

func (o *PutContactsEmailsEmailAddressConfirmDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PutContactsEmailsEmailAddressConfirmDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
