// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// PostDeviceUpsertReader is a Reader for the PostDeviceUpsert structure.
type PostDeviceUpsertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDeviceUpsertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDeviceUpsertOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostDeviceUpsertDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostDeviceUpsertOK creates a PostDeviceUpsertOK with default headers values
func NewPostDeviceUpsertOK() *PostDeviceUpsertOK {
	return &PostDeviceUpsertOK{}
}

/*
PostDeviceUpsertOK describes a response with status code 200, with default header values.

Device created/updated
*/
type PostDeviceUpsertOK struct {
	Payload *PostDeviceUpsertOKBody
}

// IsSuccess returns true when this post device upsert o k response has a 2xx status code
func (o *PostDeviceUpsertOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post device upsert o k response has a 3xx status code
func (o *PostDeviceUpsertOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post device upsert o k response has a 4xx status code
func (o *PostDeviceUpsertOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post device upsert o k response has a 5xx status code
func (o *PostDeviceUpsertOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post device upsert o k response a status code equal to that given
func (o *PostDeviceUpsertOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostDeviceUpsertOK) Error() string {
	return fmt.Sprintf("[POST /device/upsert][%d] postDeviceUpsertOK  %+v", 200, o.Payload)
}

func (o *PostDeviceUpsertOK) String() string {
	return fmt.Sprintf("[POST /device/upsert][%d] postDeviceUpsertOK  %+v", 200, o.Payload)
}

func (o *PostDeviceUpsertOK) GetPayload() *PostDeviceUpsertOKBody {
	return o.Payload
}

func (o *PostDeviceUpsertOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostDeviceUpsertOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDeviceUpsertDefault creates a PostDeviceUpsertDefault with default headers values
func NewPostDeviceUpsertDefault(code int) *PostDeviceUpsertDefault {
	return &PostDeviceUpsertDefault{
		_statusCode: code,
	}
}

/*
PostDeviceUpsertDefault describes a response with status code -1, with default header values.

Error
*/
type PostDeviceUpsertDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post device upsert default response
func (o *PostDeviceUpsertDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this post device upsert default response has a 2xx status code
func (o *PostDeviceUpsertDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post device upsert default response has a 3xx status code
func (o *PostDeviceUpsertDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post device upsert default response has a 4xx status code
func (o *PostDeviceUpsertDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post device upsert default response has a 5xx status code
func (o *PostDeviceUpsertDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post device upsert default response a status code equal to that given
func (o *PostDeviceUpsertDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PostDeviceUpsertDefault) Error() string {
	return fmt.Sprintf("[POST /device/upsert][%d] PostDeviceUpsert default  %+v", o._statusCode, o.Payload)
}

func (o *PostDeviceUpsertDefault) String() string {
	return fmt.Sprintf("[POST /device/upsert][%d] PostDeviceUpsert default  %+v", o._statusCode, o.Payload)
}

func (o *PostDeviceUpsertDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostDeviceUpsertDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostDeviceUpsertOKBody post device upsert o k body
swagger:model PostDeviceUpsertOKBody
*/
type PostDeviceUpsertOKBody struct {
	models.Envelope

	// data
	Data *models.Fid `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostDeviceUpsertOKBody) UnmarshalJSON(raw []byte) error {
	// PostDeviceUpsertOKBodyAO0
	var postDeviceUpsertOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postDeviceUpsertOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postDeviceUpsertOKBodyAO0

	// PostDeviceUpsertOKBodyAO1
	var dataPostDeviceUpsertOKBodyAO1 struct {
		Data *models.Fid `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostDeviceUpsertOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostDeviceUpsertOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostDeviceUpsertOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postDeviceUpsertOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postDeviceUpsertOKBodyAO0)
	var dataPostDeviceUpsertOKBodyAO1 struct {
		Data *models.Fid `json:"data,omitempty"`
	}

	dataPostDeviceUpsertOKBodyAO1.Data = o.Data

	jsonDataPostDeviceUpsertOKBodyAO1, errPostDeviceUpsertOKBodyAO1 := swag.WriteJSON(dataPostDeviceUpsertOKBodyAO1)
	if errPostDeviceUpsertOKBodyAO1 != nil {
		return nil, errPostDeviceUpsertOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostDeviceUpsertOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post device upsert o k body
func (o *PostDeviceUpsertOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostDeviceUpsertOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postDeviceUpsertOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postDeviceUpsertOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post device upsert o k body based on the context it is used
func (o *PostDeviceUpsertOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostDeviceUpsertOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postDeviceUpsertOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postDeviceUpsertOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostDeviceUpsertOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostDeviceUpsertOKBody) UnmarshalBinary(b []byte) error {
	var res PostDeviceUpsertOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
