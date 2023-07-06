// Code generated by go-swagger; DO NOT EDIT.

package chat

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

// PostPresetChatReader is a Reader for the PostPresetChat structure.
type PostPresetChatReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPresetChatReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostPresetChatOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostPresetChatDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostPresetChatOK creates a PostPresetChatOK with default headers values
func NewPostPresetChatOK() *PostPresetChatOK {
	return &PostPresetChatOK{}
}

/*
PostPresetChatOK describes a response with status code 200, with default header values.

Device created/updated
*/
type PostPresetChatOK struct {
	Payload *PostPresetChatOKBody
}

// IsSuccess returns true when this post preset chat o k response has a 2xx status code
func (o *PostPresetChatOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post preset chat o k response has a 3xx status code
func (o *PostPresetChatOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post preset chat o k response has a 4xx status code
func (o *PostPresetChatOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post preset chat o k response has a 5xx status code
func (o *PostPresetChatOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post preset chat o k response a status code equal to that given
func (o *PostPresetChatOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post preset chat o k response
func (o *PostPresetChatOK) Code() int {
	return 200
}

func (o *PostPresetChatOK) Error() string {
	return fmt.Sprintf("[POST /presetChat][%d] postPresetChatOK  %+v", 200, o.Payload)
}

func (o *PostPresetChatOK) String() string {
	return fmt.Sprintf("[POST /presetChat][%d] postPresetChatOK  %+v", 200, o.Payload)
}

func (o *PostPresetChatOK) GetPayload() *PostPresetChatOKBody {
	return o.Payload
}

func (o *PostPresetChatOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostPresetChatOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresetChatDefault creates a PostPresetChatDefault with default headers values
func NewPostPresetChatDefault(code int) *PostPresetChatDefault {
	return &PostPresetChatDefault{
		_statusCode: code,
	}
}

/*
PostPresetChatDefault describes a response with status code -1, with default header values.

Error
*/
type PostPresetChatDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this post preset chat default response has a 2xx status code
func (o *PostPresetChatDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post preset chat default response has a 3xx status code
func (o *PostPresetChatDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post preset chat default response has a 4xx status code
func (o *PostPresetChatDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post preset chat default response has a 5xx status code
func (o *PostPresetChatDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post preset chat default response a status code equal to that given
func (o *PostPresetChatDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post preset chat default response
func (o *PostPresetChatDefault) Code() int {
	return o._statusCode
}

func (o *PostPresetChatDefault) Error() string {
	return fmt.Sprintf("[POST /presetChat][%d] PostPresetChat default  %+v", o._statusCode, o.Payload)
}

func (o *PostPresetChatDefault) String() string {
	return fmt.Sprintf("[POST /presetChat][%d] PostPresetChat default  %+v", o._statusCode, o.Payload)
}

func (o *PostPresetChatDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostPresetChatDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostPresetChatOKBody post preset chat o k body
swagger:model PostPresetChatOKBody
*/
type PostPresetChatOKBody struct {
	models.Envelope

	// data
	Data *models.PresetChatSessionLink `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostPresetChatOKBody) UnmarshalJSON(raw []byte) error {
	// PostPresetChatOKBodyAO0
	var postPresetChatOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postPresetChatOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postPresetChatOKBodyAO0

	// PostPresetChatOKBodyAO1
	var dataPostPresetChatOKBodyAO1 struct {
		Data *models.PresetChatSessionLink `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostPresetChatOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostPresetChatOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostPresetChatOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postPresetChatOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postPresetChatOKBodyAO0)
	var dataPostPresetChatOKBodyAO1 struct {
		Data *models.PresetChatSessionLink `json:"data,omitempty"`
	}

	dataPostPresetChatOKBodyAO1.Data = o.Data

	jsonDataPostPresetChatOKBodyAO1, errPostPresetChatOKBodyAO1 := swag.WriteJSON(dataPostPresetChatOKBodyAO1)
	if errPostPresetChatOKBodyAO1 != nil {
		return nil, errPostPresetChatOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostPresetChatOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post preset chat o k body
func (o *PostPresetChatOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostPresetChatOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postPresetChatOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postPresetChatOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post preset chat o k body based on the context it is used
func (o *PostPresetChatOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostPresetChatOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postPresetChatOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postPresetChatOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostPresetChatOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostPresetChatOKBody) UnmarshalBinary(b []byte) error {
	var res PostPresetChatOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}