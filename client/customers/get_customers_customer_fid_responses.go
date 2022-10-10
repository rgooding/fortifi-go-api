// Code generated by go-swagger; DO NOT EDIT.

package customers

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

// GetCustomersCustomerFidReader is a Reader for the GetCustomersCustomerFid structure.
type GetCustomersCustomerFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersCustomerFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomersCustomerFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCustomersCustomerFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersCustomerFidOK creates a GetCustomersCustomerFidOK with default headers values
func NewGetCustomersCustomerFidOK() *GetCustomersCustomerFidOK {
	return &GetCustomersCustomerFidOK{}
}

/*
GetCustomersCustomerFidOK describes a response with status code 200, with default header values.

Loaded Customer
*/
type GetCustomersCustomerFidOK struct {
	Payload *GetCustomersCustomerFidOKBody
}

// IsSuccess returns true when this get customers customer fid o k response has a 2xx status code
func (o *GetCustomersCustomerFidOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get customers customer fid o k response has a 3xx status code
func (o *GetCustomersCustomerFidOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customers customer fid o k response has a 4xx status code
func (o *GetCustomersCustomerFidOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get customers customer fid o k response has a 5xx status code
func (o *GetCustomersCustomerFidOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get customers customer fid o k response a status code equal to that given
func (o *GetCustomersCustomerFidOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCustomersCustomerFidOK) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}][%d] getCustomersCustomerFidOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidOK) String() string {
	return fmt.Sprintf("[GET /customers/{customerFid}][%d] getCustomersCustomerFidOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidOK) GetPayload() *GetCustomersCustomerFidOKBody {
	return o.Payload
}

func (o *GetCustomersCustomerFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCustomersCustomerFidOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersCustomerFidDefault creates a GetCustomersCustomerFidDefault with default headers values
func NewGetCustomersCustomerFidDefault(code int) *GetCustomersCustomerFidDefault {
	return &GetCustomersCustomerFidDefault{
		_statusCode: code,
	}
}

/*
GetCustomersCustomerFidDefault describes a response with status code -1, with default header values.

Error
*/
type GetCustomersCustomerFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get customers customer fid default response
func (o *GetCustomersCustomerFidDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get customers customer fid default response has a 2xx status code
func (o *GetCustomersCustomerFidDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get customers customer fid default response has a 3xx status code
func (o *GetCustomersCustomerFidDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get customers customer fid default response has a 4xx status code
func (o *GetCustomersCustomerFidDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get customers customer fid default response has a 5xx status code
func (o *GetCustomersCustomerFidDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get customers customer fid default response a status code equal to that given
func (o *GetCustomersCustomerFidDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetCustomersCustomerFidDefault) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}][%d] GetCustomersCustomerFid default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidDefault) String() string {
	return fmt.Sprintf("[GET /customers/{customerFid}][%d] GetCustomersCustomerFid default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetCustomersCustomerFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetCustomersCustomerFidOKBody get customers customer fid o k body
swagger:model GetCustomersCustomerFidOKBody
*/
type GetCustomersCustomerFidOKBody struct {
	models.Envelope

	// data
	Data *models.Customer `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetCustomersCustomerFidOKBody) UnmarshalJSON(raw []byte) error {
	// GetCustomersCustomerFidOKBodyAO0
	var getCustomersCustomerFidOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getCustomersCustomerFidOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getCustomersCustomerFidOKBodyAO0

	// GetCustomersCustomerFidOKBodyAO1
	var dataGetCustomersCustomerFidOKBodyAO1 struct {
		Data *models.Customer `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetCustomersCustomerFidOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetCustomersCustomerFidOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetCustomersCustomerFidOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getCustomersCustomerFidOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getCustomersCustomerFidOKBodyAO0)
	var dataGetCustomersCustomerFidOKBodyAO1 struct {
		Data *models.Customer `json:"data,omitempty"`
	}

	dataGetCustomersCustomerFidOKBodyAO1.Data = o.Data

	jsonDataGetCustomersCustomerFidOKBodyAO1, errGetCustomersCustomerFidOKBodyAO1 := swag.WriteJSON(dataGetCustomersCustomerFidOKBodyAO1)
	if errGetCustomersCustomerFidOKBodyAO1 != nil {
		return nil, errGetCustomersCustomerFidOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetCustomersCustomerFidOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get customers customer fid o k body
func (o *GetCustomersCustomerFidOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCustomersCustomerFidOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCustomersCustomerFidOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getCustomersCustomerFidOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get customers customer fid o k body based on the context it is used
func (o *GetCustomersCustomerFidOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetCustomersCustomerFidOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCustomersCustomerFidOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getCustomersCustomerFidOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCustomersCustomerFidOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCustomersCustomerFidOKBody) UnmarshalBinary(b []byte) error {
	var res GetCustomersCustomerFidOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
