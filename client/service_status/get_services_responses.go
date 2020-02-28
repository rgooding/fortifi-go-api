// Code generated by go-swagger; DO NOT EDIT.

package service_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fortifi/go-api/models"
)

// GetServicesReader is a Reader for the GetServices structure.
type GetServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetServicesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServicesOK creates a GetServicesOK with default headers values
func NewGetServicesOK() *GetServicesOK {
	return &GetServicesOK{}
}

/*GetServicesOK handles this case with default header values.

Service listing
*/
type GetServicesOK struct {
	Payload *GetServicesOKBody
}

func (o *GetServicesOK) Error() string {
	return fmt.Sprintf("[GET /services][%d] getServicesOK  %+v", 200, o.Payload)
}

func (o *GetServicesOK) GetPayload() *GetServicesOKBody {
	return o.Payload
}

func (o *GetServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetServicesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServicesDefault creates a GetServicesDefault with default headers values
func NewGetServicesDefault(code int) *GetServicesDefault {
	return &GetServicesDefault{
		_statusCode: code,
	}
}

/*GetServicesDefault handles this case with default header values.

Error
*/
type GetServicesDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get services default response
func (o *GetServicesDefault) Code() int {
	return o._statusCode
}

func (o *GetServicesDefault) Error() string {
	return fmt.Sprintf("[GET /services][%d] GetServices default  %+v", o._statusCode, o.Payload)
}

func (o *GetServicesDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetServicesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetServicesOKBody get services o k body
swagger:model GetServicesOKBody
*/
type GetServicesOKBody struct {
	models.Envelope

	// data
	Data *models.Services `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetServicesOKBody) UnmarshalJSON(raw []byte) error {
	// GetServicesOKBodyAO0
	var getServicesOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getServicesOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getServicesOKBodyAO0

	// GetServicesOKBodyAO1
	var dataGetServicesOKBodyAO1 struct {
		Data *models.Services `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetServicesOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetServicesOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetServicesOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getServicesOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getServicesOKBodyAO0)

	var dataGetServicesOKBodyAO1 struct {
		Data *models.Services `json:"data,omitempty"`
	}

	dataGetServicesOKBodyAO1.Data = o.Data

	jsonDataGetServicesOKBodyAO1, errGetServicesOKBodyAO1 := swag.WriteJSON(dataGetServicesOKBodyAO1)
	if errGetServicesOKBodyAO1 != nil {
		return nil, errGetServicesOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetServicesOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get services o k body
func (o *GetServicesOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetServicesOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServicesOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServicesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServicesOKBody) UnmarshalBinary(b []byte) error {
	var res GetServicesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
