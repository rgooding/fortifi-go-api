// Code generated by go-swagger; DO NOT EDIT.

package products

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

// PostProductsProductFidAvailabilityCheckReader is a Reader for the PostProductsProductFidAvailabilityCheck structure.
type PostProductsProductFidAvailabilityCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostProductsProductFidAvailabilityCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostProductsProductFidAvailabilityCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostProductsProductFidAvailabilityCheckDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostProductsProductFidAvailabilityCheckOK creates a PostProductsProductFidAvailabilityCheckOK with default headers values
func NewPostProductsProductFidAvailabilityCheckOK() *PostProductsProductFidAvailabilityCheckOK {
	return &PostProductsProductFidAvailabilityCheckOK{}
}

/*
PostProductsProductFidAvailabilityCheckOK describes a response with status code 200, with default header values.

Result of availability check
*/
type PostProductsProductFidAvailabilityCheckOK struct {
	Payload *PostProductsProductFidAvailabilityCheckOKBody
}

// IsSuccess returns true when this post products product fid availability check o k response has a 2xx status code
func (o *PostProductsProductFidAvailabilityCheckOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post products product fid availability check o k response has a 3xx status code
func (o *PostProductsProductFidAvailabilityCheckOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post products product fid availability check o k response has a 4xx status code
func (o *PostProductsProductFidAvailabilityCheckOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post products product fid availability check o k response has a 5xx status code
func (o *PostProductsProductFidAvailabilityCheckOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post products product fid availability check o k response a status code equal to that given
func (o *PostProductsProductFidAvailabilityCheckOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostProductsProductFidAvailabilityCheckOK) Error() string {
	return fmt.Sprintf("[POST /products/{productFid}/availability/check][%d] postProductsProductFidAvailabilityCheckOK  %+v", 200, o.Payload)
}

func (o *PostProductsProductFidAvailabilityCheckOK) String() string {
	return fmt.Sprintf("[POST /products/{productFid}/availability/check][%d] postProductsProductFidAvailabilityCheckOK  %+v", 200, o.Payload)
}

func (o *PostProductsProductFidAvailabilityCheckOK) GetPayload() *PostProductsProductFidAvailabilityCheckOKBody {
	return o.Payload
}

func (o *PostProductsProductFidAvailabilityCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostProductsProductFidAvailabilityCheckOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProductsProductFidAvailabilityCheckDefault creates a PostProductsProductFidAvailabilityCheckDefault with default headers values
func NewPostProductsProductFidAvailabilityCheckDefault(code int) *PostProductsProductFidAvailabilityCheckDefault {
	return &PostProductsProductFidAvailabilityCheckDefault{
		_statusCode: code,
	}
}

/*
PostProductsProductFidAvailabilityCheckDefault describes a response with status code -1, with default header values.

Error
*/
type PostProductsProductFidAvailabilityCheckDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post products product fid availability check default response
func (o *PostProductsProductFidAvailabilityCheckDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this post products product fid availability check default response has a 2xx status code
func (o *PostProductsProductFidAvailabilityCheckDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post products product fid availability check default response has a 3xx status code
func (o *PostProductsProductFidAvailabilityCheckDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post products product fid availability check default response has a 4xx status code
func (o *PostProductsProductFidAvailabilityCheckDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post products product fid availability check default response has a 5xx status code
func (o *PostProductsProductFidAvailabilityCheckDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post products product fid availability check default response a status code equal to that given
func (o *PostProductsProductFidAvailabilityCheckDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PostProductsProductFidAvailabilityCheckDefault) Error() string {
	return fmt.Sprintf("[POST /products/{productFid}/availability/check][%d] PostProductsProductFidAvailabilityCheck default  %+v", o._statusCode, o.Payload)
}

func (o *PostProductsProductFidAvailabilityCheckDefault) String() string {
	return fmt.Sprintf("[POST /products/{productFid}/availability/check][%d] PostProductsProductFidAvailabilityCheck default  %+v", o._statusCode, o.Payload)
}

func (o *PostProductsProductFidAvailabilityCheckDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostProductsProductFidAvailabilityCheckDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostProductsProductFidAvailabilityCheckOKBody post products product fid availability check o k body
swagger:model PostProductsProductFidAvailabilityCheckOKBody
*/
type PostProductsProductFidAvailabilityCheckOKBody struct {
	models.Envelope

	// data
	Data *models.AvailabilityCheckResponse `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostProductsProductFidAvailabilityCheckOKBody) UnmarshalJSON(raw []byte) error {
	// PostProductsProductFidAvailabilityCheckOKBodyAO0
	var postProductsProductFidAvailabilityCheckOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postProductsProductFidAvailabilityCheckOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postProductsProductFidAvailabilityCheckOKBodyAO0

	// PostProductsProductFidAvailabilityCheckOKBodyAO1
	var dataPostProductsProductFidAvailabilityCheckOKBodyAO1 struct {
		Data *models.AvailabilityCheckResponse `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostProductsProductFidAvailabilityCheckOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostProductsProductFidAvailabilityCheckOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostProductsProductFidAvailabilityCheckOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postProductsProductFidAvailabilityCheckOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postProductsProductFidAvailabilityCheckOKBodyAO0)
	var dataPostProductsProductFidAvailabilityCheckOKBodyAO1 struct {
		Data *models.AvailabilityCheckResponse `json:"data,omitempty"`
	}

	dataPostProductsProductFidAvailabilityCheckOKBodyAO1.Data = o.Data

	jsonDataPostProductsProductFidAvailabilityCheckOKBodyAO1, errPostProductsProductFidAvailabilityCheckOKBodyAO1 := swag.WriteJSON(dataPostProductsProductFidAvailabilityCheckOKBodyAO1)
	if errPostProductsProductFidAvailabilityCheckOKBodyAO1 != nil {
		return nil, errPostProductsProductFidAvailabilityCheckOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostProductsProductFidAvailabilityCheckOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post products product fid availability check o k body
func (o *PostProductsProductFidAvailabilityCheckOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostProductsProductFidAvailabilityCheckOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postProductsProductFidAvailabilityCheckOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postProductsProductFidAvailabilityCheckOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post products product fid availability check o k body based on the context it is used
func (o *PostProductsProductFidAvailabilityCheckOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostProductsProductFidAvailabilityCheckOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postProductsProductFidAvailabilityCheckOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postProductsProductFidAvailabilityCheckOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostProductsProductFidAvailabilityCheckOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostProductsProductFidAvailabilityCheckOKBody) UnmarshalBinary(b []byte) error {
	var res PostProductsProductFidAvailabilityCheckOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
