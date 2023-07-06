// Code generated by go-swagger; DO NOT EDIT.

package accountant

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

// PostAccountantTransactionSearchReader is a Reader for the PostAccountantTransactionSearch structure.
type PostAccountantTransactionSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAccountantTransactionSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAccountantTransactionSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostAccountantTransactionSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostAccountantTransactionSearchOK creates a PostAccountantTransactionSearchOK with default headers values
func NewPostAccountantTransactionSearchOK() *PostAccountantTransactionSearchOK {
	return &PostAccountantTransactionSearchOK{}
}

/*
PostAccountantTransactionSearchOK describes a response with status code 200, with default header values.

Find Transaction Response
*/
type PostAccountantTransactionSearchOK struct {
	Payload *PostAccountantTransactionSearchOKBody
}

// IsSuccess returns true when this post accountant transaction search o k response has a 2xx status code
func (o *PostAccountantTransactionSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post accountant transaction search o k response has a 3xx status code
func (o *PostAccountantTransactionSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post accountant transaction search o k response has a 4xx status code
func (o *PostAccountantTransactionSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post accountant transaction search o k response has a 5xx status code
func (o *PostAccountantTransactionSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post accountant transaction search o k response a status code equal to that given
func (o *PostAccountantTransactionSearchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post accountant transaction search o k response
func (o *PostAccountantTransactionSearchOK) Code() int {
	return 200
}

func (o *PostAccountantTransactionSearchOK) Error() string {
	return fmt.Sprintf("[POST /accountant/transaction/search][%d] postAccountantTransactionSearchOK  %+v", 200, o.Payload)
}

func (o *PostAccountantTransactionSearchOK) String() string {
	return fmt.Sprintf("[POST /accountant/transaction/search][%d] postAccountantTransactionSearchOK  %+v", 200, o.Payload)
}

func (o *PostAccountantTransactionSearchOK) GetPayload() *PostAccountantTransactionSearchOKBody {
	return o.Payload
}

func (o *PostAccountantTransactionSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostAccountantTransactionSearchOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccountantTransactionSearchDefault creates a PostAccountantTransactionSearchDefault with default headers values
func NewPostAccountantTransactionSearchDefault(code int) *PostAccountantTransactionSearchDefault {
	return &PostAccountantTransactionSearchDefault{
		_statusCode: code,
	}
}

/*
PostAccountantTransactionSearchDefault describes a response with status code -1, with default header values.

Error
*/
type PostAccountantTransactionSearchDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this post accountant transaction search default response has a 2xx status code
func (o *PostAccountantTransactionSearchDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post accountant transaction search default response has a 3xx status code
func (o *PostAccountantTransactionSearchDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post accountant transaction search default response has a 4xx status code
func (o *PostAccountantTransactionSearchDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post accountant transaction search default response has a 5xx status code
func (o *PostAccountantTransactionSearchDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post accountant transaction search default response a status code equal to that given
func (o *PostAccountantTransactionSearchDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post accountant transaction search default response
func (o *PostAccountantTransactionSearchDefault) Code() int {
	return o._statusCode
}

func (o *PostAccountantTransactionSearchDefault) Error() string {
	return fmt.Sprintf("[POST /accountant/transaction/search][%d] PostAccountantTransactionSearch default  %+v", o._statusCode, o.Payload)
}

func (o *PostAccountantTransactionSearchDefault) String() string {
	return fmt.Sprintf("[POST /accountant/transaction/search][%d] PostAccountantTransactionSearch default  %+v", o._statusCode, o.Payload)
}

func (o *PostAccountantTransactionSearchDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostAccountantTransactionSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostAccountantTransactionSearchOKBody post accountant transaction search o k body
swagger:model PostAccountantTransactionSearchOKBody
*/
type PostAccountantTransactionSearchOKBody struct {
	models.Envelope

	// data
	Data *models.FindTransactionResponse `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostAccountantTransactionSearchOKBody) UnmarshalJSON(raw []byte) error {
	// PostAccountantTransactionSearchOKBodyAO0
	var postAccountantTransactionSearchOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postAccountantTransactionSearchOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postAccountantTransactionSearchOKBodyAO0

	// PostAccountantTransactionSearchOKBodyAO1
	var dataPostAccountantTransactionSearchOKBodyAO1 struct {
		Data *models.FindTransactionResponse `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostAccountantTransactionSearchOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostAccountantTransactionSearchOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostAccountantTransactionSearchOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postAccountantTransactionSearchOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postAccountantTransactionSearchOKBodyAO0)
	var dataPostAccountantTransactionSearchOKBodyAO1 struct {
		Data *models.FindTransactionResponse `json:"data,omitempty"`
	}

	dataPostAccountantTransactionSearchOKBodyAO1.Data = o.Data

	jsonDataPostAccountantTransactionSearchOKBodyAO1, errPostAccountantTransactionSearchOKBodyAO1 := swag.WriteJSON(dataPostAccountantTransactionSearchOKBodyAO1)
	if errPostAccountantTransactionSearchOKBodyAO1 != nil {
		return nil, errPostAccountantTransactionSearchOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostAccountantTransactionSearchOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post accountant transaction search o k body
func (o *PostAccountantTransactionSearchOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostAccountantTransactionSearchOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postAccountantTransactionSearchOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postAccountantTransactionSearchOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post accountant transaction search o k body based on the context it is used
func (o *PostAccountantTransactionSearchOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostAccountantTransactionSearchOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postAccountantTransactionSearchOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postAccountantTransactionSearchOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostAccountantTransactionSearchOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAccountantTransactionSearchOKBody) UnmarshalBinary(b []byte) error {
	var res PostAccountantTransactionSearchOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}