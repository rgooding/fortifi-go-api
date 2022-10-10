// Code generated by go-swagger; DO NOT EDIT.

package properties

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

// GetPropertiesReader is a Reader for the GetProperties structure.
type GetPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPropertiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPropertiesOK creates a GetPropertiesOK with default headers values
func NewGetPropertiesOK() *GetPropertiesOK {
	return &GetPropertiesOK{}
}

/*
GetPropertiesOK describes a response with status code 200, with default header values.

Defined Properties
*/
type GetPropertiesOK struct {
	Payload *GetPropertiesOKBody
}

// IsSuccess returns true when this get properties o k response has a 2xx status code
func (o *GetPropertiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get properties o k response has a 3xx status code
func (o *GetPropertiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get properties o k response has a 4xx status code
func (o *GetPropertiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get properties o k response has a 5xx status code
func (o *GetPropertiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get properties o k response a status code equal to that given
func (o *GetPropertiesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPropertiesOK) Error() string {
	return fmt.Sprintf("[GET /properties][%d] getPropertiesOK  %+v", 200, o.Payload)
}

func (o *GetPropertiesOK) String() string {
	return fmt.Sprintf("[GET /properties][%d] getPropertiesOK  %+v", 200, o.Payload)
}

func (o *GetPropertiesOK) GetPayload() *GetPropertiesOKBody {
	return o.Payload
}

func (o *GetPropertiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPropertiesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPropertiesDefault creates a GetPropertiesDefault with default headers values
func NewGetPropertiesDefault(code int) *GetPropertiesDefault {
	return &GetPropertiesDefault{
		_statusCode: code,
	}
}

/*
GetPropertiesDefault describes a response with status code -1, with default header values.

Error
*/
type GetPropertiesDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get properties default response
func (o *GetPropertiesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get properties default response has a 2xx status code
func (o *GetPropertiesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get properties default response has a 3xx status code
func (o *GetPropertiesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get properties default response has a 4xx status code
func (o *GetPropertiesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get properties default response has a 5xx status code
func (o *GetPropertiesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get properties default response a status code equal to that given
func (o *GetPropertiesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetPropertiesDefault) Error() string {
	return fmt.Sprintf("[GET /properties][%d] GetProperties default  %+v", o._statusCode, o.Payload)
}

func (o *GetPropertiesDefault) String() string {
	return fmt.Sprintf("[GET /properties][%d] GetProperties default  %+v", o._statusCode, o.Payload)
}

func (o *GetPropertiesDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetPropertiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetPropertiesOKBody get properties o k body
swagger:model GetPropertiesOKBody
*/
type GetPropertiesOKBody struct {
	models.Envelope

	// data
	Data *models.PropertyDefinitions `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetPropertiesOKBody) UnmarshalJSON(raw []byte) error {
	// GetPropertiesOKBodyAO0
	var getPropertiesOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getPropertiesOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getPropertiesOKBodyAO0

	// GetPropertiesOKBodyAO1
	var dataGetPropertiesOKBodyAO1 struct {
		Data *models.PropertyDefinitions `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetPropertiesOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetPropertiesOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetPropertiesOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getPropertiesOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getPropertiesOKBodyAO0)
	var dataGetPropertiesOKBodyAO1 struct {
		Data *models.PropertyDefinitions `json:"data,omitempty"`
	}

	dataGetPropertiesOKBodyAO1.Data = o.Data

	jsonDataGetPropertiesOKBodyAO1, errGetPropertiesOKBodyAO1 := swag.WriteJSON(dataGetPropertiesOKBodyAO1)
	if errGetPropertiesOKBodyAO1 != nil {
		return nil, errGetPropertiesOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetPropertiesOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get properties o k body
func (o *GetPropertiesOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetPropertiesOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPropertiesOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getPropertiesOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get properties o k body based on the context it is used
func (o *GetPropertiesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetPropertiesOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPropertiesOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getPropertiesOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPropertiesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPropertiesOKBody) UnmarshalBinary(b []byte) error {
	var res GetPropertiesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
