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

// GetCustomersFindByEmailReader is a Reader for the GetCustomersFindByEmail structure.
type GetCustomersFindByEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersFindByEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomersFindByEmailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCustomersFindByEmailDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersFindByEmailOK creates a GetCustomersFindByEmailOK with default headers values
func NewGetCustomersFindByEmailOK() *GetCustomersFindByEmailOK {
	return &GetCustomersFindByEmailOK{}
}

/* GetCustomersFindByEmailOK describes a response with status code 200, with default header values.

Located Customer
*/
type GetCustomersFindByEmailOK struct {
	Payload *GetCustomersFindByEmailOKBody
}

func (o *GetCustomersFindByEmailOK) Error() string {
	return fmt.Sprintf("[GET /customers/findByEmail][%d] getCustomersFindByEmailOK  %+v", 200, o.Payload)
}
func (o *GetCustomersFindByEmailOK) GetPayload() *GetCustomersFindByEmailOKBody {
	return o.Payload
}

func (o *GetCustomersFindByEmailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCustomersFindByEmailOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersFindByEmailDefault creates a GetCustomersFindByEmailDefault with default headers values
func NewGetCustomersFindByEmailDefault(code int) *GetCustomersFindByEmailDefault {
	return &GetCustomersFindByEmailDefault{
		_statusCode: code,
	}
}

/* GetCustomersFindByEmailDefault describes a response with status code -1, with default header values.

Error
*/
type GetCustomersFindByEmailDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get customers find by email default response
func (o *GetCustomersFindByEmailDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomersFindByEmailDefault) Error() string {
	return fmt.Sprintf("[GET /customers/findByEmail][%d] GetCustomersFindByEmail default  %+v", o._statusCode, o.Payload)
}
func (o *GetCustomersFindByEmailDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetCustomersFindByEmailDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCustomersFindByEmailOKBody get customers find by email o k body
swagger:model GetCustomersFindByEmailOKBody
*/
type GetCustomersFindByEmailOKBody struct {
	models.Envelope

	// data
	Data *models.Customer `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetCustomersFindByEmailOKBody) UnmarshalJSON(raw []byte) error {
	// GetCustomersFindByEmailOKBodyAO0
	var getCustomersFindByEmailOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getCustomersFindByEmailOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getCustomersFindByEmailOKBodyAO0

	// GetCustomersFindByEmailOKBodyAO1
	var dataGetCustomersFindByEmailOKBodyAO1 struct {
		Data *models.Customer `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetCustomersFindByEmailOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetCustomersFindByEmailOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetCustomersFindByEmailOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getCustomersFindByEmailOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getCustomersFindByEmailOKBodyAO0)
	var dataGetCustomersFindByEmailOKBodyAO1 struct {
		Data *models.Customer `json:"data,omitempty"`
	}

	dataGetCustomersFindByEmailOKBodyAO1.Data = o.Data

	jsonDataGetCustomersFindByEmailOKBodyAO1, errGetCustomersFindByEmailOKBodyAO1 := swag.WriteJSON(dataGetCustomersFindByEmailOKBodyAO1)
	if errGetCustomersFindByEmailOKBodyAO1 != nil {
		return nil, errGetCustomersFindByEmailOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetCustomersFindByEmailOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get customers find by email o k body
func (o *GetCustomersFindByEmailOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCustomersFindByEmailOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCustomersFindByEmailOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get customers find by email o k body based on the context it is used
func (o *GetCustomersFindByEmailOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetCustomersFindByEmailOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCustomersFindByEmailOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCustomersFindByEmailOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCustomersFindByEmailOKBody) UnmarshalBinary(b []byte) error {
	var res GetCustomersFindByEmailOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
