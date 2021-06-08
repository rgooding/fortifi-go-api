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

// GetCustomersCustomerFidTicketsReader is a Reader for the GetCustomersCustomerFidTickets structure.
type GetCustomersCustomerFidTicketsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersCustomerFidTicketsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomersCustomerFidTicketsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCustomersCustomerFidTicketsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersCustomerFidTicketsOK creates a GetCustomersCustomerFidTicketsOK with default headers values
func NewGetCustomersCustomerFidTicketsOK() *GetCustomersCustomerFidTicketsOK {
	return &GetCustomersCustomerFidTicketsOK{}
}

/* GetCustomersCustomerFidTicketsOK describes a response with status code 200, with default header values.

Ticket collection
*/
type GetCustomersCustomerFidTicketsOK struct {
	Payload *GetCustomersCustomerFidTicketsOKBody
}

func (o *GetCustomersCustomerFidTicketsOK) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/tickets][%d] getCustomersCustomerFidTicketsOK  %+v", 200, o.Payload)
}
func (o *GetCustomersCustomerFidTicketsOK) GetPayload() *GetCustomersCustomerFidTicketsOKBody {
	return o.Payload
}

func (o *GetCustomersCustomerFidTicketsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCustomersCustomerFidTicketsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersCustomerFidTicketsDefault creates a GetCustomersCustomerFidTicketsDefault with default headers values
func NewGetCustomersCustomerFidTicketsDefault(code int) *GetCustomersCustomerFidTicketsDefault {
	return &GetCustomersCustomerFidTicketsDefault{
		_statusCode: code,
	}
}

/* GetCustomersCustomerFidTicketsDefault describes a response with status code -1, with default header values.

Error
*/
type GetCustomersCustomerFidTicketsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get customers customer fid tickets default response
func (o *GetCustomersCustomerFidTicketsDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomersCustomerFidTicketsDefault) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/tickets][%d] GetCustomersCustomerFidTickets default  %+v", o._statusCode, o.Payload)
}
func (o *GetCustomersCustomerFidTicketsDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetCustomersCustomerFidTicketsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCustomersCustomerFidTicketsOKBody get customers customer fid tickets o k body
swagger:model GetCustomersCustomerFidTicketsOKBody
*/
type GetCustomersCustomerFidTicketsOKBody struct {
	models.Envelope

	// data
	Data *models.Tickets `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetCustomersCustomerFidTicketsOKBody) UnmarshalJSON(raw []byte) error {
	// GetCustomersCustomerFidTicketsOKBodyAO0
	var getCustomersCustomerFidTicketsOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getCustomersCustomerFidTicketsOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getCustomersCustomerFidTicketsOKBodyAO0

	// GetCustomersCustomerFidTicketsOKBodyAO1
	var dataGetCustomersCustomerFidTicketsOKBodyAO1 struct {
		Data *models.Tickets `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetCustomersCustomerFidTicketsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetCustomersCustomerFidTicketsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetCustomersCustomerFidTicketsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getCustomersCustomerFidTicketsOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getCustomersCustomerFidTicketsOKBodyAO0)
	var dataGetCustomersCustomerFidTicketsOKBodyAO1 struct {
		Data *models.Tickets `json:"data,omitempty"`
	}

	dataGetCustomersCustomerFidTicketsOKBodyAO1.Data = o.Data

	jsonDataGetCustomersCustomerFidTicketsOKBodyAO1, errGetCustomersCustomerFidTicketsOKBodyAO1 := swag.WriteJSON(dataGetCustomersCustomerFidTicketsOKBodyAO1)
	if errGetCustomersCustomerFidTicketsOKBodyAO1 != nil {
		return nil, errGetCustomersCustomerFidTicketsOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetCustomersCustomerFidTicketsOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get customers customer fid tickets o k body
func (o *GetCustomersCustomerFidTicketsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCustomersCustomerFidTicketsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCustomersCustomerFidTicketsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get customers customer fid tickets o k body based on the context it is used
func (o *GetCustomersCustomerFidTicketsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetCustomersCustomerFidTicketsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCustomersCustomerFidTicketsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCustomersCustomerFidTicketsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCustomersCustomerFidTicketsOKBody) UnmarshalBinary(b []byte) error {
	var res GetCustomersCustomerFidTicketsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
