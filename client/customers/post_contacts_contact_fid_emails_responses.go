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

// PostContactsContactFidEmailsReader is a Reader for the PostContactsContactFidEmails structure.
type PostContactsContactFidEmailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostContactsContactFidEmailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostContactsContactFidEmailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostContactsContactFidEmailsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostContactsContactFidEmailsOK creates a PostContactsContactFidEmailsOK with default headers values
func NewPostContactsContactFidEmailsOK() *PostContactsContactFidEmailsOK {
	return &PostContactsContactFidEmailsOK{}
}

/*
PostContactsContactFidEmailsOK describes a response with status code 200, with default header values.

Email Added
*/
type PostContactsContactFidEmailsOK struct {
}

// IsSuccess returns true when this post contacts contact fid emails o k response has a 2xx status code
func (o *PostContactsContactFidEmailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post contacts contact fid emails o k response has a 3xx status code
func (o *PostContactsContactFidEmailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post contacts contact fid emails o k response has a 4xx status code
func (o *PostContactsContactFidEmailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post contacts contact fid emails o k response has a 5xx status code
func (o *PostContactsContactFidEmailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post contacts contact fid emails o k response a status code equal to that given
func (o *PostContactsContactFidEmailsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post contacts contact fid emails o k response
func (o *PostContactsContactFidEmailsOK) Code() int {
	return 200
}

func (o *PostContactsContactFidEmailsOK) Error() string {
	return fmt.Sprintf("[POST /contacts/{contactFid}/emails][%d] postContactsContactFidEmailsOK ", 200)
}

func (o *PostContactsContactFidEmailsOK) String() string {
	return fmt.Sprintf("[POST /contacts/{contactFid}/emails][%d] postContactsContactFidEmailsOK ", 200)
}

func (o *PostContactsContactFidEmailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostContactsContactFidEmailsDefault creates a PostContactsContactFidEmailsDefault with default headers values
func NewPostContactsContactFidEmailsDefault(code int) *PostContactsContactFidEmailsDefault {
	return &PostContactsContactFidEmailsDefault{
		_statusCode: code,
	}
}

/*
PostContactsContactFidEmailsDefault describes a response with status code -1, with default header values.

Error
*/
type PostContactsContactFidEmailsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this post contacts contact fid emails default response has a 2xx status code
func (o *PostContactsContactFidEmailsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post contacts contact fid emails default response has a 3xx status code
func (o *PostContactsContactFidEmailsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post contacts contact fid emails default response has a 4xx status code
func (o *PostContactsContactFidEmailsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post contacts contact fid emails default response has a 5xx status code
func (o *PostContactsContactFidEmailsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post contacts contact fid emails default response a status code equal to that given
func (o *PostContactsContactFidEmailsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post contacts contact fid emails default response
func (o *PostContactsContactFidEmailsDefault) Code() int {
	return o._statusCode
}

func (o *PostContactsContactFidEmailsDefault) Error() string {
	return fmt.Sprintf("[POST /contacts/{contactFid}/emails][%d] PostContactsContactFidEmails default  %+v", o._statusCode, o.Payload)
}

func (o *PostContactsContactFidEmailsDefault) String() string {
	return fmt.Sprintf("[POST /contacts/{contactFid}/emails][%d] PostContactsContactFidEmails default  %+v", o._statusCode, o.Payload)
}

func (o *PostContactsContactFidEmailsDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostContactsContactFidEmailsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
