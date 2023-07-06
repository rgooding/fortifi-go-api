// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/fortifi/go-api/models"
)

// PostTicketsTicketFidNoteReader is a Reader for the PostTicketsTicketFidNote structure.
type PostTicketsTicketFidNoteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTicketsTicketFidNoteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTicketsTicketFidNoteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostTicketsTicketFidNoteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostTicketsTicketFidNoteOK creates a PostTicketsTicketFidNoteOK with default headers values
func NewPostTicketsTicketFidNoteOK() *PostTicketsTicketFidNoteOK {
	return &PostTicketsTicketFidNoteOK{}
}

/*
PostTicketsTicketFidNoteOK describes a response with status code 200, with default header values.

Success Response
*/
type PostTicketsTicketFidNoteOK struct {
	Payload *PostTicketsTicketFidNoteOKBody
}

// IsSuccess returns true when this post tickets ticket fid note o k response has a 2xx status code
func (o *PostTicketsTicketFidNoteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post tickets ticket fid note o k response has a 3xx status code
func (o *PostTicketsTicketFidNoteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post tickets ticket fid note o k response has a 4xx status code
func (o *PostTicketsTicketFidNoteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post tickets ticket fid note o k response has a 5xx status code
func (o *PostTicketsTicketFidNoteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post tickets ticket fid note o k response a status code equal to that given
func (o *PostTicketsTicketFidNoteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post tickets ticket fid note o k response
func (o *PostTicketsTicketFidNoteOK) Code() int {
	return 200
}

func (o *PostTicketsTicketFidNoteOK) Error() string {
	return fmt.Sprintf("[POST /tickets/{ticketFid}/note][%d] postTicketsTicketFidNoteOK  %+v", 200, o.Payload)
}

func (o *PostTicketsTicketFidNoteOK) String() string {
	return fmt.Sprintf("[POST /tickets/{ticketFid}/note][%d] postTicketsTicketFidNoteOK  %+v", 200, o.Payload)
}

func (o *PostTicketsTicketFidNoteOK) GetPayload() *PostTicketsTicketFidNoteOKBody {
	return o.Payload
}

func (o *PostTicketsTicketFidNoteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostTicketsTicketFidNoteOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTicketsTicketFidNoteDefault creates a PostTicketsTicketFidNoteDefault with default headers values
func NewPostTicketsTicketFidNoteDefault(code int) *PostTicketsTicketFidNoteDefault {
	return &PostTicketsTicketFidNoteDefault{
		_statusCode: code,
	}
}

/*
PostTicketsTicketFidNoteDefault describes a response with status code -1, with default header values.

Error
*/
type PostTicketsTicketFidNoteDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this post tickets ticket fid note default response has a 2xx status code
func (o *PostTicketsTicketFidNoteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post tickets ticket fid note default response has a 3xx status code
func (o *PostTicketsTicketFidNoteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post tickets ticket fid note default response has a 4xx status code
func (o *PostTicketsTicketFidNoteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post tickets ticket fid note default response has a 5xx status code
func (o *PostTicketsTicketFidNoteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post tickets ticket fid note default response a status code equal to that given
func (o *PostTicketsTicketFidNoteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post tickets ticket fid note default response
func (o *PostTicketsTicketFidNoteDefault) Code() int {
	return o._statusCode
}

func (o *PostTicketsTicketFidNoteDefault) Error() string {
	return fmt.Sprintf("[POST /tickets/{ticketFid}/note][%d] PostTicketsTicketFidNote default  %+v", o._statusCode, o.Payload)
}

func (o *PostTicketsTicketFidNoteDefault) String() string {
	return fmt.Sprintf("[POST /tickets/{ticketFid}/note][%d] PostTicketsTicketFidNote default  %+v", o._statusCode, o.Payload)
}

func (o *PostTicketsTicketFidNoteDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostTicketsTicketFidNoteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostTicketsTicketFidNoteOKBody post tickets ticket fid note o k body
swagger:model PostTicketsTicketFidNoteOKBody
*/
type PostTicketsTicketFidNoteOKBody struct {
	models.Envelope

	// data
	Data *models.BoolMessage `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostTicketsTicketFidNoteOKBody) UnmarshalJSON(raw []byte) error {
	// PostTicketsTicketFidNoteOKBodyAO0
	var postTicketsTicketFidNoteOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postTicketsTicketFidNoteOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postTicketsTicketFidNoteOKBodyAO0

	// PostTicketsTicketFidNoteOKBodyAO1
	var dataPostTicketsTicketFidNoteOKBodyAO1 struct {
		Data *models.BoolMessage `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostTicketsTicketFidNoteOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostTicketsTicketFidNoteOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostTicketsTicketFidNoteOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postTicketsTicketFidNoteOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postTicketsTicketFidNoteOKBodyAO0)
	var dataPostTicketsTicketFidNoteOKBodyAO1 struct {
		Data *models.BoolMessage `json:"data,omitempty"`
	}

	dataPostTicketsTicketFidNoteOKBodyAO1.Data = o.Data

	jsonDataPostTicketsTicketFidNoteOKBodyAO1, errPostTicketsTicketFidNoteOKBodyAO1 := swag.WriteJSON(dataPostTicketsTicketFidNoteOKBodyAO1)
	if errPostTicketsTicketFidNoteOKBodyAO1 != nil {
		return nil, errPostTicketsTicketFidNoteOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostTicketsTicketFidNoteOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post tickets ticket fid note o k body
func (o *PostTicketsTicketFidNoteOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Envelope
	if err := o.Envelope.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTicketsTicketFidNoteOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postTicketsTicketFidNoteOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postTicketsTicketFidNoteOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post tickets ticket fid note o k body based on the context it is used
func (o *PostTicketsTicketFidNoteOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Envelope
	if err := o.Envelope.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTicketsTicketFidNoteOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postTicketsTicketFidNoteOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postTicketsTicketFidNoteOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostTicketsTicketFidNoteOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTicketsTicketFidNoteOKBody) UnmarshalBinary(b []byte) error {
	var res PostTicketsTicketFidNoteOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}