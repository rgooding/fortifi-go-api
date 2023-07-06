// Code generated by go-swagger; DO NOT EDIT.

package interactions

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

// GetInteractionsInteractionFidReader is a Reader for the GetInteractionsInteractionFid structure.
type GetInteractionsInteractionFidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInteractionsInteractionFidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInteractionsInteractionFidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetInteractionsInteractionFidDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetInteractionsInteractionFidOK creates a GetInteractionsInteractionFidOK with default headers values
func NewGetInteractionsInteractionFidOK() *GetInteractionsInteractionFidOK {
	return &GetInteractionsInteractionFidOK{}
}

/*
GetInteractionsInteractionFidOK describes a response with status code 200, with default header values.

Interaction
*/
type GetInteractionsInteractionFidOK struct {
	Payload *GetInteractionsInteractionFidOKBody
}

// IsSuccess returns true when this get interactions interaction fid o k response has a 2xx status code
func (o *GetInteractionsInteractionFidOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get interactions interaction fid o k response has a 3xx status code
func (o *GetInteractionsInteractionFidOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get interactions interaction fid o k response has a 4xx status code
func (o *GetInteractionsInteractionFidOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get interactions interaction fid o k response has a 5xx status code
func (o *GetInteractionsInteractionFidOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get interactions interaction fid o k response a status code equal to that given
func (o *GetInteractionsInteractionFidOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get interactions interaction fid o k response
func (o *GetInteractionsInteractionFidOK) Code() int {
	return 200
}

func (o *GetInteractionsInteractionFidOK) Error() string {
	return fmt.Sprintf("[GET /interactions/{interactionFid}][%d] getInteractionsInteractionFidOK  %+v", 200, o.Payload)
}

func (o *GetInteractionsInteractionFidOK) String() string {
	return fmt.Sprintf("[GET /interactions/{interactionFid}][%d] getInteractionsInteractionFidOK  %+v", 200, o.Payload)
}

func (o *GetInteractionsInteractionFidOK) GetPayload() *GetInteractionsInteractionFidOKBody {
	return o.Payload
}

func (o *GetInteractionsInteractionFidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetInteractionsInteractionFidOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInteractionsInteractionFidDefault creates a GetInteractionsInteractionFidDefault with default headers values
func NewGetInteractionsInteractionFidDefault(code int) *GetInteractionsInteractionFidDefault {
	return &GetInteractionsInteractionFidDefault{
		_statusCode: code,
	}
}

/*
GetInteractionsInteractionFidDefault describes a response with status code -1, with default header values.

Error
*/
type GetInteractionsInteractionFidDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this get interactions interaction fid default response has a 2xx status code
func (o *GetInteractionsInteractionFidDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get interactions interaction fid default response has a 3xx status code
func (o *GetInteractionsInteractionFidDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get interactions interaction fid default response has a 4xx status code
func (o *GetInteractionsInteractionFidDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get interactions interaction fid default response has a 5xx status code
func (o *GetInteractionsInteractionFidDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get interactions interaction fid default response a status code equal to that given
func (o *GetInteractionsInteractionFidDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get interactions interaction fid default response
func (o *GetInteractionsInteractionFidDefault) Code() int {
	return o._statusCode
}

func (o *GetInteractionsInteractionFidDefault) Error() string {
	return fmt.Sprintf("[GET /interactions/{interactionFid}][%d] GetInteractionsInteractionFid default  %+v", o._statusCode, o.Payload)
}

func (o *GetInteractionsInteractionFidDefault) String() string {
	return fmt.Sprintf("[GET /interactions/{interactionFid}][%d] GetInteractionsInteractionFid default  %+v", o._statusCode, o.Payload)
}

func (o *GetInteractionsInteractionFidDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *GetInteractionsInteractionFidDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetInteractionsInteractionFidOKBody get interactions interaction fid o k body
swagger:model GetInteractionsInteractionFidOKBody
*/
type GetInteractionsInteractionFidOKBody struct {
	models.Envelope

	// data
	Data *models.InteractionResponse `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetInteractionsInteractionFidOKBody) UnmarshalJSON(raw []byte) error {
	// GetInteractionsInteractionFidOKBodyAO0
	var getInteractionsInteractionFidOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &getInteractionsInteractionFidOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = getInteractionsInteractionFidOKBodyAO0

	// GetInteractionsInteractionFidOKBodyAO1
	var dataGetInteractionsInteractionFidOKBodyAO1 struct {
		Data *models.InteractionResponse `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetInteractionsInteractionFidOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetInteractionsInteractionFidOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetInteractionsInteractionFidOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getInteractionsInteractionFidOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getInteractionsInteractionFidOKBodyAO0)
	var dataGetInteractionsInteractionFidOKBodyAO1 struct {
		Data *models.InteractionResponse `json:"data,omitempty"`
	}

	dataGetInteractionsInteractionFidOKBodyAO1.Data = o.Data

	jsonDataGetInteractionsInteractionFidOKBodyAO1, errGetInteractionsInteractionFidOKBodyAO1 := swag.WriteJSON(dataGetInteractionsInteractionFidOKBodyAO1)
	if errGetInteractionsInteractionFidOKBodyAO1 != nil {
		return nil, errGetInteractionsInteractionFidOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetInteractionsInteractionFidOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get interactions interaction fid o k body
func (o *GetInteractionsInteractionFidOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetInteractionsInteractionFidOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getInteractionsInteractionFidOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getInteractionsInteractionFidOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get interactions interaction fid o k body based on the context it is used
func (o *GetInteractionsInteractionFidOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetInteractionsInteractionFidOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getInteractionsInteractionFidOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getInteractionsInteractionFidOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetInteractionsInteractionFidOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInteractionsInteractionFidOKBody) UnmarshalBinary(b []byte) error {
	var res GetInteractionsInteractionFidOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}