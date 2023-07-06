// Code generated by go-swagger; DO NOT EDIT.

package appstore

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

// PostAppstoreCustomerFidGoogleReader is a Reader for the PostAppstoreCustomerFidGoogle structure.
type PostAppstoreCustomerFidGoogleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAppstoreCustomerFidGoogleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAppstoreCustomerFidGoogleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostAppstoreCustomerFidGoogleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostAppstoreCustomerFidGoogleOK creates a PostAppstoreCustomerFidGoogleOK with default headers values
func NewPostAppstoreCustomerFidGoogleOK() *PostAppstoreCustomerFidGoogleOK {
	return &PostAppstoreCustomerFidGoogleOK{}
}

/*
PostAppstoreCustomerFidGoogleOK describes a response with status code 200, with default header values.

Bool Response
*/
type PostAppstoreCustomerFidGoogleOK struct {
	Payload *PostAppstoreCustomerFidGoogleOKBody
}

// IsSuccess returns true when this post appstore customer fid google o k response has a 2xx status code
func (o *PostAppstoreCustomerFidGoogleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post appstore customer fid google o k response has a 3xx status code
func (o *PostAppstoreCustomerFidGoogleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post appstore customer fid google o k response has a 4xx status code
func (o *PostAppstoreCustomerFidGoogleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post appstore customer fid google o k response has a 5xx status code
func (o *PostAppstoreCustomerFidGoogleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post appstore customer fid google o k response a status code equal to that given
func (o *PostAppstoreCustomerFidGoogleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post appstore customer fid google o k response
func (o *PostAppstoreCustomerFidGoogleOK) Code() int {
	return 200
}

func (o *PostAppstoreCustomerFidGoogleOK) Error() string {
	return fmt.Sprintf("[POST /appstore/{customerFid}/google][%d] postAppstoreCustomerFidGoogleOK  %+v", 200, o.Payload)
}

func (o *PostAppstoreCustomerFidGoogleOK) String() string {
	return fmt.Sprintf("[POST /appstore/{customerFid}/google][%d] postAppstoreCustomerFidGoogleOK  %+v", 200, o.Payload)
}

func (o *PostAppstoreCustomerFidGoogleOK) GetPayload() *PostAppstoreCustomerFidGoogleOKBody {
	return o.Payload
}

func (o *PostAppstoreCustomerFidGoogleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostAppstoreCustomerFidGoogleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAppstoreCustomerFidGoogleDefault creates a PostAppstoreCustomerFidGoogleDefault with default headers values
func NewPostAppstoreCustomerFidGoogleDefault(code int) *PostAppstoreCustomerFidGoogleDefault {
	return &PostAppstoreCustomerFidGoogleDefault{
		_statusCode: code,
	}
}

/*
PostAppstoreCustomerFidGoogleDefault describes a response with status code -1, with default header values.

Error
*/
type PostAppstoreCustomerFidGoogleDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this post appstore customer fid google default response has a 2xx status code
func (o *PostAppstoreCustomerFidGoogleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post appstore customer fid google default response has a 3xx status code
func (o *PostAppstoreCustomerFidGoogleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post appstore customer fid google default response has a 4xx status code
func (o *PostAppstoreCustomerFidGoogleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post appstore customer fid google default response has a 5xx status code
func (o *PostAppstoreCustomerFidGoogleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post appstore customer fid google default response a status code equal to that given
func (o *PostAppstoreCustomerFidGoogleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post appstore customer fid google default response
func (o *PostAppstoreCustomerFidGoogleDefault) Code() int {
	return o._statusCode
}

func (o *PostAppstoreCustomerFidGoogleDefault) Error() string {
	return fmt.Sprintf("[POST /appstore/{customerFid}/google][%d] PostAppstoreCustomerFidGoogle default  %+v", o._statusCode, o.Payload)
}

func (o *PostAppstoreCustomerFidGoogleDefault) String() string {
	return fmt.Sprintf("[POST /appstore/{customerFid}/google][%d] PostAppstoreCustomerFidGoogle default  %+v", o._statusCode, o.Payload)
}

func (o *PostAppstoreCustomerFidGoogleDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostAppstoreCustomerFidGoogleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostAppstoreCustomerFidGoogleOKBody post appstore customer fid google o k body
swagger:model PostAppstoreCustomerFidGoogleOKBody
*/
type PostAppstoreCustomerFidGoogleOKBody struct {
	models.Envelope

	// data
	Data *models.BoolMessage `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostAppstoreCustomerFidGoogleOKBody) UnmarshalJSON(raw []byte) error {
	// PostAppstoreCustomerFidGoogleOKBodyAO0
	var postAppstoreCustomerFidGoogleOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postAppstoreCustomerFidGoogleOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postAppstoreCustomerFidGoogleOKBodyAO0

	// PostAppstoreCustomerFidGoogleOKBodyAO1
	var dataPostAppstoreCustomerFidGoogleOKBodyAO1 struct {
		Data *models.BoolMessage `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostAppstoreCustomerFidGoogleOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostAppstoreCustomerFidGoogleOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostAppstoreCustomerFidGoogleOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postAppstoreCustomerFidGoogleOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postAppstoreCustomerFidGoogleOKBodyAO0)
	var dataPostAppstoreCustomerFidGoogleOKBodyAO1 struct {
		Data *models.BoolMessage `json:"data,omitempty"`
	}

	dataPostAppstoreCustomerFidGoogleOKBodyAO1.Data = o.Data

	jsonDataPostAppstoreCustomerFidGoogleOKBodyAO1, errPostAppstoreCustomerFidGoogleOKBodyAO1 := swag.WriteJSON(dataPostAppstoreCustomerFidGoogleOKBodyAO1)
	if errPostAppstoreCustomerFidGoogleOKBodyAO1 != nil {
		return nil, errPostAppstoreCustomerFidGoogleOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostAppstoreCustomerFidGoogleOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post appstore customer fid google o k body
func (o *PostAppstoreCustomerFidGoogleOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostAppstoreCustomerFidGoogleOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postAppstoreCustomerFidGoogleOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postAppstoreCustomerFidGoogleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post appstore customer fid google o k body based on the context it is used
func (o *PostAppstoreCustomerFidGoogleOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostAppstoreCustomerFidGoogleOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postAppstoreCustomerFidGoogleOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postAppstoreCustomerFidGoogleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostAppstoreCustomerFidGoogleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAppstoreCustomerFidGoogleOKBody) UnmarshalBinary(b []byte) error {
	var res PostAppstoreCustomerFidGoogleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}