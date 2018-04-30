// Code generated by go-swagger; DO NOT EDIT.

package contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

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

	case 404:
		result := NewPutContactsEmailsEmailAddressConfirmNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

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

/*PutContactsEmailsEmailAddressConfirmOK handles this case with default header values.

Email Address Confirmed
*/
type PutContactsEmailsEmailAddressConfirmOK struct {
}

func (o *PutContactsEmailsEmailAddressConfirmOK) Error() string {
	return fmt.Sprintf("[PUT /contacts/emails/{emailAddress}/confirm][%d] putContactsEmailsEmailAddressConfirmOK ", 200)
}

func (o *PutContactsEmailsEmailAddressConfirmOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutContactsEmailsEmailAddressConfirmNotFound creates a PutContactsEmailsEmailAddressConfirmNotFound with default headers values
func NewPutContactsEmailsEmailAddressConfirmNotFound() *PutContactsEmailsEmailAddressConfirmNotFound {
	return &PutContactsEmailsEmailAddressConfirmNotFound{}
}

/*PutContactsEmailsEmailAddressConfirmNotFound handles this case with default header values.

Email Address not found
*/
type PutContactsEmailsEmailAddressConfirmNotFound struct {
}

func (o *PutContactsEmailsEmailAddressConfirmNotFound) Error() string {
	return fmt.Sprintf("[PUT /contacts/emails/{emailAddress}/confirm][%d] putContactsEmailsEmailAddressConfirmNotFound ", 404)
}

func (o *PutContactsEmailsEmailAddressConfirmNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutContactsEmailsEmailAddressConfirmDefault creates a PutContactsEmailsEmailAddressConfirmDefault with default headers values
func NewPutContactsEmailsEmailAddressConfirmDefault(code int) *PutContactsEmailsEmailAddressConfirmDefault {
	return &PutContactsEmailsEmailAddressConfirmDefault{
		_statusCode: code,
	}
}

/*PutContactsEmailsEmailAddressConfirmDefault handles this case with default header values.

Error
*/
type PutContactsEmailsEmailAddressConfirmDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the put contacts emails email address confirm default response
func (o *PutContactsEmailsEmailAddressConfirmDefault) Code() int {
	return o._statusCode
}

func (o *PutContactsEmailsEmailAddressConfirmDefault) Error() string {
	return fmt.Sprintf("[PUT /contacts/emails/{emailAddress}/confirm][%d] PutContactsEmailsEmailAddressConfirm default  %+v", o._statusCode, o.Payload)
}

func (o *PutContactsEmailsEmailAddressConfirmDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
