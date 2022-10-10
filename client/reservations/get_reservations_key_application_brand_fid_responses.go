// Code generated by go-swagger; DO NOT EDIT.

package reservations

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

// GetReservationsKeyApplicationBrandFidReader is a Reader for the GetReservationsKeyApplicationBrandFid structure.
type GetReservationsKeyApplicationBrandFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReservationsKeyApplicationBrandFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReservationsKeyApplicationBrandFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetReservationsKeyApplicationBrandFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetReservationsKeyApplicationBrandFidOK creates a GetReservationsKeyApplicationBrandFidOK with default headers values
func NewGetReservationsKeyApplicationBrandFidOK() *GetReservationsKeyApplicationBrandFidOK {
	return &GetReservationsKeyApplicationBrandFidOK{}
}

/*
GetReservationsKeyApplicationBrandFidOK describes a response with status code 200, with default header values.

Pixels
*/
type GetReservationsKeyApplicationBrandFidOK struct {
	Payload *GetReservationsKeyApplicationBrandFidOKBody
}

// IsSuccess returns true when this get reservations key application brand fid o k response has a 2xx status code
func (o *GetReservationsKeyApplicationBrandFidOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get reservations key application brand fid o k response has a 3xx status code
func (o *GetReservationsKeyApplicationBrandFidOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get reservations key application brand fid o k response has a 4xx status code
func (o *GetReservationsKeyApplicationBrandFidOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get reservations key application brand fid o k response has a 5xx status code
func (o *GetReservationsKeyApplicationBrandFidOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get reservations key application brand fid o k response a status code equal to that given
func (o *GetReservationsKeyApplicationBrandFidOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetReservationsKeyApplicationBrandFidOK) Error() string {
	return fmt.Sprintf("[GET /reservations/{key}/{application}/{brandFid}][%d] getReservationsKeyApplicationBrandFidOK  %+v", 200, o.Payload)
}

func (o *GetReservationsKeyApplicationBrandFidOK) String() string {
	return fmt.Sprintf("[GET /reservations/{key}/{application}/{brandFid}][%d] getReservationsKeyApplicationBrandFidOK  %+v", 200, o.Payload)
}

func (o *GetReservationsKeyApplicationBrandFidOK) GetPayload() *GetReservationsKeyApplicationBrandFidOKBody {
	return o.Payload
}

func (o *GetReservationsKeyApplicationBrandFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetReservationsKeyApplicationBrandFidOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReservationsKeyApplicationBrandFidDefault creates a GetReservationsKeyApplicationBrandFidDefault with default headers values
func NewGetReservationsKeyApplicationBrandFidDefault(code int) *GetReservationsKeyApplicationBrandFidDefault {
	return &GetReservationsKeyApplicationBrandFidDefault{
		_statusCode: code,
	}
}

/*
GetReservationsKeyApplicationBrandFidDefault describes a response with status code -1, with default header values.

Error
*/
type GetReservationsKeyApplicationBrandFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get reservations key application brand fid default response
func (o *GetReservationsKeyApplicationBrandFidDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get reservations key application brand fid default response has a 2xx status code
func (o *GetReservationsKeyApplicationBrandFidDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get reservations key application brand fid default response has a 3xx status code
func (o *GetReservationsKeyApplicationBrandFidDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get reservations key application brand fid default response has a 4xx status code
func (o *GetReservationsKeyApplicationBrandFidDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get reservations key application brand fid default response has a 5xx status code
func (o *GetReservationsKeyApplicationBrandFidDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get reservations key application brand fid default response a status code equal to that given
func (o *GetReservationsKeyApplicationBrandFidDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetReservationsKeyApplicationBrandFidDefault) Error() string {
	return fmt.Sprintf("[GET /reservations/{key}/{application}/{brandFid}][%d] GetReservationsKeyApplicationBrandFid default  %+v", o._statusCode, o.Payload)
}

func (o *GetReservationsKeyApplicationBrandFidDefault) String() string {
	return fmt.Sprintf("[GET /reservations/{key}/{application}/{brandFid}][%d] GetReservationsKeyApplicationBrandFid default  %+v", o._statusCode, o.Payload)
}

func (o *GetReservationsKeyApplicationBrandFidDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetReservationsKeyApplicationBrandFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetReservationsKeyApplicationBrandFidOKBody get reservations key application brand fid o k body
swagger:model GetReservationsKeyApplicationBrandFidOKBody
*/
type GetReservationsKeyApplicationBrandFidOKBody struct {
	models.Envelope

	// data
	Data *models.Reservations `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetReservationsKeyApplicationBrandFidOKBody) UnmarshalJSON(raw []byte) error {
	// GetReservationsKeyApplicationBrandFidOKBodyAO0
	var getReservationsKeyApplicationBrandFidOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getReservationsKeyApplicationBrandFidOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getReservationsKeyApplicationBrandFidOKBodyAO0

	// GetReservationsKeyApplicationBrandFidOKBodyAO1
	var dataGetReservationsKeyApplicationBrandFidOKBodyAO1 struct {
		Data *models.Reservations `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetReservationsKeyApplicationBrandFidOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetReservationsKeyApplicationBrandFidOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetReservationsKeyApplicationBrandFidOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getReservationsKeyApplicationBrandFidOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getReservationsKeyApplicationBrandFidOKBodyAO0)
	var dataGetReservationsKeyApplicationBrandFidOKBodyAO1 struct {
		Data *models.Reservations `json:"data,omitempty"`
	}

	dataGetReservationsKeyApplicationBrandFidOKBodyAO1.Data = o.Data

	jsonDataGetReservationsKeyApplicationBrandFidOKBodyAO1, errGetReservationsKeyApplicationBrandFidOKBodyAO1 := swag.WriteJSON(dataGetReservationsKeyApplicationBrandFidOKBodyAO1)
	if errGetReservationsKeyApplicationBrandFidOKBodyAO1 != nil {
		return nil, errGetReservationsKeyApplicationBrandFidOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetReservationsKeyApplicationBrandFidOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get reservations key application brand fid o k body
func (o *GetReservationsKeyApplicationBrandFidOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetReservationsKeyApplicationBrandFidOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getReservationsKeyApplicationBrandFidOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getReservationsKeyApplicationBrandFidOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get reservations key application brand fid o k body based on the context it is used
func (o *GetReservationsKeyApplicationBrandFidOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetReservationsKeyApplicationBrandFidOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getReservationsKeyApplicationBrandFidOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getReservationsKeyApplicationBrandFidOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetReservationsKeyApplicationBrandFidOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetReservationsKeyApplicationBrandFidOKBody) UnmarshalBinary(b []byte) error {
	var res GetReservationsKeyApplicationBrandFidOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
