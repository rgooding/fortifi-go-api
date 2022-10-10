// Code generated by go-swagger; DO NOT EDIT.

package contacts

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

// GetContactsFindByReferenceContactReferenceReader is a Reader for the GetContactsFindByReferenceContactReference structure.
type GetContactsFindByReferenceContactReferenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContactsFindByReferenceContactReferenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContactsFindByReferenceContactReferenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetContactsFindByReferenceContactReferenceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetContactsFindByReferenceContactReferenceOK creates a GetContactsFindByReferenceContactReferenceOK with default headers values
func NewGetContactsFindByReferenceContactReferenceOK() *GetContactsFindByReferenceContactReferenceOK {
	return &GetContactsFindByReferenceContactReferenceOK{}
}

/*
GetContactsFindByReferenceContactReferenceOK describes a response with status code 200, with default header values.

Person info
*/
type GetContactsFindByReferenceContactReferenceOK struct {
	Payload *GetContactsFindByReferenceContactReferenceOKBody
}

// IsSuccess returns true when this get contacts find by reference contact reference o k response has a 2xx status code
func (o *GetContactsFindByReferenceContactReferenceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get contacts find by reference contact reference o k response has a 3xx status code
func (o *GetContactsFindByReferenceContactReferenceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contacts find by reference contact reference o k response has a 4xx status code
func (o *GetContactsFindByReferenceContactReferenceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contacts find by reference contact reference o k response has a 5xx status code
func (o *GetContactsFindByReferenceContactReferenceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get contacts find by reference contact reference o k response a status code equal to that given
func (o *GetContactsFindByReferenceContactReferenceOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetContactsFindByReferenceContactReferenceOK) Error() string {
	return fmt.Sprintf("[GET /contacts/findByReference/{contactReference}][%d] getContactsFindByReferenceContactReferenceOK  %+v", 200, o.Payload)
}

func (o *GetContactsFindByReferenceContactReferenceOK) String() string {
	return fmt.Sprintf("[GET /contacts/findByReference/{contactReference}][%d] getContactsFindByReferenceContactReferenceOK  %+v", 200, o.Payload)
}

func (o *GetContactsFindByReferenceContactReferenceOK) GetPayload() *GetContactsFindByReferenceContactReferenceOKBody {
	return o.Payload
}

func (o *GetContactsFindByReferenceContactReferenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetContactsFindByReferenceContactReferenceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContactsFindByReferenceContactReferenceDefault creates a GetContactsFindByReferenceContactReferenceDefault with default headers values
func NewGetContactsFindByReferenceContactReferenceDefault(code int) *GetContactsFindByReferenceContactReferenceDefault {
	return &GetContactsFindByReferenceContactReferenceDefault{
		_statusCode: code,
	}
}

/*
GetContactsFindByReferenceContactReferenceDefault describes a response with status code -1, with default header values.

Error
*/
type GetContactsFindByReferenceContactReferenceDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the get contacts find by reference contact reference default response
func (o *GetContactsFindByReferenceContactReferenceDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get contacts find by reference contact reference default response has a 2xx status code
func (o *GetContactsFindByReferenceContactReferenceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get contacts find by reference contact reference default response has a 3xx status code
func (o *GetContactsFindByReferenceContactReferenceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get contacts find by reference contact reference default response has a 4xx status code
func (o *GetContactsFindByReferenceContactReferenceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get contacts find by reference contact reference default response has a 5xx status code
func (o *GetContactsFindByReferenceContactReferenceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get contacts find by reference contact reference default response a status code equal to that given
func (o *GetContactsFindByReferenceContactReferenceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetContactsFindByReferenceContactReferenceDefault) Error() string {
	return fmt.Sprintf("[GET /contacts/findByReference/{contactReference}][%d] GetContactsFindByReferenceContactReference default  %+v", o._statusCode, o.Payload)
}

func (o *GetContactsFindByReferenceContactReferenceDefault) String() string {
	return fmt.Sprintf("[GET /contacts/findByReference/{contactReference}][%d] GetContactsFindByReferenceContactReference default  %+v", o._statusCode, o.Payload)
}

func (o *GetContactsFindByReferenceContactReferenceDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetContactsFindByReferenceContactReferenceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetContactsFindByReferenceContactReferenceOKBody get contacts find by reference contact reference o k body
swagger:model GetContactsFindByReferenceContactReferenceOKBody
*/
type GetContactsFindByReferenceContactReferenceOKBody struct {
	models.Envelope

	// data
	Data *models.Person `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetContactsFindByReferenceContactReferenceOKBody) UnmarshalJSON(raw []byte) error {
	// GetContactsFindByReferenceContactReferenceOKBodyAO0
	var getContactsFindByReferenceContactReferenceOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getContactsFindByReferenceContactReferenceOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getContactsFindByReferenceContactReferenceOKBodyAO0

	// GetContactsFindByReferenceContactReferenceOKBodyAO1
	var dataGetContactsFindByReferenceContactReferenceOKBodyAO1 struct {
		Data *models.Person `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetContactsFindByReferenceContactReferenceOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetContactsFindByReferenceContactReferenceOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetContactsFindByReferenceContactReferenceOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getContactsFindByReferenceContactReferenceOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getContactsFindByReferenceContactReferenceOKBodyAO0)
	var dataGetContactsFindByReferenceContactReferenceOKBodyAO1 struct {
		Data *models.Person `json:"data,omitempty"`
	}

	dataGetContactsFindByReferenceContactReferenceOKBodyAO1.Data = o.Data

	jsonDataGetContactsFindByReferenceContactReferenceOKBodyAO1, errGetContactsFindByReferenceContactReferenceOKBodyAO1 := swag.WriteJSON(dataGetContactsFindByReferenceContactReferenceOKBodyAO1)
	if errGetContactsFindByReferenceContactReferenceOKBodyAO1 != nil {
		return nil, errGetContactsFindByReferenceContactReferenceOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetContactsFindByReferenceContactReferenceOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get contacts find by reference contact reference o k body
func (o *GetContactsFindByReferenceContactReferenceOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetContactsFindByReferenceContactReferenceOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getContactsFindByReferenceContactReferenceOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getContactsFindByReferenceContactReferenceOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get contacts find by reference contact reference o k body based on the context it is used
func (o *GetContactsFindByReferenceContactReferenceOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetContactsFindByReferenceContactReferenceOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getContactsFindByReferenceContactReferenceOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getContactsFindByReferenceContactReferenceOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetContactsFindByReferenceContactReferenceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetContactsFindByReferenceContactReferenceOKBody) UnmarshalBinary(b []byte) error {
	var res GetContactsFindByReferenceContactReferenceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
