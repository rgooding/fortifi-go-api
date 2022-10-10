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

// PostProductsProductFidAvailabilityReserveReader is a Reader for the PostProductsProductFidAvailabilityReserve structure.
type PostProductsProductFidAvailabilityReserveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostProductsProductFidAvailabilityReserveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostProductsProductFidAvailabilityReserveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostProductsProductFidAvailabilityReserveDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostProductsProductFidAvailabilityReserveOK creates a PostProductsProductFidAvailabilityReserveOK with default headers values
func NewPostProductsProductFidAvailabilityReserveOK() *PostProductsProductFidAvailabilityReserveOK {
	return &PostProductsProductFidAvailabilityReserveOK{}
}

/*
PostProductsProductFidAvailabilityReserveOK describes a response with status code 200, with default header values.

Result of reservation
*/
type PostProductsProductFidAvailabilityReserveOK struct {
	Payload *PostProductsProductFidAvailabilityReserveOKBody
}

// IsSuccess returns true when this post products product fid availability reserve o k response has a 2xx status code
func (o *PostProductsProductFidAvailabilityReserveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post products product fid availability reserve o k response has a 3xx status code
func (o *PostProductsProductFidAvailabilityReserveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post products product fid availability reserve o k response has a 4xx status code
func (o *PostProductsProductFidAvailabilityReserveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post products product fid availability reserve o k response has a 5xx status code
func (o *PostProductsProductFidAvailabilityReserveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post products product fid availability reserve o k response a status code equal to that given
func (o *PostProductsProductFidAvailabilityReserveOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostProductsProductFidAvailabilityReserveOK) Error() string {
	return fmt.Sprintf("[POST /products/{productFid}/availability/reserve][%d] postProductsProductFidAvailabilityReserveOK  %+v", 200, o.Payload)
}

func (o *PostProductsProductFidAvailabilityReserveOK) String() string {
	return fmt.Sprintf("[POST /products/{productFid}/availability/reserve][%d] postProductsProductFidAvailabilityReserveOK  %+v", 200, o.Payload)
}

func (o *PostProductsProductFidAvailabilityReserveOK) GetPayload() *PostProductsProductFidAvailabilityReserveOKBody {
	return o.Payload
}

func (o *PostProductsProductFidAvailabilityReserveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostProductsProductFidAvailabilityReserveOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProductsProductFidAvailabilityReserveDefault creates a PostProductsProductFidAvailabilityReserveDefault with default headers values
func NewPostProductsProductFidAvailabilityReserveDefault(code int) *PostProductsProductFidAvailabilityReserveDefault {
	return &PostProductsProductFidAvailabilityReserveDefault{
		_statusCode: code,
	}
}

/*
PostProductsProductFidAvailabilityReserveDefault describes a response with status code -1, with default header values.

Error
*/
type PostProductsProductFidAvailabilityReserveDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post products product fid availability reserve default response
func (o *PostProductsProductFidAvailabilityReserveDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this post products product fid availability reserve default response has a 2xx status code
func (o *PostProductsProductFidAvailabilityReserveDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post products product fid availability reserve default response has a 3xx status code
func (o *PostProductsProductFidAvailabilityReserveDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post products product fid availability reserve default response has a 4xx status code
func (o *PostProductsProductFidAvailabilityReserveDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post products product fid availability reserve default response has a 5xx status code
func (o *PostProductsProductFidAvailabilityReserveDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post products product fid availability reserve default response a status code equal to that given
func (o *PostProductsProductFidAvailabilityReserveDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PostProductsProductFidAvailabilityReserveDefault) Error() string {
	return fmt.Sprintf("[POST /products/{productFid}/availability/reserve][%d] PostProductsProductFidAvailabilityReserve default  %+v", o._statusCode, o.Payload)
}

func (o *PostProductsProductFidAvailabilityReserveDefault) String() string {
	return fmt.Sprintf("[POST /products/{productFid}/availability/reserve][%d] PostProductsProductFidAvailabilityReserve default  %+v", o._statusCode, o.Payload)
}

func (o *PostProductsProductFidAvailabilityReserveDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostProductsProductFidAvailabilityReserveDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostProductsProductFidAvailabilityReserveOKBody post products product fid availability reserve o k body
swagger:model PostProductsProductFidAvailabilityReserveOKBody
*/
type PostProductsProductFidAvailabilityReserveOKBody struct {
	models.Envelope

	// data
	Data *models.AvailabilityReserveResponse `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostProductsProductFidAvailabilityReserveOKBody) UnmarshalJSON(raw []byte) error {
	// PostProductsProductFidAvailabilityReserveOKBodyAO0
	var postProductsProductFidAvailabilityReserveOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postProductsProductFidAvailabilityReserveOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postProductsProductFidAvailabilityReserveOKBodyAO0

	// PostProductsProductFidAvailabilityReserveOKBodyAO1
	var dataPostProductsProductFidAvailabilityReserveOKBodyAO1 struct {
		Data *models.AvailabilityReserveResponse `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostProductsProductFidAvailabilityReserveOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostProductsProductFidAvailabilityReserveOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostProductsProductFidAvailabilityReserveOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postProductsProductFidAvailabilityReserveOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postProductsProductFidAvailabilityReserveOKBodyAO0)
	var dataPostProductsProductFidAvailabilityReserveOKBodyAO1 struct {
		Data *models.AvailabilityReserveResponse `json:"data,omitempty"`
	}

	dataPostProductsProductFidAvailabilityReserveOKBodyAO1.Data = o.Data

	jsonDataPostProductsProductFidAvailabilityReserveOKBodyAO1, errPostProductsProductFidAvailabilityReserveOKBodyAO1 := swag.WriteJSON(dataPostProductsProductFidAvailabilityReserveOKBodyAO1)
	if errPostProductsProductFidAvailabilityReserveOKBodyAO1 != nil {
		return nil, errPostProductsProductFidAvailabilityReserveOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostProductsProductFidAvailabilityReserveOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post products product fid availability reserve o k body
func (o *PostProductsProductFidAvailabilityReserveOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostProductsProductFidAvailabilityReserveOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postProductsProductFidAvailabilityReserveOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postProductsProductFidAvailabilityReserveOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post products product fid availability reserve o k body based on the context it is used
func (o *PostProductsProductFidAvailabilityReserveOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostProductsProductFidAvailabilityReserveOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postProductsProductFidAvailabilityReserveOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postProductsProductFidAvailabilityReserveOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostProductsProductFidAvailabilityReserveOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostProductsProductFidAvailabilityReserveOKBody) UnmarshalBinary(b []byte) error {
	var res PostProductsProductFidAvailabilityReserveOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
