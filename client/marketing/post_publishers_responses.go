// Code generated by go-swagger; DO NOT EDIT.

package marketing

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

// PostPublishersReader is a Reader for the PostPublishers structure.
type PostPublishersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPublishersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostPublishersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostPublishersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostPublishersOK creates a PostPublishersOK with default headers values
func NewPostPublishersOK() *PostPublishersOK {
	return &PostPublishersOK{}
}

/*
PostPublishersOK describes a response with status code 200, with default header values.

Publisher Created
*/
type PostPublishersOK struct {
	Payload *PostPublishersOKBody
}

// IsSuccess returns true when this post publishers o k response has a 2xx status code
func (o *PostPublishersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post publishers o k response has a 3xx status code
func (o *PostPublishersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post publishers o k response has a 4xx status code
func (o *PostPublishersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post publishers o k response has a 5xx status code
func (o *PostPublishersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post publishers o k response a status code equal to that given
func (o *PostPublishersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post publishers o k response
func (o *PostPublishersOK) Code() int {
	return 200
}

func (o *PostPublishersOK) Error() string {
	return fmt.Sprintf("[POST /publishers][%d] postPublishersOK  %+v", 200, o.Payload)
}

func (o *PostPublishersOK) String() string {
	return fmt.Sprintf("[POST /publishers][%d] postPublishersOK  %+v", 200, o.Payload)
}

func (o *PostPublishersOK) GetPayload() *PostPublishersOKBody {
	return o.Payload
}

func (o *PostPublishersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostPublishersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPublishersDefault creates a PostPublishersDefault with default headers values
func NewPostPublishersDefault(code int) *PostPublishersDefault {
	return &PostPublishersDefault{
		_statusCode: code,
	}
}

/*
PostPublishersDefault describes a response with status code -1, with default header values.

Error
*/
type PostPublishersDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this post publishers default response has a 2xx status code
func (o *PostPublishersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post publishers default response has a 3xx status code
func (o *PostPublishersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post publishers default response has a 4xx status code
func (o *PostPublishersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post publishers default response has a 5xx status code
func (o *PostPublishersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post publishers default response a status code equal to that given
func (o *PostPublishersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post publishers default response
func (o *PostPublishersDefault) Code() int {
	return o._statusCode
}

func (o *PostPublishersDefault) Error() string {
	return fmt.Sprintf("[POST /publishers][%d] PostPublishers default  %+v", o._statusCode, o.Payload)
}

func (o *PostPublishersDefault) String() string {
	return fmt.Sprintf("[POST /publishers][%d] PostPublishers default  %+v", o._statusCode, o.Payload)
}

func (o *PostPublishersDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostPublishersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostPublishersOKBody post publishers o k body
swagger:model PostPublishersOKBody
*/
type PostPublishersOKBody struct {
	models.Envelope

	// data
	Data *models.Publisher `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostPublishersOKBody) UnmarshalJSON(raw []byte) error {
	// PostPublishersOKBodyAO0
	var postPublishersOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postPublishersOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postPublishersOKBodyAO0

	// PostPublishersOKBodyAO1
	var dataPostPublishersOKBodyAO1 struct {
		Data *models.Publisher `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostPublishersOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostPublishersOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostPublishersOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postPublishersOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postPublishersOKBodyAO0)
	var dataPostPublishersOKBodyAO1 struct {
		Data *models.Publisher `json:"data,omitempty"`
	}

	dataPostPublishersOKBodyAO1.Data = o.Data

	jsonDataPostPublishersOKBodyAO1, errPostPublishersOKBodyAO1 := swag.WriteJSON(dataPostPublishersOKBodyAO1)
	if errPostPublishersOKBodyAO1 != nil {
		return nil, errPostPublishersOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostPublishersOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post publishers o k body
func (o *PostPublishersOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostPublishersOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postPublishersOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postPublishersOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post publishers o k body based on the context it is used
func (o *PostPublishersOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostPublishersOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postPublishersOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postPublishersOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostPublishersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostPublishersOKBody) UnmarshalBinary(b []byte) error {
	var res PostPublishersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
