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

// PostInteractionsJourneyReader is a Reader for the PostInteractionsJourney structure.
type PostInteractionsJourneyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostInteractionsJourneyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostInteractionsJourneyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostInteractionsJourneyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostInteractionsJourneyOK creates a PostInteractionsJourneyOK with default headers values
func NewPostInteractionsJourneyOK() *PostInteractionsJourneyOK {
	return &PostInteractionsJourneyOK{}
}

/*
PostInteractionsJourneyOK describes a response with status code 200, with default header values.

Customer Journey Created
*/
type PostInteractionsJourneyOK struct {
	Payload *PostInteractionsJourneyOKBody
}

// IsSuccess returns true when this post interactions journey o k response has a 2xx status code
func (o *PostInteractionsJourneyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post interactions journey o k response has a 3xx status code
func (o *PostInteractionsJourneyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post interactions journey o k response has a 4xx status code
func (o *PostInteractionsJourneyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post interactions journey o k response has a 5xx status code
func (o *PostInteractionsJourneyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post interactions journey o k response a status code equal to that given
func (o *PostInteractionsJourneyOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostInteractionsJourneyOK) Error() string {
	return fmt.Sprintf("[POST /interactions/journey][%d] postInteractionsJourneyOK  %+v", 200, o.Payload)
}

func (o *PostInteractionsJourneyOK) String() string {
	return fmt.Sprintf("[POST /interactions/journey][%d] postInteractionsJourneyOK  %+v", 200, o.Payload)
}

func (o *PostInteractionsJourneyOK) GetPayload() *PostInteractionsJourneyOKBody {
	return o.Payload
}

func (o *PostInteractionsJourneyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostInteractionsJourneyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInteractionsJourneyDefault creates a PostInteractionsJourneyDefault with default headers values
func NewPostInteractionsJourneyDefault(code int) *PostInteractionsJourneyDefault {
	return &PostInteractionsJourneyDefault{
		_statusCode: code,
	}
}

/*
PostInteractionsJourneyDefault describes a response with status code -1, with default header values.

Error
*/
type PostInteractionsJourneyDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// Code gets the status code for the post interactions journey default response
func (o *PostInteractionsJourneyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this post interactions journey default response has a 2xx status code
func (o *PostInteractionsJourneyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post interactions journey default response has a 3xx status code
func (o *PostInteractionsJourneyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post interactions journey default response has a 4xx status code
func (o *PostInteractionsJourneyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post interactions journey default response has a 5xx status code
func (o *PostInteractionsJourneyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post interactions journey default response a status code equal to that given
func (o *PostInteractionsJourneyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PostInteractionsJourneyDefault) Error() string {
	return fmt.Sprintf("[POST /interactions/journey][%d] PostInteractionsJourney default  %+v", o._statusCode, o.Payload)
}

func (o *PostInteractionsJourneyDefault) String() string {
	return fmt.Sprintf("[POST /interactions/journey][%d] PostInteractionsJourney default  %+v", o._statusCode, o.Payload)
}

func (o *PostInteractionsJourneyDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostInteractionsJourneyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostInteractionsJourneyOKBody post interactions journey o k body
swagger:model PostInteractionsJourneyOKBody
*/
type PostInteractionsJourneyOKBody struct {
	models.Envelope

	// data
	Data *models.Fid `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostInteractionsJourneyOKBody) UnmarshalJSON(raw []byte) error {
	// PostInteractionsJourneyOKBodyAO0
	var postInteractionsJourneyOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postInteractionsJourneyOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postInteractionsJourneyOKBodyAO0

	// PostInteractionsJourneyOKBodyAO1
	var dataPostInteractionsJourneyOKBodyAO1 struct {
		Data *models.Fid `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostInteractionsJourneyOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostInteractionsJourneyOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostInteractionsJourneyOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postInteractionsJourneyOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postInteractionsJourneyOKBodyAO0)
	var dataPostInteractionsJourneyOKBodyAO1 struct {
		Data *models.Fid `json:"data,omitempty"`
	}

	dataPostInteractionsJourneyOKBodyAO1.Data = o.Data

	jsonDataPostInteractionsJourneyOKBodyAO1, errPostInteractionsJourneyOKBodyAO1 := swag.WriteJSON(dataPostInteractionsJourneyOKBodyAO1)
	if errPostInteractionsJourneyOKBodyAO1 != nil {
		return nil, errPostInteractionsJourneyOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostInteractionsJourneyOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post interactions journey o k body
func (o *PostInteractionsJourneyOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostInteractionsJourneyOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postInteractionsJourneyOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postInteractionsJourneyOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post interactions journey o k body based on the context it is used
func (o *PostInteractionsJourneyOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostInteractionsJourneyOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postInteractionsJourneyOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postInteractionsJourneyOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostInteractionsJourneyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostInteractionsJourneyOKBody) UnmarshalBinary(b []byte) error {
	var res PostInteractionsJourneyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
