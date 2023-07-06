// Code generated by go-swagger; DO NOT EDIT.

package interactions

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

// PostInteractionsInviteInviteCodeReader is a Reader for the PostInteractionsInviteInviteCode structure.
type PostInteractionsInviteInviteCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostInteractionsInviteInviteCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostInteractionsInviteInviteCodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostInteractionsInviteInviteCodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostInteractionsInviteInviteCodeOK creates a PostInteractionsInviteInviteCodeOK with default headers values
func NewPostInteractionsInviteInviteCodeOK() *PostInteractionsInviteInviteCodeOK {
	return &PostInteractionsInviteInviteCodeOK{}
}

/*
PostInteractionsInviteInviteCodeOK describes a response with status code 200, with default header values.

Accepted an interaction invite
*/
type PostInteractionsInviteInviteCodeOK struct {
	Payload *PostInteractionsInviteInviteCodeOKBody
}

// IsSuccess returns true when this post interactions invite invite code o k response has a 2xx status code
func (o *PostInteractionsInviteInviteCodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post interactions invite invite code o k response has a 3xx status code
func (o *PostInteractionsInviteInviteCodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post interactions invite invite code o k response has a 4xx status code
func (o *PostInteractionsInviteInviteCodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post interactions invite invite code o k response has a 5xx status code
func (o *PostInteractionsInviteInviteCodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post interactions invite invite code o k response a status code equal to that given
func (o *PostInteractionsInviteInviteCodeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post interactions invite invite code o k response
func (o *PostInteractionsInviteInviteCodeOK) Code() int {
	return 200
}

func (o *PostInteractionsInviteInviteCodeOK) Error() string {
	return fmt.Sprintf("[POST /interactions/invite/{inviteCode}][%d] postInteractionsInviteInviteCodeOK  %+v", 200, o.Payload)
}

func (o *PostInteractionsInviteInviteCodeOK) String() string {
	return fmt.Sprintf("[POST /interactions/invite/{inviteCode}][%d] postInteractionsInviteInviteCodeOK  %+v", 200, o.Payload)
}

func (o *PostInteractionsInviteInviteCodeOK) GetPayload() *PostInteractionsInviteInviteCodeOKBody {
	return o.Payload
}

func (o *PostInteractionsInviteInviteCodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostInteractionsInviteInviteCodeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInteractionsInviteInviteCodeDefault creates a PostInteractionsInviteInviteCodeDefault with default headers values
func NewPostInteractionsInviteInviteCodeDefault(code int) *PostInteractionsInviteInviteCodeDefault {
	return &PostInteractionsInviteInviteCodeDefault{
		_statusCode: code,
	}
}

/*
PostInteractionsInviteInviteCodeDefault describes a response with status code -1, with default header values.

Error
*/
type PostInteractionsInviteInviteCodeDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this post interactions invite invite code default response has a 2xx status code
func (o *PostInteractionsInviteInviteCodeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post interactions invite invite code default response has a 3xx status code
func (o *PostInteractionsInviteInviteCodeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post interactions invite invite code default response has a 4xx status code
func (o *PostInteractionsInviteInviteCodeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post interactions invite invite code default response has a 5xx status code
func (o *PostInteractionsInviteInviteCodeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post interactions invite invite code default response a status code equal to that given
func (o *PostInteractionsInviteInviteCodeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post interactions invite invite code default response
func (o *PostInteractionsInviteInviteCodeDefault) Code() int {
	return o._statusCode
}

func (o *PostInteractionsInviteInviteCodeDefault) Error() string {
	return fmt.Sprintf("[POST /interactions/invite/{inviteCode}][%d] PostInteractionsInviteInviteCode default  %+v", o._statusCode, o.Payload)
}

func (o *PostInteractionsInviteInviteCodeDefault) String() string {
	return fmt.Sprintf("[POST /interactions/invite/{inviteCode}][%d] PostInteractionsInviteInviteCode default  %+v", o._statusCode, o.Payload)
}

func (o *PostInteractionsInviteInviteCodeDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostInteractionsInviteInviteCodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostInteractionsInviteInviteCodeOKBody post interactions invite invite code o k body
swagger:model PostInteractionsInviteInviteCodeOKBody
*/
type PostInteractionsInviteInviteCodeOKBody struct {
	models.Envelope

	// data
	Data *models.InteractionResponse `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostInteractionsInviteInviteCodeOKBody) UnmarshalJSON(raw []byte) error {
	// PostInteractionsInviteInviteCodeOKBodyAO0
	var postInteractionsInviteInviteCodeOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postInteractionsInviteInviteCodeOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postInteractionsInviteInviteCodeOKBodyAO0

	// PostInteractionsInviteInviteCodeOKBodyAO1
	var dataPostInteractionsInviteInviteCodeOKBodyAO1 struct {
		Data *models.InteractionResponse `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostInteractionsInviteInviteCodeOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostInteractionsInviteInviteCodeOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostInteractionsInviteInviteCodeOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postInteractionsInviteInviteCodeOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postInteractionsInviteInviteCodeOKBodyAO0)
	var dataPostInteractionsInviteInviteCodeOKBodyAO1 struct {
		Data *models.InteractionResponse `json:"data,omitempty"`
	}

	dataPostInteractionsInviteInviteCodeOKBodyAO1.Data = o.Data

	jsonDataPostInteractionsInviteInviteCodeOKBodyAO1, errPostInteractionsInviteInviteCodeOKBodyAO1 := swag.WriteJSON(dataPostInteractionsInviteInviteCodeOKBodyAO1)
	if errPostInteractionsInviteInviteCodeOKBodyAO1 != nil {
		return nil, errPostInteractionsInviteInviteCodeOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostInteractionsInviteInviteCodeOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post interactions invite invite code o k body
func (o *PostInteractionsInviteInviteCodeOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostInteractionsInviteInviteCodeOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postInteractionsInviteInviteCodeOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postInteractionsInviteInviteCodeOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post interactions invite invite code o k body based on the context it is used
func (o *PostInteractionsInviteInviteCodeOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostInteractionsInviteInviteCodeOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postInteractionsInviteInviteCodeOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postInteractionsInviteInviteCodeOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostInteractionsInviteInviteCodeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostInteractionsInviteInviteCodeOKBody) UnmarshalBinary(b []byte) error {
	var res PostInteractionsInviteInviteCodeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}