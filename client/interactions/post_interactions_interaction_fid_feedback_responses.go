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

// PostInteractionsInteractionFidFeedbackReader is a Reader for the PostInteractionsInteractionFidFeedback structure.
type PostInteractionsInteractionFidFeedbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostInteractionsInteractionFidFeedbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostInteractionsInteractionFidFeedbackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostInteractionsInteractionFidFeedbackDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostInteractionsInteractionFidFeedbackOK creates a PostInteractionsInteractionFidFeedbackOK with default headers values
func NewPostInteractionsInteractionFidFeedbackOK() *PostInteractionsInteractionFidFeedbackOK {
	return &PostInteractionsInteractionFidFeedbackOK{}
}

/*
PostInteractionsInteractionFidFeedbackOK describes a response with status code 200, with default header values.

Feedback Updated Response
*/
type PostInteractionsInteractionFidFeedbackOK struct {
	Payload *PostInteractionsInteractionFidFeedbackOKBody
}

// IsSuccess returns true when this post interactions interaction fid feedback o k response has a 2xx status code
func (o *PostInteractionsInteractionFidFeedbackOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post interactions interaction fid feedback o k response has a 3xx status code
func (o *PostInteractionsInteractionFidFeedbackOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post interactions interaction fid feedback o k response has a 4xx status code
func (o *PostInteractionsInteractionFidFeedbackOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post interactions interaction fid feedback o k response has a 5xx status code
func (o *PostInteractionsInteractionFidFeedbackOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post interactions interaction fid feedback o k response a status code equal to that given
func (o *PostInteractionsInteractionFidFeedbackOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post interactions interaction fid feedback o k response
func (o *PostInteractionsInteractionFidFeedbackOK) Code() int {
	return 200
}

func (o *PostInteractionsInteractionFidFeedbackOK) Error() string {
	return fmt.Sprintf("[POST /interactions/{interactionFid}/feedback][%d] postInteractionsInteractionFidFeedbackOK  %+v", 200, o.Payload)
}

func (o *PostInteractionsInteractionFidFeedbackOK) String() string {
	return fmt.Sprintf("[POST /interactions/{interactionFid}/feedback][%d] postInteractionsInteractionFidFeedbackOK  %+v", 200, o.Payload)
}

func (o *PostInteractionsInteractionFidFeedbackOK) GetPayload() *PostInteractionsInteractionFidFeedbackOKBody {
	return o.Payload
}

func (o *PostInteractionsInteractionFidFeedbackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostInteractionsInteractionFidFeedbackOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInteractionsInteractionFidFeedbackDefault creates a PostInteractionsInteractionFidFeedbackDefault with default headers values
func NewPostInteractionsInteractionFidFeedbackDefault(code int) *PostInteractionsInteractionFidFeedbackDefault {
	return &PostInteractionsInteractionFidFeedbackDefault{
		_statusCode: code,
	}
}

/*
PostInteractionsInteractionFidFeedbackDefault describes a response with status code -1, with default header values.

Error
*/
type PostInteractionsInteractionFidFeedbackDefault struct {
	_statusCode int

	Payload *models.Envelope
}

// IsSuccess returns true when this post interactions interaction fid feedback default response has a 2xx status code
func (o *PostInteractionsInteractionFidFeedbackDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post interactions interaction fid feedback default response has a 3xx status code
func (o *PostInteractionsInteractionFidFeedbackDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post interactions interaction fid feedback default response has a 4xx status code
func (o *PostInteractionsInteractionFidFeedbackDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post interactions interaction fid feedback default response has a 5xx status code
func (o *PostInteractionsInteractionFidFeedbackDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post interactions interaction fid feedback default response a status code equal to that given
func (o *PostInteractionsInteractionFidFeedbackDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the post interactions interaction fid feedback default response
func (o *PostInteractionsInteractionFidFeedbackDefault) Code() int {
	return o._statusCode
}

func (o *PostInteractionsInteractionFidFeedbackDefault) Error() string {
	return fmt.Sprintf("[POST /interactions/{interactionFid}/feedback][%d] PostInteractionsInteractionFidFeedback default  %+v", o._statusCode, o.Payload)
}

func (o *PostInteractionsInteractionFidFeedbackDefault) String() string {
	return fmt.Sprintf("[POST /interactions/{interactionFid}/feedback][%d] PostInteractionsInteractionFidFeedback default  %+v", o._statusCode, o.Payload)
}

func (o *PostInteractionsInteractionFidFeedbackDefault) GetPayload() *models.Envelope {
	return o.Payload
}

func (o *PostInteractionsInteractionFidFeedbackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Envelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostInteractionsInteractionFidFeedbackOKBody post interactions interaction fid feedback o k body
swagger:model PostInteractionsInteractionFidFeedbackOKBody
*/
type PostInteractionsInteractionFidFeedbackOKBody struct {
	models.Envelope

	// data
	Data *models.BoolMessage `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostInteractionsInteractionFidFeedbackOKBody) UnmarshalJSON(raw []byte) error {
	// PostInteractionsInteractionFidFeedbackOKBodyAO0
	var postInteractionsInteractionFidFeedbackOKBodyAO0 models.Envelope
	if err := swag.ReadJSON(raw, &postInteractionsInteractionFidFeedbackOKBodyAO0); err != nil {
		return err
	}
	o.Envelope = postInteractionsInteractionFidFeedbackOKBodyAO0

	// PostInteractionsInteractionFidFeedbackOKBodyAO1
	var dataPostInteractionsInteractionFidFeedbackOKBodyAO1 struct {
		Data *models.BoolMessage `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostInteractionsInteractionFidFeedbackOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostInteractionsInteractionFidFeedbackOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostInteractionsInteractionFidFeedbackOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postInteractionsInteractionFidFeedbackOKBodyAO0, err := swag.WriteJSON(o.Envelope)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postInteractionsInteractionFidFeedbackOKBodyAO0)
	var dataPostInteractionsInteractionFidFeedbackOKBodyAO1 struct {
		Data *models.BoolMessage `json:"data,omitempty"`
	}

	dataPostInteractionsInteractionFidFeedbackOKBodyAO1.Data = o.Data

	jsonDataPostInteractionsInteractionFidFeedbackOKBodyAO1, errPostInteractionsInteractionFidFeedbackOKBodyAO1 := swag.WriteJSON(dataPostInteractionsInteractionFidFeedbackOKBodyAO1)
	if errPostInteractionsInteractionFidFeedbackOKBodyAO1 != nil {
		return nil, errPostInteractionsInteractionFidFeedbackOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostInteractionsInteractionFidFeedbackOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post interactions interaction fid feedback o k body
func (o *PostInteractionsInteractionFidFeedbackOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostInteractionsInteractionFidFeedbackOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postInteractionsInteractionFidFeedbackOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postInteractionsInteractionFidFeedbackOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post interactions interaction fid feedback o k body based on the context it is used
func (o *PostInteractionsInteractionFidFeedbackOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostInteractionsInteractionFidFeedbackOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postInteractionsInteractionFidFeedbackOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postInteractionsInteractionFidFeedbackOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostInteractionsInteractionFidFeedbackOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostInteractionsInteractionFidFeedbackOKBody) UnmarshalBinary(b []byte) error {
	var res PostInteractionsInteractionFidFeedbackOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}