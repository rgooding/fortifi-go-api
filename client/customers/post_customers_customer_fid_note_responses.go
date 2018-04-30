// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PostCustomersCustomerFidNoteReader is a Reader for the PostCustomersCustomerFidNote structure.
type PostCustomersCustomerFidNoteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomersCustomerFidNoteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostCustomersCustomerFidNoteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostCustomersCustomerFidNoteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostCustomersCustomerFidNoteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCustomersCustomerFidNoteOK creates a PostCustomersCustomerFidNoteOK with default headers values
func NewPostCustomersCustomerFidNoteOK() *PostCustomersCustomerFidNoteOK {
	return &PostCustomersCustomerFidNoteOK{}
}

/*PostCustomersCustomerFidNoteOK handles this case with default header values.

Customer note saved
*/
type PostCustomersCustomerFidNoteOK struct {
}

func (o *PostCustomersCustomerFidNoteOK) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/note][%d] postCustomersCustomerFidNoteOK ", 200)
}

func (o *PostCustomersCustomerFidNoteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomersCustomerFidNoteBadRequest creates a PostCustomersCustomerFidNoteBadRequest with default headers values
func NewPostCustomersCustomerFidNoteBadRequest() *PostCustomersCustomerFidNoteBadRequest {
	return &PostCustomersCustomerFidNoteBadRequest{}
}

/*PostCustomersCustomerFidNoteBadRequest handles this case with default header values.

Invalid Payload
*/
type PostCustomersCustomerFidNoteBadRequest struct {
}

func (o *PostCustomersCustomerFidNoteBadRequest) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/note][%d] postCustomersCustomerFidNoteBadRequest ", 400)
}

func (o *PostCustomersCustomerFidNoteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomersCustomerFidNoteDefault creates a PostCustomersCustomerFidNoteDefault with default headers values
func NewPostCustomersCustomerFidNoteDefault(code int) *PostCustomersCustomerFidNoteDefault {
	return &PostCustomersCustomerFidNoteDefault{
		_statusCode: code,
	}
}

/*PostCustomersCustomerFidNoteDefault handles this case with default header values.

Error
*/
type PostCustomersCustomerFidNoteDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post customers customer fid note default response
func (o *PostCustomersCustomerFidNoteDefault) Code() int {
	return o._statusCode
}

func (o *PostCustomersCustomerFidNoteDefault) Error() string {
	return fmt.Sprintf("[POST /customers/{customerFid}/note][%d] PostCustomersCustomerFidNote default  %+v", o._statusCode, o.Payload)
}

func (o *PostCustomersCustomerFidNoteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
