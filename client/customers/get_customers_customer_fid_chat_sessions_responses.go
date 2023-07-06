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

// GetCustomersCustomerFidChatSessionsReader is a Reader for the GetCustomersCustomerFidChatSessions structure.
type GetCustomersCustomerFidChatSessionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersCustomerFidChatSessionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomersCustomerFidChatSessionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCustomersCustomerFidChatSessionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersCustomerFidChatSessionsOK creates a GetCustomersCustomerFidChatSessionsOK with default headers values
func NewGetCustomersCustomerFidChatSessionsOK() *GetCustomersCustomerFidChatSessionsOK {
	return &GetCustomersCustomerFidChatSessionsOK{}
}

/*
GetCustomersCustomerFidChatSessionsOK describes a response with status code 200, with default header values.

Chat sessions
*/
type GetCustomersCustomerFidChatSessionsOK struct {
	Payload *GetCustomersCustomerFidChatSessionsOKBody
}

// IsSuccess returns true when this get customers customer fid chat sessions o k response has a 2xx status code
func (o *GetCustomersCustomerFidChatSessionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get customers customer fid chat sessions o k response has a 3xx status code
func (o *GetCustomersCustomerFidChatSessionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get customers customer fid chat sessions o k response has a 4xx status code
func (o *GetCustomersCustomerFidChatSessionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get customers customer fid chat sessions o k response has a 5xx status code
func (o *GetCustomersCustomerFidChatSessionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get customers customer fid chat sessions o k response a status code equal to that given
func (o *GetCustomersCustomerFidChatSessionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get customers customer fid chat sessions o k response
func (o *GetCustomersCustomerFidChatSessionsOK) Code() int {
	return 200
}

func (o *GetCustomersCustomerFidChatSessionsOK) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/chatSessions][%d] getCustomersCustomerFidChatSessionsOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidChatSessionsOK) String() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/chatSessions][%d] getCustomersCustomerFidChatSessionsOK  %+v", 200, o.Payload)
}

func (o *GetCustomersCustomerFidChatSessionsOK) GetPayload() *GetCustomersCustomerFidChatSessionsOKBody {
	return o.Payload
}

func (o *GetCustomersCustomerFidChatSessionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCustomersCustomerFidChatSessionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomersCustomerFidChatSessionsDefault creates a GetCustomersCustomerFidChatSessionsDefault with default headers values
func NewGetCustomersCustomerFidChatSessionsDefault(code int) *GetCustomersCustomerFidChatSessionsDefault {
	return &GetCustomersCustomerFidChatSessionsDefault{
		_statusCode: code,
	}
}

/*
GetCustomersCustomerFidChatSessionsDefault describes a response with status code -1, with default header values.

Error
*/
type GetCustomersCustomerFidChatSessionsDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this get customers customer fid chat sessions default response has a 2xx status code
func (o *GetCustomersCustomerFidChatSessionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get customers customer fid chat sessions default response has a 3xx status code
func (o *GetCustomersCustomerFidChatSessionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get customers customer fid chat sessions default response has a 4xx status code
func (o *GetCustomersCustomerFidChatSessionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get customers customer fid chat sessions default response has a 5xx status code
func (o *GetCustomersCustomerFidChatSessionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get customers customer fid chat sessions default response a status code equal to that given
func (o *GetCustomersCustomerFidChatSessionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get customers customer fid chat sessions default response
func (o *GetCustomersCustomerFidChatSessionsDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomersCustomerFidChatSessionsDefault) Error() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/chatSessions][%d] GetCustomersCustomerFidChatSessions default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidChatSessionsDefault) String() string {
	return fmt.Sprintf("[GET /customers/{customerFid}/chatSessions][%d] GetCustomersCustomerFidChatSessions default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersCustomerFidChatSessionsDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetCustomersCustomerFidChatSessionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetCustomersCustomerFidChatSessionsOKBody get customers customer fid chat sessions o k body
swagger:model GetCustomersCustomerFidChatSessionsOKBody
*/
type GetCustomersCustomerFidChatSessionsOKBody struct {
	models.Envelope

	// data
	Data *models.ChatSessions `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetCustomersCustomerFidChatSessionsOKBody) UnmarshalJSON(raw []byte) error {
	// GetCustomersCustomerFidChatSessionsOKBodyAO0
	var getCustomersCustomerFidChatSessionsOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getCustomersCustomerFidChatSessionsOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getCustomersCustomerFidChatSessionsOKBodyAO0

	// GetCustomersCustomerFidChatSessionsOKBodyAO1
	var dataGetCustomersCustomerFidChatSessionsOKBodyAO1 struct {
		Data *models.ChatSessions `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetCustomersCustomerFidChatSessionsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetCustomersCustomerFidChatSessionsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetCustomersCustomerFidChatSessionsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getCustomersCustomerFidChatSessionsOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getCustomersCustomerFidChatSessionsOKBodyAO0)
	var dataGetCustomersCustomerFidChatSessionsOKBodyAO1 struct {
		Data *models.ChatSessions `json:"data,omitempty"`
	}

	dataGetCustomersCustomerFidChatSessionsOKBodyAO1.Data = o.Data

	jsonDataGetCustomersCustomerFidChatSessionsOKBodyAO1, errGetCustomersCustomerFidChatSessionsOKBodyAO1 := swag.WriteJSON(dataGetCustomersCustomerFidChatSessionsOKBodyAO1)
	if errGetCustomersCustomerFidChatSessionsOKBodyAO1 != nil {
		return nil, errGetCustomersCustomerFidChatSessionsOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetCustomersCustomerFidChatSessionsOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get customers customer fid chat sessions o k body
func (o *GetCustomersCustomerFidChatSessionsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetCustomersCustomerFidChatSessionsOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCustomersCustomerFidChatSessionsOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getCustomersCustomerFidChatSessionsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get customers customer fid chat sessions o k body based on the context it is used
func (o *GetCustomersCustomerFidChatSessionsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetCustomersCustomerFidChatSessionsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCustomersCustomerFidChatSessionsOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getCustomersCustomerFidChatSessionsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCustomersCustomerFidChatSessionsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCustomersCustomerFidChatSessionsOKBody) UnmarshalBinary(b []byte) error {
	var res GetCustomersCustomerFidChatSessionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
