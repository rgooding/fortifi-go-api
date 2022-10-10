// Code generated by go-swagger; DO NOT EDIT.

package polymers

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

// GetPolymersParentFidPolymerFidReader is a Reader for the GetPolymersParentFidPolymerFid structure.
type GetPolymersParentFidPolymerFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPolymersParentFidPolymerFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPolymersParentFidPolymerFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPolymersParentFidPolymerFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPolymersParentFidPolymerFidOK creates a GetPolymersParentFidPolymerFidOK with default headers values
func NewGetPolymersParentFidPolymerFidOK() *GetPolymersParentFidPolymerFidOK {
	return &GetPolymersParentFidPolymerFidOK{}
}

/*
GetPolymersParentFidPolymerFidOK describes a response with status code 200, with default header values.

Polymer Information
*/
type GetPolymersParentFidPolymerFidOK struct {
	Payload *GetPolymersParentFidPolymerFidOKBody
}

// IsSuccess returns true when this get polymers parent fid polymer fid o k response has a 2xx status code
func (o *GetPolymersParentFidPolymerFidOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get polymers parent fid polymer fid o k response has a 3xx status code
func (o *GetPolymersParentFidPolymerFidOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get polymers parent fid polymer fid o k response has a 4xx status code
func (o *GetPolymersParentFidPolymerFidOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get polymers parent fid polymer fid o k response has a 5xx status code
func (o *GetPolymersParentFidPolymerFidOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get polymers parent fid polymer fid o k response a status code equal to that given
func (o *GetPolymersParentFidPolymerFidOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPolymersParentFidPolymerFidOK) Error() string {
	return fmt.Sprintf("[GET /polymers/{parentFid}/{polymerFid}][%d] getPolymersParentFidPolymerFidOK  %+v", 200, o.Payload)
}

func (o *GetPolymersParentFidPolymerFidOK) String() string {
	return fmt.Sprintf("[GET /polymers/{parentFid}/{polymerFid}][%d] getPolymersParentFidPolymerFidOK  %+v", 200, o.Payload)
}

func (o *GetPolymersParentFidPolymerFidOK) GetPayload() *GetPolymersParentFidPolymerFidOKBody {
	return o.Payload
}

func (o *GetPolymersParentFidPolymerFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPolymersParentFidPolymerFidOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPolymersParentFidPolymerFidDefault creates a GetPolymersParentFidPolymerFidDefault with default headers values
func NewGetPolymersParentFidPolymerFidDefault(code int) *GetPolymersParentFidPolymerFidDefault {
	return &GetPolymersParentFidPolymerFidDefault{
		_statusCode: code,
	}
}

/*
GetPolymersParentFidPolymerFidDefault describes a response with status code -1, with default header values.

Error
*/
type GetPolymersParentFidPolymerFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get polymers parent fid polymer fid default response
func (o *GetPolymersParentFidPolymerFidDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get polymers parent fid polymer fid default response has a 2xx status code
func (o *GetPolymersParentFidPolymerFidDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get polymers parent fid polymer fid default response has a 3xx status code
func (o *GetPolymersParentFidPolymerFidDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get polymers parent fid polymer fid default response has a 4xx status code
func (o *GetPolymersParentFidPolymerFidDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get polymers parent fid polymer fid default response has a 5xx status code
func (o *GetPolymersParentFidPolymerFidDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get polymers parent fid polymer fid default response a status code equal to that given
func (o *GetPolymersParentFidPolymerFidDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetPolymersParentFidPolymerFidDefault) Error() string {
	return fmt.Sprintf("[GET /polymers/{parentFid}/{polymerFid}][%d] GetPolymersParentFidPolymerFid default  %+v", o._statusCode, o.Payload)
}

func (o *GetPolymersParentFidPolymerFidDefault) String() string {
	return fmt.Sprintf("[GET /polymers/{parentFid}/{polymerFid}][%d] GetPolymersParentFidPolymerFid default  %+v", o._statusCode, o.Payload)
}

func (o *GetPolymersParentFidPolymerFidDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetPolymersParentFidPolymerFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetPolymersParentFidPolymerFidOKBody get polymers parent fid polymer fid o k body
swagger:model GetPolymersParentFidPolymerFidOKBody
*/
type GetPolymersParentFidPolymerFidOKBody struct {
	models.Envelope

	// data
	Data *models.Polymer `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetPolymersParentFidPolymerFidOKBody) UnmarshalJSON(raw []byte) error {
	// GetPolymersParentFidPolymerFidOKBodyAO0
	var getPolymersParentFidPolymerFidOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getPolymersParentFidPolymerFidOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getPolymersParentFidPolymerFidOKBodyAO0

	// GetPolymersParentFidPolymerFidOKBodyAO1
	var dataGetPolymersParentFidPolymerFidOKBodyAO1 struct {
		Data *models.Polymer `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetPolymersParentFidPolymerFidOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetPolymersParentFidPolymerFidOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetPolymersParentFidPolymerFidOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getPolymersParentFidPolymerFidOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getPolymersParentFidPolymerFidOKBodyAO0)
	var dataGetPolymersParentFidPolymerFidOKBodyAO1 struct {
		Data *models.Polymer `json:"data,omitempty"`
	}

	dataGetPolymersParentFidPolymerFidOKBodyAO1.Data = o.Data

	jsonDataGetPolymersParentFidPolymerFidOKBodyAO1, errGetPolymersParentFidPolymerFidOKBodyAO1 := swag.WriteJSON(dataGetPolymersParentFidPolymerFidOKBodyAO1)
	if errGetPolymersParentFidPolymerFidOKBodyAO1 != nil {
		return nil, errGetPolymersParentFidPolymerFidOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetPolymersParentFidPolymerFidOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get polymers parent fid polymer fid o k body
func (o *GetPolymersParentFidPolymerFidOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetPolymersParentFidPolymerFidOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPolymersParentFidPolymerFidOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getPolymersParentFidPolymerFidOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get polymers parent fid polymer fid o k body based on the context it is used
func (o *GetPolymersParentFidPolymerFidOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetPolymersParentFidPolymerFidOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPolymersParentFidPolymerFidOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getPolymersParentFidPolymerFidOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPolymersParentFidPolymerFidOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPolymersParentFidPolymerFidOKBody) UnmarshalBinary(b []byte) error {
	var res GetPolymersParentFidPolymerFidOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
