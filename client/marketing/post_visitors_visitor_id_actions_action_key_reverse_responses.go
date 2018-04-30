// Code generated by go-swagger; DO NOT EDIT.

package marketing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// PostVisitorsVisitorIDActionsActionKeyReverseReader is a Reader for the PostVisitorsVisitorIDActionsActionKeyReverse structure.
type PostVisitorsVisitorIDActionsActionKeyReverseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVisitorsVisitorIDActionsActionKeyReverseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostVisitorsVisitorIDActionsActionKeyReverseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPostVisitorsVisitorIDActionsActionKeyReverseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostVisitorsVisitorIDActionsActionKeyReverseConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostVisitorsVisitorIDActionsActionKeyReverseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostVisitorsVisitorIDActionsActionKeyReverseOK creates a PostVisitorsVisitorIDActionsActionKeyReverseOK with default headers values
func NewPostVisitorsVisitorIDActionsActionKeyReverseOK() *PostVisitorsVisitorIDActionsActionKeyReverseOK {
	return &PostVisitorsVisitorIDActionsActionKeyReverseOK{}
}

/*PostVisitorsVisitorIDActionsActionKeyReverseOK handles this case with default header values.

Action Reversed
*/
type PostVisitorsVisitorIDActionsActionKeyReverseOK struct {
	Payload *models.PostVisitorsVisitorIDActionsActionKeyReverseOKBody
}

func (o *PostVisitorsVisitorIDActionsActionKeyReverseOK) Error() string {
	return fmt.Sprintf("[POST /visitors/{visitorId}/actions/{actionKey}/reverse][%d] postVisitorsVisitorIdActionsActionKeyReverseOK  %+v", 200, o.Payload)
}

func (o *PostVisitorsVisitorIDActionsActionKeyReverseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostVisitorsVisitorIDActionsActionKeyReverseOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVisitorsVisitorIDActionsActionKeyReverseNotFound creates a PostVisitorsVisitorIDActionsActionKeyReverseNotFound with default headers values
func NewPostVisitorsVisitorIDActionsActionKeyReverseNotFound() *PostVisitorsVisitorIDActionsActionKeyReverseNotFound {
	return &PostVisitorsVisitorIDActionsActionKeyReverseNotFound{}
}

/*PostVisitorsVisitorIDActionsActionKeyReverseNotFound handles this case with default header values.

The action you are trying to reverse cannot be found

*/
type PostVisitorsVisitorIDActionsActionKeyReverseNotFound struct {
}

func (o *PostVisitorsVisitorIDActionsActionKeyReverseNotFound) Error() string {
	return fmt.Sprintf("[POST /visitors/{visitorId}/actions/{actionKey}/reverse][%d] postVisitorsVisitorIdActionsActionKeyReverseNotFound ", 404)
}

func (o *PostVisitorsVisitorIDActionsActionKeyReverseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostVisitorsVisitorIDActionsActionKeyReverseConflict creates a PostVisitorsVisitorIDActionsActionKeyReverseConflict with default headers values
func NewPostVisitorsVisitorIDActionsActionKeyReverseConflict() *PostVisitorsVisitorIDActionsActionKeyReverseConflict {
	return &PostVisitorsVisitorIDActionsActionKeyReverseConflict{}
}

/*PostVisitorsVisitorIDActionsActionKeyReverseConflict handles this case with default header values.

The action specified has already been reversed

*/
type PostVisitorsVisitorIDActionsActionKeyReverseConflict struct {
}

func (o *PostVisitorsVisitorIDActionsActionKeyReverseConflict) Error() string {
	return fmt.Sprintf("[POST /visitors/{visitorId}/actions/{actionKey}/reverse][%d] postVisitorsVisitorIdActionsActionKeyReverseConflict ", 409)
}

func (o *PostVisitorsVisitorIDActionsActionKeyReverseConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostVisitorsVisitorIDActionsActionKeyReverseDefault creates a PostVisitorsVisitorIDActionsActionKeyReverseDefault with default headers values
func NewPostVisitorsVisitorIDActionsActionKeyReverseDefault(code int) *PostVisitorsVisitorIDActionsActionKeyReverseDefault {
	return &PostVisitorsVisitorIDActionsActionKeyReverseDefault{
		_statusCode: code,
	}
}

/*PostVisitorsVisitorIDActionsActionKeyReverseDefault handles this case with default header values.

Error
*/
type PostVisitorsVisitorIDActionsActionKeyReverseDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post visitors visitor ID actions action key reverse default response
func (o *PostVisitorsVisitorIDActionsActionKeyReverseDefault) Code() int {
	return o._statusCode
}

func (o *PostVisitorsVisitorIDActionsActionKeyReverseDefault) Error() string {
	return fmt.Sprintf("[POST /visitors/{visitorId}/actions/{actionKey}/reverse][%d] PostVisitorsVisitorIDActionsActionKeyReverse default  %+v", o._statusCode, o.Payload)
}

func (o *PostVisitorsVisitorIDActionsActionKeyReverseDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
